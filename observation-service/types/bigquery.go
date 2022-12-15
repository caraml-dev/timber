package types

import (
	"encoding/json"

	"cloud.google.com/go/bigquery"
	"github.com/caraml-dev/timber/observation-service/errors"
)

// BqLogEntry wraps a ObservationLogEntry and implements the bigquery.ValueSaver interface
type BqLogEntry struct {
	*ObservationLogEntry
}

// Save implements the ValueSaver interface on BqLogEntry, for saving the data to BigQuery
func (e *BqLogEntry) Save() (map[string]bigquery.Value, string, error) {
	var kvPairs map[string]bigquery.Value
	bytes, err := json.Marshal(e)
	if err != nil {
		return kvPairs, "", err
	}

	// Unmarshal into map[string]bigquery.Value
	err = json.Unmarshal(bytes, &kvPairs)
	if err != nil {
		return kvPairs, "", errors.Wrapf(err, "Error unmarshaling the result log for save to BQ")
	}

	return kvPairs, "", nil
}
