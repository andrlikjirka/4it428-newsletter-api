package repositories

import (
	"4it428-newsletter-api/services/newsletter-service/internal/service/model"
	"context"
	"github.com/google/uuid"
)

type IPostRepository interface {
	Add(ctx context.Context, post *model.Post) (*model.Post, error)
	List(ctx context.Context, newsletterID uuid.UUID) ([]*model.Post, error)
	GetById(ctx context.Context, postID uuid.UUID, newsletterID uuid.UUID) (*model.Post, error)
	Update(ctx context.Context, post *model.Post) (*model.Post, error)
	Delete(ctx context.Context, postID uuid.UUID, newsletterID uuid.UUID) error
}
