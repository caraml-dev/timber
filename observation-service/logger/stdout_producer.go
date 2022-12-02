package logger

import (
	"github.com/caraml-dev/observation-service/observation-service/log"
	"github.com/caraml-dev/observation-service/observation-service/types"
)

// StdOutLogProducer captures configs for logging ObservationLog to standard output
type StdOutLogProducer struct{}

// NewStdOutLogProducer initializes a StdOutLogProducer struct
func NewStdOutLogProducer() (*StdOutLogProducer, error) {
	return &StdOutLogProducer{}, nil
}

// Produce logs ObservationLog to standard output
func (p *StdOutLogProducer) Produce(logs []*types.ObservationLogEntry) error {
	for _, observationLog := range logs {
		log.Info(observationLog)
	}
	return nil
}
