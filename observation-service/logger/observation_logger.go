package logger

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/caraml-dev/observation-service/observation-service/config"
	"github.com/caraml-dev/observation-service/observation-service/types"
)

type LogConsumer interface {
	Consume(logsChannel chan *types.ObservationLogEntry) error
}

type LogProducer interface {
	Produce(log []*types.ObservationLogEntry) error
}

type ObservationLogger struct {
	logsChannel chan *types.ObservationLogEntry
	consumer    LogConsumer
	producer    LogProducer

	flushInterval time.Duration
}

func (l *ObservationLogger) Consume(ctx context.Context) error {
	err := l.consumer.Consume(l.logsChannel)
	if err != nil {
		return err
	}
	return nil
}

// worker is a goroutine that periodically calls Produce method
func (l *ObservationLogger) worker() {
	for range time.Tick(l.flushInterval) {
		logs := make([]*types.ObservationLogEntry, 0)

	collection:
		for {
			select {
			case log := <-l.logsChannel:
				logs = append(logs, log)
			default:
				break collection
			}
		}

		if len(logs) > 0 {
			err := l.producer.Produce(logs)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

type NoopLogConsumer struct{}

func NewNoopLogConsumer() (*NoopLogConsumer, error) {
	return &NoopLogConsumer{}, nil
}

func (k *NoopLogConsumer) Consume(logsChannel chan *types.ObservationLogEntry) error {
	return nil
}

type NoopLogProducer struct{}

func NewNoopLogProducer() (*NoopLogProducer, error) {
	return &NoopLogProducer{}, nil
}

func (k *NoopLogProducer) Produce(log []*types.ObservationLogEntry) error {
	return nil
}

func NewObservationLogger(
	consumerConfig config.LogConsumerConfig,
	producerConfig config.LogProducerConfig,
) (*ObservationLogger, error) {
	var err error
	var consumer LogConsumer
	// Instantiate Consumer
	switch consumerConfig.Kind {
	case config.LoggerNoopConsumer:
		consumer, err = NewNoopLogConsumer()
	case config.LoggerKafkaConsumer:
		consumer, err = NewKafkaLogConsumer(*consumerConfig.KafkaConsumerConfig)
	default:
		return nil, fmt.Errorf("invalid consumer (%s) was provided", consumerConfig.Kind)
	}
	if err != nil {
		return nil, err
	}

	// Instantiate Producer
	var producer LogProducer
	c := make(chan *types.ObservationLogEntry, producerConfig.QueueLength)
	switch producerConfig.Kind {
	case config.LoggerNoopProducer:
		producer, err = NewNoopLogProducer()
	case config.LoggerStdOutProducer:
		producer, err = NewStdOutLogProducer()
	case config.LoggerKafkaProducer:
		producer, err = NewKafkaLogProducer(*producerConfig.KafkaProducerConfig)
	default:
		return nil, fmt.Errorf("invalid producer (%s) was provided", producerConfig.Kind)
	}
	if err != nil {
		return nil, err
	}

	logger := &ObservationLogger{
		logsChannel:   c,
		consumer:      consumer,
		producer:      producer,
		flushInterval: time.Duration(producerConfig.FlushIntervalSeconds),
	}

	go logger.worker()

	return logger, nil
}
