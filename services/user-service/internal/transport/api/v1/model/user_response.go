package model

import (
	"4it428-newsletter-api/services/user-service/internal/service/model"
	"github.com/google/uuid"
)

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
}

// MAPPERS:

func FromUser(u *model.User) *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

func FromUserList(users []*model.User) []*UserResponse {
	responses := make([]*UserResponse, 0, len(users))
	for _, u := range users {
		responses = append(responses, FromUser(u))
	}
	return responses
}
