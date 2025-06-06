package impl

import (
	"4it428-newsletter-api/pkg/logger"
	"4it428-newsletter-api/services/user-service/internal/service/auth"
	"4it428-newsletter-api/services/user-service/internal/service/errors"
	"4it428-newsletter-api/services/user-service/internal/service/model"
	"4it428-newsletter-api/services/user-service/internal/service/repositories"
	"4it428-newsletter-api/services/user-service/internal/service/services"
	"context"
)

type userService struct {
	authProvider auth.IAuthProvider
	repo         repositories.UserRepository
}

func NewUserService(authProvider auth.IAuthProvider, repo repositories.UserRepository) services.UserService {
	return &userService{
		authProvider: authProvider,
		repo:         repo,
	}
}

func (u *userService) CreateUser(ctx context.Context, user *model.User) error {
	err := u.repo.Add(ctx, user)
	if err != nil {
		logger.Error("Failed to create user", "error", err)
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
		return nil, errors.ErrUserNotFound
	}

	logger.Info("Getting user by email: " + email)
	return user, nil
}

func (u *userService) UpdateUser(ctx context.Context, email string, userToUpdate *model.UserUpdate) (*model.User, error) {
	existingUser, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	if userToUpdate.FirstName != nil {
		existingUser.FirstName = *userToUpdate.FirstName
	}
	if userToUpdate.LastName != nil {
		existingUser.LastName = *userToUpdate.LastName
	}

	user, err := u.repo.Update(ctx, existingUser)
	if err != nil {
		logger.Error("Failed to update user", "id", user.ID, "error", err)
		return nil, err
	}

	logger.Info("User updated successfully", "id", user.ID)
	return user, nil
}

func (u *userService) DeleteUser(ctx context.Context, email string) error {
	userRecord, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return errors.ErrUserNotFound
	}

	err = u.repo.Delete(ctx, email)
	if err != nil {
		logger.Error("Failed to delete user", "id", userRecord, "error", err)
		return err
	}
	logger.Info("Deleted user from DB with email: " + email)

	err = u.authProvider.DeleteUser(ctx, userRecord.FirebaseUID)
	if err != nil {
		logger.Error("Failed to delete user from Firebase: "+email, "error", err)
	}
	return nil
}
