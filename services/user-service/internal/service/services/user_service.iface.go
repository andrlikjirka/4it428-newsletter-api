package services

import (
	"4it428-newsletter-api/services/user-service/internal/service/model"
	"context"
)

type IUserService interface {
	CreateUser(ctx context.Context, user *model.User) error
	ListUsers(ctx context.Context) []*model.User
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateUser(ctx context.Context, email string, user *model.UserUpdate) (*model.User, error)
	DeleteUser(ctx context.Context, email string) error
}
