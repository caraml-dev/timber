package storage

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
	// Starting offset of the list request
	Offset int
	// Limit number of entities returned
	Limit int
}
