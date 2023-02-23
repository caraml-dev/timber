package appcontext

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/caraml-dev/timber/dataset-service/config"
	"github.com/caraml-dev/timber/dataset-service/mlp"
	"github.com/caraml-dev/timber/dataset-service/service"
	"github.com/caraml-dev/timber/dataset-service/storage"
)

// AppContext captures the config of all related internal services to run Dataset Service
type AppContext struct {
	Services service.Services
	DB       *gorm.DB
}

// NewAppContext initializes a AppContext struct
func NewAppContext(cfg *config.Config) (*AppContext, error) {
	// Init Services
	var allServices service.Services

	mlpSvc, err := mlp.NewClient(cfg.DatasetServiceConfig.MlpURL)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed initializing MLP Service")
	}

	obsSvc, err := service.NewObservationService(
		cfg.CommonDeploymentConfig,
		cfg.ObservationServiceConfig,
	)

	if err != nil {
		return nil, err
	}

	logWriterSvc, err := service.NewLogWriterService(
		cfg.CommonDeploymentConfig,
		cfg.LogWriterConfig,
	)

	if err != nil {
		return nil, err
	}

	allServices = service.NewServices(
		mlpSvc,
		obsSvc,
		logWriterSvc,
	)

	db, err := storage.InitDB(cfg.DatasetServiceConfig.DatabaseConfig)
	if err != nil {
		return nil, err
	}

	appContext := &AppContext{
		Services: allServices,
		DB:       db,
	}

	return appContext, nil
}
