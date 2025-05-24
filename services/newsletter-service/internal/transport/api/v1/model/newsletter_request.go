package model

import (
	"4it428-newsletter-api/services/newsletter-service/internal/service/model"
	"github.com/google/uuid"
)

type CreateNewsletterRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=50"`
	Description string `json:"description" validate:"required,max=500"`
}

type UpdateNewsletterRequest struct {
	Title       *string `json:"title,omitempty" validate:"omitempty,min=3,max=50"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=500"`
}

func (n *CreateNewsletterRequest) ToNewsletter() *model.Newsletter {
	return &model.Newsletter{
		ID:          uuid.New(),
		Title:       n.Title,
		Description: n.Description,
	}
}

func (n *UpdateNewsletterRequest) ToNewsletterUpdate() *model.NewsletterUpdate {
	return &model.NewsletterUpdate{
		Title:       n.Title,
		Description: n.Description,
	}
}
