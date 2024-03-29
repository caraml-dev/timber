package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/caraml-dev/timber/observation-service/types"
)

func TestNoopLogConsumer(t *testing.T) {
	logConsumer, err := NewNoopLogConsumer()
	expected := &NoopLogConsumer{}

	assert.NoError(t, nil, err)
	assert.Equal(t, expected, logConsumer)

	observationLogEntryChannel := make(chan *types.ObservationLogEntry, 1)
	err = logConsumer.Consume(observationLogEntryChannel)
	assert.NoError(t, nil, err)
}
