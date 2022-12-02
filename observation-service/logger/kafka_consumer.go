package logger

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	upiv1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"

	"github.com/caraml-dev/observation-service/observation-service/config"
	"github.com/caraml-dev/observation-service/observation-service/log"
	"github.com/caraml-dev/observation-service/observation-service/monitoring"
	"github.com/caraml-dev/observation-service/observation-service/services"
	"github.com/caraml-dev/observation-service/observation-service/types"
)

type kafkaConsumer interface {
	GetMetadata(*string, bool, int) (*kafka.Metadata, error)
	Subscribe(string, kafka.RebalanceCb) error
	Poll(int) kafka.Event
	Close() error
}

// KafkaLogConsumer captures configs for polling ObservationLog from a Kafka topic
type KafkaLogConsumer struct {
	pollInterval   int
	topic          string
	consumer       kafkaConsumer
	metricsService services.MetricService
}

// NewKafkaLogConsumer initializes a KafkaLogConsumer struct
func NewKafkaLogConsumer(
	cfg config.KafkaConfig,
	metricsService services.MetricService,
) (*KafkaLogConsumer, error) {
	consumer, err := newKafkaConsumer(cfg)
	if err != nil {
		return nil, err
	}
	// Test that we are able to query the broker on the topic. If the topic
	// does not already exist on the broker, this should create it.
	_, err = consumer.GetMetadata(&cfg.Topic, false, cfg.ConnectTimeoutMS)
	if err != nil {
		return nil, fmt.Errorf("error Querying topic %s from Kafka broker(s): %s", cfg.Topic, err)
	}

	kafkaLogConsumer := &KafkaLogConsumer{
		pollInterval:   cfg.PollInterval,
		topic:          cfg.Topic,
		consumer:       consumer,
		metricsService: metricsService,
	}

	return kafkaLogConsumer, nil
}

// newKafkaConsumer creates a new Kafka consumer and subscribes to relevant Kafka topic
func newKafkaConsumer(
	cfg config.KafkaConfig,
) (kafkaConsumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Brokers,
		"group.id":          "observation-service",
		"auto.offset.reset": cfg.AutoOffsetReset,
	})
	if err != nil {
		return nil, err
	}

	err = consumer.Subscribe(cfg.Topic, nil)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

// Consume polls for ObservationLog from a Kafka topic to a buffered Go channel
func (k *KafkaLogConsumer) Consume(logsChannel chan *types.ObservationLogEntry) error {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case sig := <-sigchan:
			// Capture Ctrl-C interrupt
			log.Glob().Infof("System interrupt detected: %s", sig)

			// Close consumer before exit
			if err := k.consumer.Close(); err != nil {
				log.Glob().Errorf("Failed to close consumer:", err)
				return err
			}
			// Wait for awhile before close
			time.Sleep(2 * time.Second)
			break
		default:
			// Log errors as we don't want to crash the server due to bad records
			ev := k.consumer.Poll(k.pollInterval)
			switch e := ev.(type) {
			case *kafka.Message:
				decodedLogMessage := &upiv1.ObservationLog{}
				err := proto.Unmarshal(e.Value, decodedLogMessage)
				if err != nil {
					log.Glob().Error(err)
				}
				convertedLogMessage := types.NewObservationLogEntry(decodedLogMessage)
				k.metricsService.LogRequestCount(http.StatusOK, monitoring.ReadCount)

				logsChannel <- convertedLogMessage
			case kafka.PartitionEOF:
				k.metricsService.LogRequestCount(http.StatusInternalServerError, monitoring.ReadCount)
				log.Glob().Errorf("%% Reached %v\n", e)
			case kafka.Error:
				k.metricsService.LogRequestCount(http.StatusInternalServerError, monitoring.ReadCount)
				log.Glob().Errorf("%% Error: %v\n", os.Stderr)
			default:
			}
		}
	}
}
