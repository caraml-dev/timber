package logger

import (
	"context"
	"fmt"
	"time"

	"github.com/caraml-dev/observation-service/observation-service/config"
	upiv1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
)

type LogConsumer interface {
	Consume(queueChannel chan *upiv1.ObservationLog) error
}

type ObservationLogger struct {
	queue    chan *upiv1.ObservationLog
	consumer LogConsumer

	flushInterval time.Duration
}

func (l *ObservationLogger) Consume(ctx context.Context) error {
	// TODO: Pass queueLength from config
	queueLength := 100
	queueChannel := make(chan *upiv1.ObservationLog, queueLength)

	err := l.consumer.Consume(queueChannel)
	if err != nil {
		return err
	}
	return nil
}

func (l *ObservationLogger) Append(log *upiv1.ObservationLog) error {
	l.queue <- log
	return nil
}

func (l *ObservationLogger) worker() {
	for range time.Tick(l.flushInterval) {
		logs := make([]*upiv1.ObservationLog, 0)

	collection:
		for {
			select {
			case log := <-l.queue:
				logs = append(logs, log)
			default:
				break collection
			}
		}

		if len(logs) > 0 {
			// TODO: Implement publishing of Observation logs
			fmt.Println(logs)
		}
	}
}

func NewNoopObservationLogger() (*ObservationLogger, error) {
	return nil, nil
}

func NewObservationLogger(
	consumerConfig config.LogConsumerConfig,
	queueLength int,
	flushInterval time.Duration,
) (*ObservationLogger, error) {
	c := make(chan *upiv1.ObservationLog, queueLength)

	var err error
	var consumer LogConsumer
	// Instantiate Consumer
	switch consumerConfig.Kind {
	case config.LoggerKafkaConsumer:
		consumer, err = NewKafkaLogConsumer(
			consumerConfig.KafkaConsumerConfig.Brokers,
			consumerConfig.KafkaConsumerConfig.Topic,
			consumerConfig.KafkaConsumerConfig.ConnectTimeoutMS,
			c,
		)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("invalid consumer (%s) was provided", consumerConfig.Kind)
	}

	// TODO: Instantiate Producer

	logger := &ObservationLogger{
		queue:         c,
		consumer:      consumer,
		flushInterval: flushInterval,
	}

	// go logger.worker()

	return logger, nil
}
