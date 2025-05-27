package errors

import "errors"

var (
	ErrNoNewsletterId = errors.New("newsletter id is required")
	ErrInvalidUUID    = errors.New("invalid UUID format")
)
