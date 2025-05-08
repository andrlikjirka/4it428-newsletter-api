package services

import (
	"4it428-newsletter-api/services/newsletter-service/internal/service/model"
	"context"
)

type NewsletterService interface {
	CreateNewsletter(ctx context.Context, newsletter *model.Newsletter) (*model.Newsletter, error)
	ListNewsletters(ctx context.Context) ([]*model.Newsletter, error)
	GetNewsletterById(ctx context.Context, id string) (*model.Newsletter, error)
	UpdateNewsletter(ctx context.Context, id string, newsletter *model.NewsletterUpdate) (*model.Newsletter, error)
	DeleteNewsletter(ctx context.Context, id string) error
}
