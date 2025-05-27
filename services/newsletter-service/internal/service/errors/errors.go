package errors

import "errors"

var (
	ErrInvalidUUID      = errors.New("invalid UUID format")
	ErrNotFound         = errors.New("newsletter not found")
	ErrPostNotFound     = errors.New("post not found")
	ErrAlreadyPublished = errors.New("post already published")
)
