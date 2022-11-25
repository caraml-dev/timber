package models

import (
	"encoding/json"

	upiv1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/caraml-dev/observation-service/observation-service/errors"
)

type ObservationLogKey struct {
	EventTimestamp int64
}

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
		return kvPairs, errors.Wrapf(err, "Error marshaling the result log")
	}
	// Unmarshal into map[string]interface{}
	err = json.Unmarshal(bytes, &kvPairs)
	if err != nil {
		return kvPairs, errors.Wrapf(err, "Error unmarshaling the result log")
	}
	return kvPairs, nil
}

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
