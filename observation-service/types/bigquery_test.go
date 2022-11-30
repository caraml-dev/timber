package types

import (
	"testing"

	"cloud.google.com/go/bigquery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBQLogEntryValue(t *testing.T) {
	observationLogEntry := makeTestObservationLogEntry(t)
	bqLogEntry := BqLogEntry{observationLogEntry}

	// Get loggable data and validate
	kvPairs, str, err := bqLogEntry.Save()
	require.NoError(t, err)
	assert.Equal(t, str, "")
	assert.Equal(t, map[string]bigquery.Value{
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
