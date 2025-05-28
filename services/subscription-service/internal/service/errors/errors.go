package errors

import "errors"

var (
	ErrNoNewsletterId     = errors.New("newsletter id is required")
	ErrInvalidUUID        = errors.New("invalid UUID format")
	ErrNewsletterNotFound = errors.New("newsletter not found")
	ErrUnauthorized       = errors.New("unauthorized access to list of subscribers")
)
