package errors

import (
	"fmt"
)

// NotFoundError error type for missing entity
type NotFoundError struct {
	// Error message
	err error
}

// NewNotFoundError creates new instance of NotFoundError
func NewNotFoundError(entity string, identity any) *NotFoundError {
	return &NotFoundError{
		err: fmt.Errorf("%s with %s is not found", entity, identity),
	}
}

// Error Return error message
func (e *NotFoundError) Error() string { return e.err.Error() }

// Is check whether the error is NotFoundError
func (e *NotFoundError) Is(target error) bool {
	_, ok := target.(*NotFoundError)
	return ok
}

// InvalidInputError error type for invalid users input
type InvalidInputError struct {
	// Error message
	err error
}

// NewInvalidInputErrorf creates new instance of InvalidInputError
func NewInvalidInputErrorf(format string, a ...any) *InvalidInputError {
	return &InvalidInputError{
		err: fmt.Errorf(format, a...),
	}
}

// Error Return error message
func (e *InvalidInputError) Error() string { return e.err.Error() }

// Is check whether the error is InvalidInputError
func (e *InvalidInputError) Is(target error) bool {
	_, ok := target.(*InvalidInputError)
	return ok
}

// ConflictError error type for creating an existing entity
type ConflictError struct {
	// Error message
	err error
}

// NewConflictError creates new ConflictError instance
func NewConflictError(entity string, err error) *ConflictError {
	return &ConflictError{
		err: fmt.Errorf("%s exists: %w", entity, err),
	}
}

// Error Return error message
func (e *ConflictError) Error() string { return e.err.Error() }

// Is checks whether an error is due to ConflictError
func (e *ConflictError) Is(target error) bool {
	_, ok := target.(*ConflictError)
	return ok
}
