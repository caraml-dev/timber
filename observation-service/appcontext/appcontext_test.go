package appcontext

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/caraml-dev/timber/observation-service/config"
)

func TestNewAppContext(t *testing.T) {
	cfg := &config.Config{}
	appCtx, err := NewAppContext(cfg)

	expectedAppContext := &AppContext{
		ObservationLogger: appCtx.ObservationLogger,
		MetricService:     appCtx.MetricService,
	}

	require.NoError(t, err)
	assert.Equal(t, expectedAppContext, appCtx)
}
