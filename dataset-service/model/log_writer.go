package model

import (
	"database/sql/driver"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
)

// LogWriter data model for log writer
type LogWriter struct {
	// Base provides common DB model field, namely ID, ProjectID, CreatedAt, UpdatedAt
	Base
	// Name of the log writer
	Name string
	// Log Writer Source configuration
	Source *LogWriterSource `gorm:"type:jsonb"`
	// Deployment status of the log writer
	Status Status
	// Error message
	Error string `gorm:"size:2048"`
}

// ToLogWriterProto convert internal LogWriter representation into LogWriter proto message
func (w *LogWriter) ToLogWriterProto() *timberv1.LogWriter {
	return &timberv1.LogWriter{
		ProjectId: w.ProjectID,
		Id:        w.ID,
		Name:      w.Name,
		Source:    w.Source.LogWriterSource,
		Status:    w.Status.ToStatusProto(),
		Error:     w.Error,
		CreatedAt: timestamppb.New(w.CreatedAt),
		UpdatedAt: timestamppb.New(w.UpdatedAt),
	}
}

// LogWriterFromProto convert LogWriter proto to internal representation of LogWriter
func LogWriterFromProto(msg *timberv1.LogWriter) *LogWriter {
	var createdTime time.Time
	var updatedTime time.Time

	if msg.CreatedAt != nil {
		createdTime = msg.CreatedAt.AsTime()
	}

	if msg.UpdatedAt != nil {
		updatedTime = msg.UpdatedAt.AsTime()
	}

	return &LogWriter{
		Base: Base{
			ID:        msg.Id,
			ProjectID: msg.ProjectId,
			CreatedAt: createdTime,
			UpdatedAt: updatedTime,
		},
		Name:   msg.Name,
		Source: &LogWriterSource{msg.Source},
		Status: StatusFromProto(msg.Status),
		Error:  msg.Error,
	}
}

// LogWriterSource is wrapper of LogWriterSource proto message to allow marshalling and unmarshalling to DB
type LogWriterSource struct {
	*timberv1.LogWriterSource
}

// Value marshall LogWriterSource to be stored as json blob to DB
func (l *LogWriterSource) Value() (driver.Value, error) {
	valueString, err := protojson.Marshal(l)
	return string(valueString), err
}

// Scan parses jsonb as LogWriterSource
func (l *LogWriterSource) Scan(value interface{}) error {
	l.LogWriterSource = &timberv1.LogWriterSource{}
	if err := protojson.Unmarshal(value.([]byte), l.LogWriterSource); err != nil {
		return err
	}
	return nil
}
