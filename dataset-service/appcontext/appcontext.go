package appcontext

import (
	"github.com/caraml-dev/timber/dataset-service/config"
)

// AppContext captures the config of all related internal services to run Dataset Service
type AppContext struct{}

// NewAppContext initializes a AppContext struct
func NewAppContext(cfg *config.Config) (*AppContext, error) {
	appContext := &AppContext{}

	return appContext, nil
}
