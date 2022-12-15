package logger

import "github.com/caraml-dev/timber/observation-service/types"

// NoopLogConsumer is the struct for no operation to ObservationLog
type NoopLogConsumer struct{}

// NewNoopLogConsumer initializes a NoopLogConsumer struct
func NewNoopLogConsumer() (*NoopLogConsumer, error) {
	return &NoopLogConsumer{}, nil
}

// Consume does nothing to ObservationLog
func (k *NoopLogConsumer) Consume(logsChannel chan *types.ObservationLogEntry) error {
	return nil
}
