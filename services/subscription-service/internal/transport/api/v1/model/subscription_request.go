package model

import (
	"4it428-newsletter-api/services/subscription-service/internal/service/model"
	"github.com/google/uuid"
)

type SubscriptionRequest struct {
	Email        string    `json:"email" validate:"required,email,min=3,max=50"`
	NewsletterID uuid.UUID `json:"newsletter_id" validate:"required"`
}

func (n *SubscriptionRequest) ToSubscription() *model.Subscription {
	return &model.Subscription{
		ID:           uuid.New().String(),
		Email:        n.Email,
		NewsletterID: n.NewsletterID.String(),
	}
}

type NotifySubscribersRequest struct {
	Title       string `json:"title" validate:"required"`
	Content     string `json:"content" validate:"required"`
	HtmlContent string `json:"html_content" validate:"required"`
}

func (n *NotifySubscribersRequest) ToNotification() *model.Notification {
	return &model.Notification{
		Title:       n.Title,
		Content:     n.Content,
		HtmlContent: n.HtmlContent,
	}
}
