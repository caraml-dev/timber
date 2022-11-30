package types

import (
	"encoding/json"

	upiv1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
	"github.com/google/uuid"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/caraml-dev/observation-service/observation-service/errors"
)

// ObservationLogKey is an alias for upiv1.ObservationLogKey proto, to support extension of default
// methods such as Value, to consolidate conversions required to write to different sinks
type ObservationLogKey upiv1.ObservationLogKey

// NewObservationLogKey initializes a ObservationLogKey struct
func NewObservationLogKey(rawObservationKey *upiv1.ObservationLogKey) *ObservationLogKey {
	return &ObservationLogKey{
		ObservationBatchId: uuid.New().String(),
		PredictionId:       rawObservationKey.GetPredictionId(),
		RowId:              rawObservationKey.GetRowId(),
	}
}

// Value returns the NewObservationLogKey in a loggable format
func (logEntry *ObservationLogKey) Value() (map[string]interface{}, error) {
	var kvPairs map[string]interface{}
	// Marshal into bytes
	bytes, err := json.Marshal(&logEntry)
	if err != nil {
		return kvPairs, errors.Wrapf(err, "Error marshaling the observation log key")
	}
	// Unmarshal into map[string]interface{}
	err = json.Unmarshal(bytes, &kvPairs)
	if err != nil {
		return kvPairs, errors.Wrapf(err, "Error unmarshaling the observation log key")
	}
	return kvPairs, nil
}

// ObservationLogEntry is an alias for upiv1.ObservationLog proto, to support extension of default
// methods such as MarshalJSON and Value, to consolidate conversions required to write to different sinks
type ObservationLogEntry upiv1.ObservationLog

// MarshalJSON implements custom Marshaling for ObservationLogEntry, using the underlying proto def
func (logEntry *ObservationLogEntry) MarshalJSON() ([]byte, error) {
	m := &protojson.MarshalOptions{
		UseProtoNames: true, // Use the json field name instead of the camel case struct field name
	}
	message := (*upiv1.ObservationLog)(logEntry)
	return m.Marshal(message)
}

// Value returns the ObservationLogEntry in a loggable format
func (logEntry *ObservationLogEntry) Value() (map[string]interface{}, error) {
	var kvPairs map[string]interface{}
	// Marshal into bytes
	bytes, err := json.Marshal(&logEntry)
	if err != nil {
		return kvPairs, errors.Wrapf(err, "Error marshaling the observation log entry")
	}
	// Unmarshal into map[string]interface{}
	err = json.Unmarshal(bytes, &kvPairs)
	if err != nil {
		return kvPairs, errors.Wrapf(err, "Error unmarshaling the observation log entry")
	}
	return kvPairs, nil
}

// NewObservationLogEntry initializes a ObservationLogEntry struct
func NewObservationLogEntry(rawObservation *upiv1.ObservationLog) *ObservationLogEntry {
	return &ObservationLogEntry{
		PredictionId:         rawObservation.GetPredictionId(),
		RowId:                rawObservation.GetRowId(),
		TargetName:           rawObservation.GetTargetName(),
		ObservationContext:   rawObservation.GetObservationContext(),
		ObservationValues:    rawObservation.GetObservationValues(),
		ObservationTimestamp: rawObservation.GetObservationTimestamp(),
	}
}
