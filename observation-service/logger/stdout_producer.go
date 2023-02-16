package logger

import (
	"github.com/caraml-dev/timber/common/log"
	"github.com/caraml-dev/timber/observation-service/types"
)

// StdOutLogProducer captures configs for logging ObservationLog to standard output
type StdOutLogProducer struct{}

// NewStdOutLogProducer initializes a StdOutLogProducer struct
func NewStdOutLogProducer() (*StdOutLogProducer, error) {
	return &StdOutLogProducer{}, nil
}

// Produce logs ObservationLog to standard output
func (p *StdOutLogProducer) Produce(observationLog *types.ObservationLogEntry) {
	log.Info(observationLog)
}
