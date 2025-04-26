package services

import (
	"4it428-newsletter-api/libs/logger"
	"4it428-newsletter-api/services/user-service/internal/service/model"
	"context"
	"errors"
)

var (
	users = map[string]*model.User{}
)

type userService struct{}

func NewUserService() IUserService {
	return &userService{}
}

// CreateUser method saves user in map under email as a key.
func (u *userService) CreateUser(_ context.Context, user *model.User) error {
	if _, exists := users[user.Email]; exists {
		return errors.New("user already exists")
	}
	users[user.Email] = user

	logger.Info("Creating new user with email " + user.Email)
	return nil
}

// ListUsers method returns list of users in array of users.
func (u *userService) ListUsers(_ context.Context) []*model.User {
	logger.Info("Listing all users")
	usersList := make([]*model.User, 0, len(users))
	for _, user := range users {
		usersList = append(usersList, user)
	}

	return usersList
}

// GetUserByEmail method returns an user with specified email.
func (u *userService) GetUserByEmail(_ context.Context, email string) (*model.User, error) {
	user, exists := users[email]
	if !exists {
		return &model.User{}, errors.New("user does not exist")
	}

	logger.Info("Getting user by email: " + email)
	return user, nil
}

// UpdateUser updates attributes of a specified user.
func (u *userService) UpdateUser(_ context.Context, email string, userToUpdate *model.UserUpdate) (*model.User, error) {
	user, exists := users[email]
	if !exists {
		return nil, errors.New("user not found")
	}

	if userToUpdate.FirstName != nil {
		user.FirstName = *userToUpdate.FirstName
	}
	if userToUpdate.LastName != nil {
		user.LastName = *userToUpdate.LastName
	}
	if userToUpdate.Password != nil {
		user.Password = *userToUpdate.Password // hash it !!!
	}

	logger.Info("Updating user: %+v\n", user)
	return user, nil
}

// DeleteUser deletes user from memory.
func (u *userService) DeleteUser(_ context.Context, email string) error {
	logger.Info("Deleting user with email: " + email)
	if _, exists := users[email]; !exists {
		return errors.New("user does not exist")
	}

	delete(users, email)
	return nil
}
