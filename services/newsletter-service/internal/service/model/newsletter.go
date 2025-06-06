package model

import (
	"github.com/google/uuid"
	"time"
)

type Newsletter struct {
	ID          uuid.UUID
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	UserID      uuid.UUID
}

type NewsletterUpdate struct {
	Title       *string
	Description *string
}
