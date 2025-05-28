package model

import (
	"4it428-newsletter-api/services/newsletter-service/internal/service/model"
	"github.com/google/uuid"
	"time"
)

type NewsletterResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      uuid.UUID `json:"user_id"`
}

func FromNewsletter(u *model.Newsletter) *NewsletterResponse {
	return &NewsletterResponse{
		ID:          u.ID,
		Title:       u.Title,
		Description: u.Description,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		UserID:      u.UserID,
	}
}

func FromNewsletterList(newsletters []*model.Newsletter) []*NewsletterResponse {
	responses := make([]*NewsletterResponse, 0, len(newsletters))
	for _, u := range newsletters {
		responses = append(responses, FromNewsletter(u))
	}
	return responses
}
