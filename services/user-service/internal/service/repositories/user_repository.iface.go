package repositories

import (
	"4it428-newsletter-api/services/user-service/internal/service/model"
	"context"
)

type UserRepository interface {
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, email string) error
	List(ctx context.Context) ([]*model.User, error)
}
