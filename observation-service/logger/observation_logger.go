package logger

import (
	"context"
	"fmt"
	"time"

	"github.com/caraml-dev/timber/observation-service/config"
	"github.com/caraml-dev/timber/observation-service/services"
	"github.com/caraml-dev/timber/observation-service/types"
)

const (
	// SleepTime is the duration to sleep for prior to polling for new messages in ObservationLog channel
	SleepTime = time.Microsecond * 100
	// BaseNanoseconds is the base denomination for converting duration
	BaseNanoseconds = 1000000000
)

// LogConsumer captures the methods exposed by a ObservationLog consumer
type LogConsumer interface {
	Consume(logsChannel chan *types.ObservationLogEntry) error
}

// LogProducer captures the methods exposed by a ObservationLog producer
type LogProducer interface {
	Produce(log *types.ObservationLogEntry)
}

// ObservationLogger captures the config related to consume and produce ObservationLog to data sources and sinks respectively
type ObservationLogger struct {
	logsChannel   chan *types.ObservationLogEntry
	consumer      LogConsumer
	producer      LogProducer
	metricService services.MetricService
}

// Consume runs the Consume method of the underlying configured LogConsumer
func (l *ObservationLogger) Consume(ctx context.Context) error {
	err := l.consumer.Consume(l.logsChannel)
	if err != nil {
		return err
	}
	return nil
}

// worker is a goroutine that periodically calls Produce method
func (l *ObservationLogger) worker() {
	for {
		log := <-l.logsChannel
		l.producer.Produce(log)
	}
}

// NewObservationLogger initializes a ObservationLogger struct
func NewObservationLogger(
	consumerConfig config.LogConsumerConfig,
	producerConfig config.LogProducerConfig,
	metricService services.MetricService,
) (*ObservationLogger, error) {
	var err error
	var consumer LogConsumer
	// Instantiate Consumer
	switch consumerConfig.Kind {
	case config.LoggerNoopConsumer:
		consumer, err = NewNoopLogConsumer()
	case config.LoggerKafkaConsumer:
		consumer, err = NewKafkaLogConsumer(*consumerConfig.KafkaConfig, metricService)
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
		producer, err = NewKafkaLogProducer(*producerConfig.KafkaConfig, metricService)
	case config.LoggerFluentdProducer:
		producer, err = NewFluentdLogProducer(*producerConfig.FluentdConfig, metricService)
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
		metricService: metricService,
	}

	go logger.worker()

	return logger, nil
}
