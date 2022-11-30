package logger

import (
	"fmt"
	"log"

	"github.com/fluent/fluent-logger-golang/fluent"

	"github.com/caraml-dev/observation-service/observation-service/config"
	"github.com/caraml-dev/observation-service/observation-service/types"
)

// FluentdLogProducer captures configs for publishing ObservationLog via Fluentd to a configured sink
type FluentdLogProducer struct {
	tag      string
	logger   *fluent.Fluent
	bqLogger *BigQueryLogger
}

// NewFluentdLogProducer creates a new FluentdLogProducer
func NewFluentdLogProducer(cfg config.FluentdConfig) (*FluentdLogProducer, error) {
	logger, err := fluent.New(fluent.Config{FluentPort: cfg.Port, FluentHost: cfg.Host})
	if err != nil {
		log.Println(err)
	}

	var bqLogger BigQueryLogger
	switch cfg.Kind {
	case config.LoggerBQSinkFluentdProducer:
		log.Println("Initiating BQ Sink configurations via Fluentd...")
		bqLogger, err = newBigQueryLogger(cfg.BQConfig)
		if err != nil {
			log.Println(err)
		}
	default:
		return nil, fmt.Errorf("invalid fluentd sink (%s) was provided", cfg.Kind)
	}

	return &FluentdLogProducer{
		tag:      cfg.Tag,
		logger:   logger,
		bqLogger: &bqLogger,
	}, nil
}

// Produce logs ObservationLog via Fluentd to the configured sink
func (p *FluentdLogProducer) Produce(logs []*types.ObservationLogEntry) error {
	for _, observationLog := range logs {
		logFormattedVal, err := observationLog.Value()
		if err != nil {
			log.Println(err)
		}
		err = p.logger.Post(p.tag, logFormattedVal)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}
