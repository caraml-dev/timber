package logger

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/caraml-dev/observation-service/observation-service/config"
	"github.com/caraml-dev/observation-service/observation-service/services"
)

func TestObservationLogger(t *testing.T) {
	// Configs
	consumerConfig := config.LogConsumerConfig{}
	producerConfig := config.LogProducerConfig{}
	deploymentConfig := config.DeploymentConfig{}
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
	expected := &ObservationLogger{
		logsChannel:   observationLogger.logsChannel,
		consumer:      logConsumer,
		producer:      logProducer,
		metricService: metricService,
		flushInterval: time.Duration(producerConfig.FlushIntervalSeconds),
	}
	expected.batcherInfo = observationLogger.batcherInfo

	assert.NoError(t, nil, err)
	assert.Equal(t, expected, observationLogger)

	ctx := context.Background()
	err = expected.Consume(ctx)
	assert.NoError(t, nil, err)
}
