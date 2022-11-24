package logger

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"
	"google.golang.org/protobuf/proto"

	"github.com/caraml-dev/observation-service/observation-service/models"
	upiv1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// kafkaConsumer contains GetMetadata and Produce methods for mocking in unit tests
type kafkaConsumer interface {
	GetMetadata(*string, bool, int) (*kafka.Metadata, error)
	Subscribe(string, kafka.RebalanceCb) error
	Poll(int) kafka.Event
	Close() error
}

type KafkaLogConsumer struct {
	channel chan *upiv1.ObservationLog

	topic    string
	consumer kafkaConsumer
}

func NewKafkaLogConsumer(
	kafkaBrokers string,
	kafkaTopic string,
	KafkaConnectTimeoutMS int,
	queueChannel chan *upiv1.ObservationLog,
) (*KafkaLogConsumer, error) {
	consumer, err := newKafkaConsumer(kafkaBrokers, kafkaTopic, KafkaConnectTimeoutMS)
	if err != nil {
		return nil, err
	}

	kafkaLogConsumer := &KafkaLogConsumer{
		channel:  queueChannel,
		topic:    kafkaTopic,
		consumer: consumer,
	}

	return kafkaLogConsumer, nil
}

// newKafkaConsumer creates a new Kafka consumer and subscribes to relevant Kafka topic
func newKafkaConsumer(
	kafkaBrokers string,
	kafkaTopic string,
	KafkaConnectTimeoutMS int,
) (kafkaConsumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaBrokers,
		"group.id":          "observation-service",
	})
	if err != nil {
		return nil, err
	}

	err = consumer.Subscribe(kafkaTopic, nil)
	if err != nil {
		return nil, err
	}

	return consumer, nil
}

func (k *KafkaLogConsumer) Consume(queueChannel chan *upiv1.ObservationLog) error {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// TODO: How do we want to handle poll errors from Kafka topic?
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
			ev := k.consumer.Poll(1000)
			switch e := ev.(type) {
			case *kafka.Message:
				log.Printf("%% Message on %s:\n%s\n",
					e.TopicPartition, string(e.Value))
				decodedLogMessage := &upiv1.ObservationLog{}
				err := proto.Unmarshal(e.Value, decodedLogMessage)
				if err != nil {
					log.Println(err)
				}

				// Configure FluentD
				logger, err := fluent.New(fluent.Config{FluentPort: 24224, FluentHost: "localhost"})
				if err != nil {
					log.Println(err)
				}
				tag := "observation-service.access"

				convertedLogMessage, err := models.NewObservationLogEntry(decodedLogMessage).Value()
				if err != nil {
					log.Println(err)
				}

				err = logger.Post(tag, convertedLogMessage)
				if err != nil {
					log.Println(err)
				}
			case kafka.PartitionEOF:
				log.Printf("%% Reached %v\n", e)
			case kafka.Error:
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			default:
			}
		}
	}
}
