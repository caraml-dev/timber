package appcontext

import (
	"github.com/pkg/errors"

	"github.com/caraml-dev/timber/dataset-service/config"
	"github.com/caraml-dev/timber/dataset-service/services"
)

// AppContext captures the config of all related internal services to run Dataset Service
type AppContext struct {
	Services services.Services
}

// NewAppContext initializes a AppContext struct
func NewAppContext(cfg *config.Config) (*AppContext, error) {
	mlpSvc, err := services.NewMLPService(cfg.MLPConfig.URL)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed initializing MLP Service")
	}

	allServices := services.NewServices(
		mlpSvc,
	)

	appContext := &AppContext{
		Services: allServices,
	}

	return appContext, nil
}
