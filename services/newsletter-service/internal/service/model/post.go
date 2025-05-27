package model

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID           uuid.UUID
	NewsletterID uuid.UUID
	Title        string
	Content      string
	HtmlContent  string
	Published    bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type PostUpdate struct {
	Title       *string
	Content     *string
	HtmlContent *string
	Published   *bool
}
