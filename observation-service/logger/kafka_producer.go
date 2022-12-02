package logger

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/caraml-dev/observation-service/observation-service/config"
	"github.com/caraml-dev/observation-service/observation-service/log"
	"github.com/caraml-dev/observation-service/observation-service/monitoring"
	"github.com/caraml-dev/observation-service/observation-service/services"
	"github.com/caraml-dev/observation-service/observation-service/types"
)

type kafkaProducer interface {
	GetMetadata(*string, bool, int) (*kafka.Metadata, error)
	Produce(*kafka.Message, chan kafka.Event) error
}

// KafkaLogPublisher captures configs for publishing ObservationLog to a Kafka topic
type KafkaLogPublisher struct {
	topic          string
	producer       kafkaProducer
	metricsService services.MetricService
}

// NewKafkaLogProducer initializes a KafkaLogPublisher struct
func NewKafkaLogProducer(
	cfg config.KafkaConfig,
	metricsService services.MetricService,
) (*KafkaLogPublisher, error) {
	// Create Kafka Producer
	producer, err := newKafkaProducer(cfg.Brokers, cfg.MaxMessageBytes, cfg.CompressionType)
	if err != nil {
		return nil, err
	}
	// Test that we are able to query the broker on the topic. If the topic
	// does not already exist on the broker, this should create it.
	_, err = producer.GetMetadata(&cfg.Topic, false, cfg.ConnectTimeoutMS)
	if err != nil {
		return nil, fmt.Errorf("error Querying topic %s from Kafka broker(s): %s", cfg.Topic, err)
	}
	// Create Kafka Logger
	return &KafkaLogPublisher{
		topic:          cfg.Topic,
		producer:       producer,
		metricsService: metricsService,
	}, nil
}

func newKafkaProducer(
	kafkaBrokers string,
	kafkaMaxMessageBytes int,
	kafkaCompressionType string,
) (kafkaProducer, error) {
	producer, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": kafkaBrokers,
			"message.max.bytes": kafkaMaxMessageBytes,
			"compression.type":  kafkaCompressionType,
		},
	)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

// Produce logs ObservationLog to a Kafka topic
func (p *KafkaLogPublisher) Produce(logs []*types.ObservationLogEntry) {
	deliveryChan := make(chan kafka.Event, 1)
	defer close(deliveryChan)

	for _, l := range logs {
		keyBytes, valueBytes, err := newKafkaLogEntry(l)
		if err != nil {
			log.Error(err)
		}

		err = p.producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &p.topic,
				Partition: kafka.PartitionAny},
			Value: valueBytes,
			Key:   keyBytes,
		}, deliveryChan)
		if err != nil {
			log.Error(err)
		}

		// Get delivery response
		event := <-deliveryChan
		msg := event.(*kafka.Message)
		if msg.TopicPartition.Error != nil {
			// TODO: Send failed ObservationLog to deadletter sink
			p.metricsService.LogRequestCount(http.StatusInternalServerError, monitoring.FlushObservationCount)
			log.Errorf("delivery failed: %v", msg.TopicPartition.Error)
		}
		p.metricsService.LogRequestCount(http.StatusOK, monitoring.FlushObservationCount)
		p.metricsService.LogLatencyHistogram(l.StartTime, http.StatusOK, monitoring.FlushDurationMs)
	}
}

func newKafkaLogEntry(
	log *types.ObservationLogEntry,
) (keyBytes []byte, valueBytes []byte, err error) {
	// Create the Kafka key
	key := &types.ObservationLogKey{
		ObservationBatchId: log.BatchID,
		PredictionId:       log.PredictionId,
		RowId:              log.RowId,
	}
	// Marshal the key
	keyBytes, err = json.Marshal(key)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to marshal log entry key, %s", err)
	}

	// Marshal the message
	valueBytes, err = log.MarshalJSON()
	if err != nil {
		return nil, nil, fmt.Errorf("unable to marshal log entry value, %s", err)
	}

	return keyBytes, valueBytes, nil
}
