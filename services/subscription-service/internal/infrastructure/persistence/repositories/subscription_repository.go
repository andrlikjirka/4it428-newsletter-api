package repositories

import (
	"4it428-newsletter-api/services/subscription-service/internal/service/model"
	"context"
	"github.com/google/uuid"
)

type SubscriptionRepository struct {
}

func NewSubscriptionRepository() *SubscriptionRepository {
	return &SubscriptionRepository{}
}

func (r *SubscriptionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	// TODO: delete subscription by id from the database
	return nil
}

func (r *SubscriptionRepository) Add(ctx context.Context, subscription *model.Subscription) (*model.Subscription, error) {
	// TODO: insert subscription into the database
	return nil, nil // Replace with actual implementation
}

func (r *SubscriptionRepository) ListByNewsletterId(ctx context.Context, newsletterID uuid.UUID) ([]*model.Subscription, error) {
	// TODO: list all subscriptions by newsletter ID from the database
	return nil, nil // Replace with actual implementation
}
