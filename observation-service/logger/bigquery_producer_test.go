package logger

import (
	"fmt"
	"path/filepath"
	"testing"
	"time"

	"cloud.google.com/go/bigquery"
	upiv1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"

	tu "github.com/caraml-dev/observation-service/observation-service/internal/testutils"
	"github.com/caraml-dev/observation-service/observation-service/types"
)

type testSuiteBQSchema struct {
	filepath1 string
	filepath2 string
	isUpdated bool
	isError   bool
}

func TestGetObservationLogTableSchema(t *testing.T) {
	// Get expected schema
	bytes, err := tu.ReadFile(filepath.Join("testdata", "bq_observation_log_schema.json"))
	assert.NoError(t, err)
	expectedSchema, err := bigquery.SchemaFromJSON(bytes)
	assert.NoError(t, err)

	// Actual schema
	schema := getObservationLogTableSchema()

	// Enclose schema in a struct for go-cmp
	type bqSchema struct {
		Schema bigquery.Schema
	}
	wantSchema := &bqSchema{Schema: expectedSchema}
	gotSchema := &bqSchema{Schema: *schema}

	// Compare all fields except Description
	opt := cmpopts.IgnoreFields(bigquery.FieldSchema{}, "Description")
	if !cmp.Equal(wantSchema, gotSchema, opt) {
		t.Log(cmp.Diff(wantSchema, gotSchema, opt))
		t.Fail()
	}
}

func TestCheckTableSchema(t *testing.T) {
	tests := map[string]testSuiteBQSchema{
		"order_diff": {
			filepath1: filepath.Join("testdata", "bq_schema_1_order_diff.json"),
			filepath2: filepath.Join("testdata", "bq_schema_1_original.json"),
			isError:   false,
		},
		"field_diff": {
			filepath1: filepath.Join("testdata", "bq_schema_2_field_diff.json"),
			filepath2: filepath.Join("testdata", "bq_schema_1_original.json"),
			isError:   true,
		},
		"required_diff": {
			filepath1: filepath.Join("testdata", "bq_schema_3_required_diff.json"),
			filepath2: filepath.Join("testdata", "bq_schema_1_original.json"),
			isError:   true,
		},
		"nested_schema_diff": {
			filepath1: filepath.Join("testdata", "bq_schema_4_nested_schema_diff.json"),
			filepath2: filepath.Join("testdata", "bq_schema_1_original.json"),
			isUpdated: true,
			isError:   false,
		},
	}

	for name, data := range tests {
		t.Run(name, func(t *testing.T) {
			// Read in the JSON schema from the two files
			filebytes1, _ := tu.ReadFile(data.filepath1)
			filebytes2, _ := tu.ReadFile(data.filepath2)

			// Create BQ schema
			schema1, _ := bigquery.SchemaFromJSON(filebytes1)
			schema2, _ := bigquery.SchemaFromJSON(filebytes2)

			// Compare and check the success state
			newSchema, isUpdated, err := compareTableSchema(&schema1, &schema2)
			assert.Equal(t, data.isError, err != nil)
			assert.Equal(t, data.isUpdated, isUpdated)
			// If updated, check that the new schema and the expected schema match
			if isUpdated {
				_, isUpdated, err = compareTableSchema(&schema1, newSchema)
				assert.NoError(t, err)
				assert.False(t, isUpdated)
			}
		})
	}
}

func TestBigQueryLoggerGetData(t *testing.T) {
	// Create new logger
	testLogger := &bigQueryLogger{}

	// Generate record
	predictionId := "pred-1"
	rowId := "row-1"
	targetName := "target-name"
	timestamp := &timestamppb.Timestamp{Seconds: time.Date(2000, 2, 1, 4, 5, 6, 7, time.UTC).Unix()}
	message := &upiv1.ObservationLog{
		PredictionId: predictionId,
		RowId:        rowId,
		TargetName:   targetName,
		ObservationValues: []*upiv1.Variable{
			{
				Name:        "variable1",
				Type:        upiv1.Type_TYPE_STRING,
				StringValue: "variable_value",
			},
		},
		ObservationContext: []*upiv1.Variable{
			{
				Name:        "project",
				Type:        upiv1.Type_TYPE_STRING,
				StringValue: "local",
			},
		},
		ObservationTimestamp: timestamp,
	}

	// Create a ObservationLogEntry record and add the data
	entry := types.NewObservationLogEntry(message)

	// Get the log data and validate
	logData := testLogger.getLogData(entry)
	// Cast to map[string]bigquery.Value
	if logMap, ok := logData.(map[string]bigquery.Value); ok {
		fmt.Println(logMap)
		assert.Equal(t, predictionId, logMap["prediction_id"])
		assert.Equal(t, rowId, logMap["row_id"])
		assert.Equal(t, targetName, logMap["target_name"])
		assert.Equal(t, []interface{}{
			map[string]interface{}{
				"name":         "variable1",
				"type":         "TYPE_STRING",
				"string_value": "variable_value",
			},
		}, logMap["observation_values"])
		assert.Equal(t, []interface{}{
			map[string]interface{}{
				"name":         "project",
				"type":         "TYPE_STRING",
				"string_value": "local",
			},
		}, logMap["observation_context"])
		assert.Equal(t, "2000-02-01T04:05:06Z", logMap["observation_timestamp"])
	}
}
