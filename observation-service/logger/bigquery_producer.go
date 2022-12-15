package logger

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"cloud.google.com/go/bigquery"
	"go.einride.tech/protobuf-bigquery/encoding/protobq"

	"github.com/caraml-dev/timber/observation-service/config"
	"github.com/caraml-dev/timber/observation-service/log"
	"github.com/caraml-dev/timber/observation-service/types"
	upiv1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
)

// BigQueryLogger defines methods implemented by the logger
type BigQueryLogger interface {
	getLogData(*types.ObservationLogEntry) interface{}
}

// bigQueryLogger implements the BigQueryLogger interface and wraps the bigquery.Client
// and other necessary information to save the data to BigQuery
type bigQueryLogger struct {
	dataset  string
	table    string
	bqClient *bigquery.Client
	schema   *bigquery.Schema
}

// newBigQueryLogger creates a new BigQueryLogger
func newBigQueryLogger(cfg *config.BQConfig) (BigQueryLogger, error) {
	ctx := context.Background()
	bqClient, err := bigquery.NewClient(ctx, cfg.Project)
	if err != nil {
		return nil, err
		// return nil, errors.Wrapf(err, "Failed to initialize BigQuery Client")
	}
	// Create the BigQuery logger
	bqLogger := &bigQueryLogger{
		dataset:  cfg.Dataset,
		table:    cfg.Table,
		bqClient: bqClient,
		schema:   getObservationLogTableSchema(),
	}
	// Set up Observation Log table
	err = bqLogger.setUpObservationLogTable()
	if err != nil {
		return nil, err
	}
	return bqLogger, nil
}

// getLogData returns the log information as a generic interface{} object. Internally, it calls
// the Save method defined on the BqLogEntry structure which implements the
// bigquery.ValueSaver interface and returns the log data as a map. This can be returned
// as is for logging by other loggers whose destination is a BQ table.
func (l *bigQueryLogger) getLogData(obsLogEntry *types.ObservationLogEntry) interface{} {
	entry := &types.BqLogEntry{ObservationLogEntry: obsLogEntry}
	record, _, err := entry.Save()
	if err != nil {
		log.Errorf("failed to create log entry %s", err)
	}
	return record
}

// setUpObservationLogTable checks that the logging table is set up in BQ as expected.
// If the specified dataset does not exist in the project, it returns an error.
// If the dataset + table exists and the schema does not match the expected,
// an error is returned as well. If the dataset exists but not the table, a
// new table is created.
func (l *bigQueryLogger) setUpObservationLogTable() error {
	ctx := context.Background()

	// Check that the dataset exists
	dataset := l.bqClient.Dataset(l.dataset)
	_, err := dataset.Metadata(ctx)
	if err != nil {
		// return errors.Wrapf(err, "BigQuery dataset %s not found", l.dataset)
		return err
	}

	// Check if the table exists
	table := dataset.Table(l.table)
	metadata, err := table.Metadata(ctx)

	// If not, create
	if err != nil {
		err = createObservationLogTable(&ctx, table, l.schema)
		if err != nil {
			// return errors.Wrapf(err, "Failed creating BigQuery table %s", l.table)
			return err
		}
	} else {
		// Table exists, compare schema
		schema, isUpdated, err := compareTableSchema(&metadata.Schema, l.schema)
		if err != nil {
			// return errors.Wrapf(err, "Unexpected schema for BigQuery table %s", l.table)
			return err
		}
		// Update schema, if it changed
		if isUpdated {
			update := bigquery.TableMetadataToUpdate{
				Schema: *schema,
			}
			if _, err := table.Update(ctx, update, metadata.ETag); err != nil {
				return err
			}
		} else {
			// No update to schema required, check that we have the required perms
			// for data write
			return checkBQTableWritePermissions(table)
		}
	}

	return nil
}

// checkPermissions checks that the BQ client has the required permissions on the dataset
func checkBQTableWritePermissions(table *bigquery.Table) error {
	// Ref: https://cloud.google.com/bigquery/docs/access-control
	requiredPerms := []string{
		"bigquery.tables.get",
		"bigquery.tables.getData",
		"bigquery.tables.update",
	}
	perms, err := table.IAM().TestPermissions(context.Background(), requiredPerms)
	if err != nil {
		return err
	}
	if len(perms) < len(requiredPerms) {
		return fmt.Errorf("Insufficient permissions. Got: %s; Want: %s",
			strings.Join(perms, ","),
			strings.Join(requiredPerms, ","))
	}
	return nil
}

// createObservationLogTable creates the specified table if not exists
func createObservationLogTable(
	ctx *context.Context,
	table *bigquery.Table,
	schema *bigquery.Schema,
) error {
	// Set partitioning
	metaData := &bigquery.TableMetadata{
		Schema: *schema,
		TimePartitioning: &bigquery.TimePartitioning{
			Field:                  "observation_timestamp",
			RequirePartitionFilter: false,
		},
	}

	// Create the table
	if err := table.Create(*ctx, metaData); err != nil {
		return err
	}

	return nil
}

// compareTableSchema validates the important properties of each field in the schema
// recursively. If the expected schema has more columns than the actual, and these
// columns are nullable, they are added to the actual schema and the 'updated' flag
// is set to true. The (updated) actual schema, the flag and any error is returned.
func compareTableSchema(
	tableSchema *bigquery.Schema,
	expectedSchema *bigquery.Schema,
) (*bigquery.Schema, bool, error) {
	isUpdated := false

	// Create a map of the tableSchema column name to the field
	tableSchemaMap := map[string]*bigquery.FieldSchema{}
	for _, item := range *tableSchema {
		tableSchemaMap[item.Name] = item
	}

	// For each field in the expected schema, add it to the tableSchema if it
	// doesn't already exist. Compare the properties otherwise.
	for _, ef := range *expectedSchema {
		var af *bigquery.FieldSchema
		var ok bool
		if af, ok = tableSchemaMap[ef.Name]; ok {
			// Compare current schema
			if af.Name != ef.Name ||
				af.Type != ef.Type ||
				af.Required != ef.Required ||
				af.Repeated != ef.Repeated {
				return tableSchema, false, fmt.Errorf(
					"BigQuery schema mismatch for field %s", ef.Name,
				)
				// return tableSchema, false, errors.Newf(errors.BadConfig,
				// 	"BigQuery schema mismatch for field %s", ef.Name)
			}
			// Compare nested schema
			nestedSchema, itemUpdated, err := compareTableSchema(&af.Schema, &ef.Schema)
			// If error, return
			if err != nil {
				return tableSchema, false, err
			}
			// Save the (new) nested schema to the current field
			af.Schema = *nestedSchema
			// Set the overall updated flag
			isUpdated = isUpdated || itemUpdated
		} else if !ef.Required {
			// Append NULLABLE missing field to the tableSchema
			*tableSchema = append(*tableSchema, ef)
			isUpdated = true
		} else {
			// Return error
			return tableSchema, false, fmt.Errorf(
				"Cannot add Required field %s to the existing BQ table", ef.Name,
			)
		}
	}

	// At the end of any additions, the two schemas must have the same number of fields
	if len(*tableSchema) != len(*expectedSchema) {
		return tableSchema, false, errors.New("BigQuery schema mismatch")
	}

	return tableSchema, isUpdated, nil
}

// getObservationLogTableSchema returns the expected schema defined for logging results
// to BigQuery
func getObservationLogTableSchema() *bigquery.Schema {
	schema := protobq.InferSchema(&upiv1.ObservationLog{})
	return &schema
}
