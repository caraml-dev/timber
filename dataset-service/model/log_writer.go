package model

import (
	"database/sql/driver"

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"google.golang.org/protobuf/encoding/protojson"
)

// LogWriter data model for log writer
type LogWriter struct {
	// Base provides common DB model field
	// ID        		uint `gorm:"primarykey"`
	// ProjectID        uint `gorm:"primarykey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt DeletedAt `gorm:"index"`
	Base
	// Name of the log writer
	Name string `gorm:"primaryKey"`
	// Deployment status of the log writer
	Status Status
	// Log Writer Source configuration
	LogWriterSource *LogWriterSource `gorm:"type:jsonb"`
}

// ToLogWriterProto convert internal LogWriter representation into LogWriter proto message
func (w *LogWriter) ToLogWriterProto() *timberv1.LogWriter {
	return nil
}

// FromLogWriterProto convert LogWriter proto to internal representation of LogWriter
func FromLogWriterProto(msg *timberv1.LogWriter) *LogWriter {
	return nil
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
