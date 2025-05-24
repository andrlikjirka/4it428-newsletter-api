package repositories

import (
	"4it428-newsletter-api/services/newsletter-service/internal/service/model"
	"context"
	"github.com/google/uuid"
)

type INewsletterRepository interface {
	Add(ctx context.Context, newsletter *model.Newsletter) (*model.Newsletter, error)
	List(ctx context.Context) ([]*model.Newsletter, error)
	GetById(ctx context.Context, id uuid.UUID) (*model.Newsletter, error)
	Update(ctx context.Context, newsletter *model.Newsletter) (*model.Newsletter, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
