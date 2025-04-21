package model

import (
	"4it428-newsletter-api/services/user-service/internal/service/model"
	"github.com/google/uuid"
)

type UserRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

func (u *UserRequest) ToUser() model.User {
	return model.User{
		ID:        uuid.New(),
		Email:     u.Email,
		Password:  u.Password,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}
