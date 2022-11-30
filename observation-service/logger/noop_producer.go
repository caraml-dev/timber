package logger

import "github.com/caraml-dev/observation-service/observation-service/types"

// NoopLogProducer is the struct for no operation to ObservationLog
type NoopLogProducer struct{}

// NewNoopLogProducer initializes a NoopLogProducer struct
func NewNoopLogProducer() (*NoopLogProducer, error) {
	return &NoopLogProducer{}, nil
}

// Produce does nothing to ObservationLog
func (k *NoopLogProducer) Produce(log []*types.ObservationLogEntry) error {
	return nil
}
