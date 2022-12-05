package logger

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/caraml-dev/observation-service/observation-service/config"
	"github.com/caraml-dev/observation-service/observation-service/log"
	"github.com/caraml-dev/observation-service/observation-service/monitoring"
	"github.com/caraml-dev/observation-service/observation-service/services"
	"github.com/caraml-dev/observation-service/observation-service/types"
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
	Produce(log []*types.ObservationLogEntry)
}

type batcherInfo struct {
	Records         []*types.ObservationLogEntry
	Start           time.Time
	Now             time.Time
	CurrentInputLen int
}

func getCurrentTime() time.Time {
	return time.Now().UTC()
}

func (batcherInfo *batcherInfo) InitializeInfo() {
	batcherInfo.CurrentInputLen = 0
	batcherInfo.Records = make([]*types.ObservationLogEntry, 0)
	batcherInfo.Start = getCurrentTime()
	batcherInfo.Now = batcherInfo.Start
}

// ObservationLogger captures the config related to consume and produce ObservationLog to data sources and sinks respectively
type ObservationLogger struct {
	logsChannel   chan *types.ObservationLogEntry
	consumer      LogConsumer
	producer      LogProducer
	metricService services.MetricService

	batcherInfo   batcherInfo
	flushInterval time.Duration
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
	log.Infof("starting periodic flush: Max Flush Duration %s / Max Queue Size %d", l.flushInterval, cap(l.logsChannel))
	for {
		select {
		case log := <-l.logsChannel:
			if len(l.batcherInfo.Records) == 0 {
				l.batcherInfo.Start = getCurrentTime()
			}
			l.batcherInfo.Records = append(l.batcherInfo.Records, log)
			l.batcherInfo.CurrentInputLen = len(l.batcherInfo.Records)
		case <-time.After(SleepTime):
		}
		l.batcherInfo.Now = getCurrentTime()
		// Flushing should be either time-based or if x no. of messages have been reached
		if l.batcherInfo.CurrentInputLen >= cap(l.logsChannel) ||
			(l.batcherInfo.Now.Sub(l.batcherInfo.Start).Seconds() >= l.flushInterval.Seconds() &&
				l.batcherInfo.CurrentInputLen > 0) {
			l.producer.Produce(l.batcherInfo.Records)
			// Reset batcherInfo after flush
			l.batcherInfo.InitializeInfo()
			// Increment Flush count
			l.metricService.LogRequestCount(http.StatusOK, monitoring.FlushCount)
		}
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
		flushInterval: time.Duration(producerConfig.FlushIntervalSeconds * BaseNanoseconds),
	}
	logger.batcherInfo.InitializeInfo()

	go logger.worker()

	return logger, nil
}
