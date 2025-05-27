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
	id, err := uuid.Parse(u.ID)
	if err != nil {
		id = uuid.Nil
	}
	newsletterID, err := uuid.Parse(u.NewsletterID)
	if err != nil {
		newsletterID = uuid.Nil
	}
	return &SubscriptionResponse{
		ID:           id,
		Email:        u.Email,
		NewsletterID: newsletterID,
	}
}

func FromSubscriptionList(subscriptions []*model.Subscription) []*SubscriptionResponse {
	responses := make([]*SubscriptionResponse, 0, len(subscriptions))
	for _, s := range subscriptions {
		responses = append(responses, FromSubscription(s))
	}
	return responses
}
