package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStdOutLogProducer(t *testing.T) {
	logProducer, err := NewStdOutLogProducer()
	expected := &StdOutLogProducer{}

	assert.NoError(t, nil, err)
	assert.Equal(t, expected, logProducer)
}
