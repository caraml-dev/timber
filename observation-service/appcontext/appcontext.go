package appcontext

import (
	"github.com/caraml-dev/observation-service/observation-service/config"
	"github.com/caraml-dev/observation-service/observation-service/log"
	"github.com/caraml-dev/observation-service/observation-service/logger"
	"github.com/caraml-dev/observation-service/observation-service/services"
)

// AppContext captures the config of all related internal services to run Observation Service
type AppContext struct {
	ObservationLogger *logger.ObservationLogger

	MetricService services.MetricService
}

// NewAppContext initializes a AppContext struct
func NewAppContext(cfg *config.Config) (*AppContext, error) {
	log.Glob().Info("Initializing metric service...")
	metricService, err := services.NewMetricService(cfg.DeploymentConfig, cfg.MonitoringConfig)
	if err != nil {
		return nil, err
	}

	log.Glob().Info("Initializing Observation Service logger...")
	var observationLogger *logger.ObservationLogger
	observationLogger, err = logger.NewObservationLogger(
		cfg.LogConsumerConfig,
		cfg.LogProducerConfig,
		metricService,
	)
	if err != nil {
		return nil, err
	}

	appContext := &AppContext{
		ObservationLogger: observationLogger,
		MetricService:     metricService,
	}

	return appContext, nil
}
