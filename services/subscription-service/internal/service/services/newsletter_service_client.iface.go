package services

import (
	"context"
	"github.com/google/uuid"
)

type Newsletter struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	UserID      string `json:"user_id"`
}

type INewsletterServiceClient interface {
	GetNewsletter(ctx context.Context, newsletterID uuid.UUID) (*Newsletter, error)
}
