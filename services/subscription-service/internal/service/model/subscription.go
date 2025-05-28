package model

import "github.com/google/uuid"

type Subscription struct {
	ID           uuid.UUID
	Email        string
	NewsletterID uuid.UUID
}
