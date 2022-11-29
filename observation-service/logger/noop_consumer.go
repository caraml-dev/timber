package logger

import "github.com/caraml-dev/observation-service/observation-service/types"

type NoopLogConsumer struct{}

func NewNoopLogConsumer() (*NoopLogConsumer, error) {
	return &NoopLogConsumer{}, nil
}

func (k *NoopLogConsumer) Consume(logsChannel chan *types.ObservationLogEntry) error {
	return nil
}
