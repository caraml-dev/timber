package storage

import (
	"context"
	"errors"

	"gorm.io/gorm"

	dserrors "github.com/caraml-dev/timber/dataset-service/errors"
	"github.com/caraml-dev/timber/dataset-service/model"
)

const LogWriterEntityName = "log_writer"

// LogWriter interface providing access for LogWriter storage
type LogWriter interface {
	// Get a log writer given its identifier
	Get(ctx context.Context, input GetInput) (model.LogWriter, error)
	// Create a new log writer and return the stored log writer with ID populated or error
	Create(ctx context.Context, lw model.LogWriter) (model.LogWriter, error)
	// Update an existing log writer and return the stored log writer or error
	Update(ctx context.Context, lw model.LogWriter) (model.LogWriter, error)
	// List all log writer given the list input
	List(ctx context.Context, listInput ListInput) ([]model.LogWriter, error)
}

// log writer storage implementation
type logWriter struct {
	db *gorm.DB
}

// NewLogWriter creates new Log Writer storage
func NewLogWriter(db *gorm.DB) LogWriter {
	return &logWriter{db: db}
}

// Get a log writer given its identifier
func (l *logWriter) Get(ctx context.Context, input GetInput) (model.LogWriter, error) {
	var logWriter model.LogWriter
	tx := l.db.WithContext(ctx).Where(&model.LogWriter{
		Base: model.Base{
			ID:        input.ID,
			ProjectID: input.ProjectID,
		},
		Name: input.Name,
	}).Take(&logWriter)

	if tx.Error != nil && errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return logWriter, dserrors.NewNotFoundError(LogWriterEntityName, input)
	}

	return logWriter, tx.Error
}

// Create a new log writer and return the stored log writer with ID populated or error
func (l *logWriter) Create(ctx context.Context, lw model.LogWriter) (model.LogWriter, error) {
	tx := l.db.WithContext(ctx).Create(&lw)
	return lw, tx.Error
}

// Update an existing log writer and return the stored log writer or error
func (l *logWriter) Update(ctx context.Context, lw model.LogWriter) (model.LogWriter, error) {
	tx := l.db.WithContext(ctx).Updates(&lw)
	return lw, tx.Error
}

// List all log writer given the list input
func (l *logWriter) List(ctx context.Context, listInput ListInput) ([]model.LogWriter, error) {
	var logWriters []model.LogWriter
	tx := l.db.WithContext(ctx).Where(&model.LogWriter{
		Base: model.Base{
			ProjectID: listInput.ProjectID,
		},
	}).Limit(listInput.Limit).
		Offset(listInput.Offset).
		Find(&logWriters)
	return logWriters, tx.Error
}
