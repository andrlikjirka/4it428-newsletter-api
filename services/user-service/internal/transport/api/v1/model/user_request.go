package model

import (
	"4it428-newsletter-api/services/user-service/internal/service/model"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

type UpdateUserRequest struct {
	Password  *string `json:"password,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
}

func (u *CreateUserRequest) ToUser() *model.User {
	return &model.User{
		ID:        uuid.New(),
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

func (u *UpdateUserRequest) ToUserUpdate() *model.UserUpdate {
	return &model.UserUpdate{
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}
