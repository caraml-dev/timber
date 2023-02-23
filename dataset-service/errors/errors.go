package errors

import (
	"fmt"
)

// NotFoundError error type for missing entity
type NotFoundError struct {
	// Error message
	err error
}

func NewNotFoundError(entity string, identity any) *NotFoundError {
	return &NotFoundError{
		err: fmt.Errorf("%s with %s is not found", entity, identity),
	}
}

// Error Return error message
func (e *NotFoundError) Error() string { return e.err.Error() }

func (e *NotFoundError) Is(target error) bool {
	_, ok := target.(*NotFoundError)
	return ok
}

// InvalidInputError error type for invalid users input
type InvalidInputError struct {
	// Error message
	err error
}

func NewInvalidInputErrorf(format string, a ...any) *InvalidInputError {
	return &InvalidInputError{
		err: fmt.Errorf(format, a...),
	}
}

// Error Return error message
func (e *InvalidInputError) Error() string { return e.err.Error() }

func (e *InvalidInputError) Is(target error) bool {
	_, ok := target.(*InvalidInputError)
	return ok
}

type ConflictError struct {
	// Error message
	err error
}

func NewConflictError(entity string, err error) *ConflictError {
	return &ConflictError{
		err: fmt.Errorf("%s exists: %w", entity, err),
	}
}

// Error Return error message
func (e *ConflictError) Error() string { return e.err.Error() }

func (e *ConflictError) Is(target error) bool {
	_, ok := target.(*ConflictError)
	return ok
}
