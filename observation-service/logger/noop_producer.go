package logger

import "github.com/caraml-dev/observation-service/observation-service/types"

type NoopLogProducer struct{}

func NewNoopLogProducer() (*NoopLogProducer, error) {
	return &NoopLogProducer{}, nil
}

func (k *NoopLogProducer) Produce(log []*types.ObservationLogEntry) error {
	return nil
}
