package logger

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	common_config "github.com/caraml-dev/timber/common/config"
	"github.com/caraml-dev/timber/observation-service/config"
	"github.com/caraml-dev/timber/observation-service/services"
)

func TestObservationLogger(t *testing.T) {
	// Configs
	consumerConfig := config.LogConsumerConfig{}
	producerConfig := config.LogProducerConfig{}
	deploymentConfig := common_config.DeploymentConfig{}
	metricConfig := config.MonitoringConfig{}
	metricService, err := services.NewMetricService(deploymentConfig, metricConfig)
	assert.NoError(t, nil, err)

	observationLogger, err := NewObservationLogger(
		consumerConfig,
		producerConfig,
		metricService,
	)
	assert.NoError(t, nil, err)

	// Expected
	logConsumer, err := NewNoopLogConsumer()
	assert.NoError(t, nil, err)
	logProducer, err := NewNoopLogProducer()
	assert.NoError(t, nil, err)

	assert.Equal(t, logConsumer, observationLogger.consumer)
	assert.Equal(t, logProducer, observationLogger.producer)
	assert.Equal(t, metricService, observationLogger.metricService)

	ctx := context.Background()
	err = observationLogger.Consume(ctx)
	assert.NoError(t, nil, err)
}
