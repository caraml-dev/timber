package appcontext

import (
	"log"

	"github.com/caraml-dev/observation-service/observation-service/config"
	"github.com/caraml-dev/observation-service/observation-service/logger"
)

type AppContext struct {
	ObservationLogger *logger.ObservationLogger
}

func NewAppContext(cfg *config.Config) (*AppContext, error) {
	log.Println("Initializing Observation Service logger...")
	var observationLogger *logger.ObservationLogger
	observationLogger, err := logger.NewObservationLogger(
		cfg.LogConsumerConfig,
		cfg.LogProducerConfig,
	)
	if err != nil {
		return nil, err
	}

	appContext := &AppContext{
		ObservationLogger: observationLogger,
	}

	return appContext, nil
}
