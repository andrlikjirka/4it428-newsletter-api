package services

import (
	"context"
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
	GetNewsletter(ctx context.Context, newsletterID string) (*Newsletter, error)
}
