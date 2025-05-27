package services

import (
	"4it428-newsletter-api/services/subscription-service/internal/service/model"
	"context"
)

type SubscriptionService interface {
	Unsubscribe(ctx context.Context, subscriptionID string) error
	Subscribe(ctx context.Context, subscription *model.Subscription) (*model.Subscription, error)
	ListSubscriptions(ctx context.Context, newsletterID string, userID string) ([]*model.Subscription, error)
	NotifySubscribers(ctx context.Context, newsletterID string, notification *model.Notification) error
}
