package storage

import (
	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/jackc/pgconn"
)

// GetInput is common type used for querying specific entity in its storage
type GetInput struct {
	// ID of the entity
	ID int64
	// Name of the entity
	Name string
	// ProjectID of the entity belongs to
	ProjectID int64
}

// ListInput is common type used for querying list of entities in its storage
type ListInput struct {
	// Project ID
	ProjectID int64
	// Starting offset of the list request
	Offset int
	// Limit number of entities returned
	Limit int
}

func ListInputFromOption(projectID int64, detail *timberv1.ListOption) ListInput {
	if detail == nil {
		detail = &timberv1.ListOption{
			Offset: 0,
			Limit:  10,
		}
	}

	return ListInput{
		ProjectID: projectID,
		Offset:    int(detail.Offset),
		Limit:     int(detail.Limit),
	}
}

var duplicateEntryError = &pgconn.PgError{Code: "23505"}
