package errors

import "fmt"

// NotFoundError error type for missing entity
type NotFoundError struct {
	// Error message
	err error
}

// Error Return error message
func (e *NotFoundError) Error() string { return e.err.Error() }

func NewNotFoundError(entity string, identity any) *NotFoundError {
	return &NotFoundError{
		err: fmt.Errorf("%s with %s is not found", entity, identity),
	}
}

// InvalidInputError error type for invalid users input
type InvalidInputError struct {
	// Error message
	err error
}

// Error Return error message
func (e *InvalidInputError) Error() string { return e.err.Error() }

type ConflictError struct {
	// Error message
	err error
}

// Error Return error message
func (e *ConflictError) Error() string { return e.err.Error() }

func NewConflictError(entity string, err error) *ConflictError {
	return &ConflictError{
		err: fmt.Errorf("%s exists: %w", entity, err),
	}
}
