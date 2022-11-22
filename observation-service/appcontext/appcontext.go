package appcontext

import "github.com/caraml-dev/observation-service/observation-service/config"

type AppContext struct{}

func NewAppContext(cfg *config.Config) (*AppContext, error) {
	appContext := &AppContext{}

	return appContext, nil
}
