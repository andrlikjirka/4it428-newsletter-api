package iface

import (
	"4it428-newsletter-api/services/user-service/internal/service/model"
	"context"
)

type IUserService interface {
	CreateUser(ctx context.Context, user *model.User)
	ListUsers(ctx context.Context) ([]*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, email string) error
}
