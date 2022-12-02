package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNoopLogProducer(t *testing.T) {
	logProducer, err := NewNoopLogProducer()
	expected := &NoopLogProducer{}

	assert.NoError(t, nil, err)
	assert.Equal(t, expected, logProducer)
}
