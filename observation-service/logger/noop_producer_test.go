package logger

import (
	"testing"

	"github.com/caraml-dev/observation-service/observation-service/types"
	"github.com/stretchr/testify/assert"
)

func TestNoopLogProducer(t *testing.T) {
	logProducer, err := NewNoopLogProducer()
	expected := &NoopLogProducer{}

	assert.NoError(t, nil, err)
	assert.Equal(t, expected, logProducer)

	observationLogs := []*types.ObservationLogEntry{}
	err = logProducer.Produce(observationLogs)
	assert.NoError(t, nil, err)
}
