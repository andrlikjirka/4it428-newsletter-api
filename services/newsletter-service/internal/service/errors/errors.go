package errors

import "errors"

var (
	ErrInvalidUUID = errors.New("invalid UUID format")
	ErrNotFound    = errors.New("newsletter not found")
)
