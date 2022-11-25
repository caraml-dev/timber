package logger

import (
	"log"

	"github.com/caraml-dev/observation-service/observation-service/types"
)

type StdOutLogProducer struct{}

func NewStdOutLogProducer() (*StdOutLogProducer, error) {
	return &StdOutLogProducer{}, nil
}

func (p *StdOutLogProducer) Produce(logs []*types.ObservationLogEntry) error {
	for _, observationLog := range logs {
		log.Println(observationLog)
	}
	return nil
}
