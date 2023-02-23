package model

import (
	"database/sql/driver"

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"google.golang.org/protobuf/encoding/protojson"
)

type ObservationService struct {
	// Base provides common DB model field, namely ID, ProjectID, CreatedAt, UpdatedAt
	Base
	// Name of observation service
	Name string
	// Source of the observation service
	Source *ObservationServiceSource `gorm:"type:jsonb"`
	// Deployment status of the observation service
	Status Status
	// Error message
	Error string `gorm:"size:2048"`
}

// ObservationServiceSource is wrapper of timberv1.ObservationServiceSource proto message to allow marshalling and unmarshalling to DB
type ObservationServiceSource struct {
	*timberv1.ObservationServiceSource
}

// Value marshall ObservationServiceSource to be stored as json blob to DB
func (l *ObservationServiceSource) Value() (driver.Value, error) {
	valueString, err := protojson.Marshal(l)
	return string(valueString), err
}

// Scan parses jsonb as ObservationServiceSource
func (l *ObservationServiceSource) Scan(value interface{}) error {
	l.ObservationServiceSource = &timberv1.ObservationServiceSource{}
	if err := protojson.Unmarshal(value.([]byte), l.ObservationServiceSource); err != nil {
		return err
	}
	return nil
}
