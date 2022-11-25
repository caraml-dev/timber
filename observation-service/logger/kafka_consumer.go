package logger

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	upiv1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"

	"github.com/caraml-dev/observation-service/observation-service/config"
	"github.com/caraml-dev/observation-service/observation-service/types"
)

type kafkaConsumer interface {
	GetMetadata(*string, bool, int) (*kafka.Metadata, error)
	Subscribe(string, kafka.RebalanceCb) error
	Poll(int) kafka.Event
	Close() error
}

type KafkaLogConsumer struct {
	topic    string
	consumer kafkaConsumer
}

func NewKafkaLogConsumer(
	cfg config.KafkaConsumerConfig,
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
		topic:    cfg.Topic,
		consumer: consumer,
	}

	return kafkaLogConsumer, nil
}

// newKafkaConsumer creates a new Kafka consumer and subscribes to relevant Kafka topic
func newKafkaConsumer(
	cfg config.KafkaConsumerConfig,
) (kafkaConsumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Brokers,
		"group.id":          "observation-service",
		// At-least-once semantics
		"auto.offset.reset": "earliest",
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

func (k *KafkaLogConsumer) Consume(logsChannel chan *types.ObservationLogEntry) error {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case sig := <-sigchan:
			// Capture Ctrl-C interrupt
			log.Println("System interrupt detected:", sig)

			// Close consumer before exit
			if err := k.consumer.Close(); err != nil {
				log.Println("Failed to close consumer:", err)
				return err
			}
			// Wait for awhile before close
			time.Sleep(2 * time.Second)
			break
		default:
			// Log errors as we don't want to crash the server due to bad records
			ev := k.consumer.Poll(1000)
			switch e := ev.(type) {
			case *kafka.Message:
				decodedLogMessage := &upiv1.ObservationLog{}
				err := proto.Unmarshal(e.Value, decodedLogMessage)
				if err != nil {
					log.Println(err)
				}
				convertedLogMessage := types.NewObservationLogEntry(decodedLogMessage)

				logsChannel <- convertedLogMessage
			case kafka.PartitionEOF:
				log.Printf("%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			default:
			}
		}
	}
}
