package appcontext

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/caraml-dev/timber/dataset-service/services"
	"github.com/caraml-dev/timber/dataset-service/services/mocks"
)

func TestNewAppContext(t *testing.T) {
	// Create mock MLP service
	mlpSvc := &mocks.MLPService{}
	appCtx := AppContext{
		Services: services.Services{
			MLPService: mlpSvc,
		},
	}
	assert.Equal(t, mlpSvc, appCtx.Services.MLPService)
}
