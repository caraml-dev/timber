package types

import (
	"testing"
	"time"

	upiv1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestObservationLogKeyValue(t *testing.T) {
	logKey, uuidString := makeTestObservationLogKey(t)

	// Get loggable data and validate
	kvPairs, err := logKey.Value()
	require.NoError(t, err)
	assert.Equal(t, map[string]interface{}{
		"observation_batch_id": uuidString,
		"prediction_id":        "1",
		"row_id":               "1",
	}, kvPairs)
}

func TestObservationLogEntryValue(t *testing.T) {
	logEntry := makeTestObservationLogEntry(t)

	// Get loggable data and validate
	kvPairs, err := logEntry.Value()
	require.NoError(t, err)
	assert.Equal(t, map[string]interface{}{
		"observation_timestamp": "2000-02-01T04:05:06.000000007Z",
		"prediction_id":         "1",
		"row_id":                "1",
		"target_name":           "target_variable",
		"observation_values": []interface{}{
			map[string]interface{}{
				"name":         "variable1",
				"type":         "TYPE_STRING",
				"string_value": "variable_value",
			},
		},
		"observation_context": []interface{}{
			map[string]interface{}{
				"name":         "project",
				"type":         "TYPE_STRING",
				"string_value": "local",
			},
		},
	}, kvPairs)
}

// Helper methods for types package tests
func makeTestObservationLogKey(t *testing.T) (*ObservationLogKey, string) {
	// Create a ObservationLogKey record and add the data
	batchID := uuid.New().String()
	predictionID := "1"
	rowID := "1"

	key := &ObservationLogKey{
		ObservationBatchId: batchID,
		PredictionId:       predictionID,
		RowId:              rowID,
	}

	return key, batchID
}

func makeTestObservationLogEntry(t *testing.T) *ObservationLogEntry {

	// Create a ObservationLogEntry record and add the data
	predictionID := "1"
	rowID := "1"
	targetName := "target_variable"
	observationContext := []*upiv1.Variable{
		{
			Name:        "project",
			Type:        upiv1.Type_TYPE_STRING,
			StringValue: "local",
		},
	}
	observationValues := []*upiv1.Variable{
		{
			Name:        "variable1",
			Type:        upiv1.Type_TYPE_STRING,
			StringValue: "variable_value",
		},
	}
	timestamp := time.Date(2000, 2, 1, 4, 5, 6, 7, time.UTC)

	entry := &ObservationLogEntry{
		ObservationLog: upiv1.ObservationLog{
			PredictionId:         predictionID,
			RowId:                rowID,
			TargetName:           targetName,
			ObservationContext:   observationContext,
			ObservationValues:    observationValues,
			ObservationTimestamp: timestamppb.New(timestamp),
		},
		StartTime: time.Now(),
	}

	return entry
}
