package impl

import (
	"4it428-newsletter-api/libs/logger"
	"4it428-newsletter-api/services/user-service/internal/service/iface"
	"4it428-newsletter-api/services/user-service/internal/service/model"
	"context"
)

type userService struct{}

func NewUserService() iface.IUserService {
	return &userService{}
}

func (u *userService) CreateUser(_ context.Context, user *model.User) {
	logger.Info("Creating user: %+v\n", user)
}

func (u *userService) ListUsers(_ context.Context) ([]*model.User, error) {
	logger.Info("Listing all users")
	return []*model.User{}, nil
}

func (u *userService) GetByEmail(_ context.Context, email string) (*model.User, error) {
	logger.Info("Getting user by email: %s\n", email)
	return &model.User{Email: email}, nil
}

func (u *userService) UpdateUser(_ context.Context, user *model.User) (*model.User, error) {
	logger.Info("Updating user: %+v\n", user)
	return user, nil
}

func (u *userService) DeleteUser(_ context.Context, email string) error {
	logger.Info("Deleting user with email: %s\n", email)
	return nil
}
