package logger

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fluent/fluent-logger-golang/fluent"

	"github.com/caraml-dev/timber/common/log"
	"github.com/caraml-dev/timber/observation-service/config"
	"github.com/caraml-dev/timber/observation-service/monitoring"
	"github.com/caraml-dev/timber/observation-service/services"
	"github.com/caraml-dev/timber/observation-service/types"
)

// FluentdLogProducer captures configs for publishing ObservationLog via Fluentd to a configured sink
type FluentdLogProducer struct {
	tag            string
	logger         *fluent.Fluent
	bqLogger       *BigQueryLogger
	metricsService services.MetricService
}

// NewFluentdLogProducer creates a new FluentdLogProducer
func NewFluentdLogProducer(
	cfg config.FluentdConfig,
	metricsService services.MetricService,
) (*FluentdLogProducer, error) {
	logger, err := fluent.New(
		fluent.Config{
			FluentPort:             cfg.Port,
			FluentHost:             cfg.Host,
			Async:                  true,
			AsyncReconnectInterval: 10000,
			BufferLimit:            cfg.BufferLimit,
		},
	)
	if err != nil {
		log.Error(err)
	}

	var bqLogger BigQueryLogger
	switch cfg.Kind {
	case config.LoggerBQSinkFluentdProducer:
		log.Info("Initiating BQ Sink configurations via Fluentd...")
		bqLogger, err = newBigQueryLogger(cfg.BQConfig)
		if err != nil {
			log.Error(err)
		}
	default:
		return nil, fmt.Errorf("invalid fluentd sink (%s) was provided", cfg.Kind)
	}

	return &FluentdLogProducer{
		tag:            cfg.Tag,
		logger:         logger,
		bqLogger:       &bqLogger,
		metricsService: metricsService,
	}, nil
}

// Produce logs ObservationLog via Fluentd to the configured sink
func (p *FluentdLogProducer) Produce(observationLog *types.ObservationLogEntry) {
	logFormattedVal, err := observationLog.Value()
	if err != nil {
		// TODO: Send failed ObservationLog to deadletter sink
		p.metricsService.LogRequestCount(http.StatusInternalServerError, monitoring.FlushObservationCount)
		log.Error(err)
	}
	labels := map[string]string{
		"component": "fluentd",
	}
	fluentdFlushStartTime := time.Now()
	err = p.logger.Post(p.tag, logFormattedVal)
	if err != nil {
		p.metricsService.LogRequestCount(http.StatusInternalServerError, monitoring.FlushObservationCount)
		log.Error(err)
	} else {
		// Log fluentd latency
		p.metricsService.LogLatencyHistogram(fluentdFlushStartTime, http.StatusOK, monitoring.FlushDurationMs, labels)
		p.metricsService.LogRequestCount(http.StatusOK, monitoring.FlushObservationCount)
		// Log E2E latency
		labels = map[string]string{
			"component": "e2e",
		}
		p.metricsService.LogLatencyHistogram(observationLog.StartTime, http.StatusOK, monitoring.FlushDurationMs, labels)
	}
}
