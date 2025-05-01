package services

import (
	"4it428-newsletter-api/libs/logger"
	"4it428-newsletter-api/services/user-service/internal/service/errors"
	"4it428-newsletter-api/services/user-service/internal/service/model"
	"4it428-newsletter-api/services/user-service/internal/service/repositories"
	"context"
)

type userService struct {
	repo repositories.IUserRepository
}

func NewUserService(repo repositories.IUserRepository) IUserService {
	return &userService{
		repo: repo,
	}
}

func (u *userService) CreateUser(ctx context.Context, user *model.User) error {
	err := u.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	logger.Info("Creating new user with email " + user.Email)
	return nil
}

func (u *userService) ListUsers(ctx context.Context) ([]*model.User, error) {
	users, err := u.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	logger.Info("Listing all users")
	return users, nil
}

func (u *userService) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return &model.User{}, errors.ErrUserNotFound
	}

	logger.Info("Getting user by email: " + email)
	return user, nil
}

func (u *userService) UpdateUser(ctx context.Context, email string, userToUpdate *model.UserUpdate) (*model.User, error) {
	existingUser, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if existingUser == nil {
		return nil, errors.ErrUserNotFound
	}

	if userToUpdate.FirstName != nil {
		existingUser.FirstName = *userToUpdate.FirstName
	}
	if userToUpdate.LastName != nil {
		existingUser.LastName = *userToUpdate.LastName
	}

	user, err := u.repo.Update(ctx, existingUser)
	logger.Info("Updating user: %+v\n", user)
	return user, nil
}

func (u *userService) DeleteUser(ctx context.Context, email string) error {
	err := u.repo.Delete(ctx, email)
	if err != nil {
		return err
	}

	logger.Info("Deleting user with email: " + email)
	return nil
}
