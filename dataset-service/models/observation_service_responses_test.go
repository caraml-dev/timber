package models

import (
	"testing"

	"github.com/stretchr/testify/assert"

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
)

func TestObservationServiceResponseToApiSchema(t *testing.T) {
	id := "999"
	resp := &ObservationServiceResponse{Id: id}
	expected := &timberv1.ObservationServiceResponse{Id: id}

	assert.Equal(t, expected, resp.ToApiSchema())
}
