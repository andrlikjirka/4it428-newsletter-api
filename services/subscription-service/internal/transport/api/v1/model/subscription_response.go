package model

import (
	"4it428-newsletter-api/services/subscription-service/internal/service/model"
	"github.com/google/uuid"
)

type SubscriptionResponse struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	NewsletterID uuid.UUID `json:"newsletter_id"`
}

func FromSubscription(u *model.Subscription) *SubscriptionResponse {
	return &SubscriptionResponse{
		ID:           u.ID,
		Email:        u.Email,
		NewsletterID: u.NewsletterID,
	}
}

func FromSubscriptionList(subscriptions []*model.Subscription) []*SubscriptionResponse {
	responses := make([]*SubscriptionResponse, 0, len(subscriptions))
	for _, s := range subscriptions {
		responses = append(responses, FromSubscription(s))
	}
	return responses
}
