package model

import "4it428-newsletter-api/services/user-service/internal/service/model"

type SignUpRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

func (s *SignUpRequest) ToServiceInput() *model.SignUpInput {
	return &model.SignUpInput{
		Email:     s.Email,
		Password:  s.Password,
		FirstName: s.FirstName,
		LastName:  s.LastName,
	}
}
