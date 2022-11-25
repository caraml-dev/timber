package logger

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/caraml-dev/observation-service/observation-service/models"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type kafkaProducer interface {
	GetMetadata(*string, bool, int) (*kafka.Metadata, error)
	Produce(*kafka.Message, chan kafka.Event) error
}

type KafkaLogPublisher struct {
	topic    string
	producer kafkaProducer
}

func NewKafkaLogProducer(
	kafkaBrokers string,
	kafkaTopic string,
	kafkaMaxMessageBytes int,
	kafkaCompressionType string,
	KafkaConnectTimeoutMS int,
) (*KafkaLogPublisher, error) {
	// Create Kafka Producer
	producer, err := newKafkaProducer(kafkaBrokers, kafkaMaxMessageBytes, kafkaCompressionType)
	if err != nil {
		return nil, err
	}
	// Test that we are able to query the broker on the topic. If the topic
	// does not already exist on the broker, this should create it.
	_, err = producer.GetMetadata(&kafkaTopic, false, KafkaConnectTimeoutMS)
	if err != nil {
		return nil, fmt.Errorf("error Querying topic %s from Kafka broker(s): %s", kafkaTopic, err)
	}
	// Create Kafka Logger
	return &KafkaLogPublisher{
		topic:    kafkaTopic,
		producer: producer,
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

func (p *KafkaLogPublisher) Produce(logs []*models.ObservationLogEntry) error {
	deliveryChan := make(chan kafka.Event, 1)
	defer close(deliveryChan)

	for _, l := range logs {
		keyBytes, valueBytes, err := newKafkaLogEntry(l)
		if err != nil {
			return err
		}

		err = p.producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &p.topic,
				Partition: kafka.PartitionAny},
			Value: valueBytes,
			Key:   keyBytes,
		}, deliveryChan)
		if err != nil {
			return err
		}

		// Get delivery response
		event := <-deliveryChan
		msg := event.(*kafka.Message)
		if msg.TopicPartition.Error != nil {
			err = fmt.Errorf("Delivery failed: %v\n", msg.TopicPartition.Error)
			return err
		}
	}

	return nil
}

func newKafkaLogEntry(
	log *models.ObservationLogEntry,
) (keyBytes []byte, valueBytes []byte, err error) {
	// Create the Kafka key
	key := &models.ObservationLogKey{
		EventTimestamp: time.Now().Unix(),
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
