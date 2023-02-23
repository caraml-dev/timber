package model

import (
	"database/sql/driver"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
)

// ObservationService observation service internal model representation
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

// ToObservationServiceProto convert internal ObservationService representation into ObservationService proto message
func (w *ObservationService) ToObservationServiceProto() *timberv1.ObservationService {
	return &timberv1.ObservationService{
		ProjectId: w.ProjectID,
		Id:        w.ID,
		Name:      w.Name,
		Source:    w.Source.ObservationServiceSource,
		Status:    w.Status.ToStatusProto(),
		Error:     w.Error,
		CreatedAt: timestamppb.New(w.CreatedAt),
		UpdatedAt: timestamppb.New(w.UpdatedAt),
	}
}

// ObservationServiceFromProto convert ObservationService proto to internal representation of ObservationService
func ObservationServiceFromProto(msg *timberv1.ObservationService) *ObservationService {
	return &ObservationService{
		Base: Base{
			ID:        msg.Id,
			ProjectID: msg.ProjectId,
		},
		Name:   msg.Name,
		Source: &ObservationServiceSource{msg.Source},
		Status: StatusFromProto(msg.Status),
		Error:  msg.Error,
	}
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
