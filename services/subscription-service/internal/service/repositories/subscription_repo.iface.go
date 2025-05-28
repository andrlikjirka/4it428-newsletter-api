package repositories

import (
	"4it428-newsletter-api/services/subscription-service/internal/service/model"
	"context"
	"github.com/google/uuid"
)

type ISubscriptionRepository interface {
	Delete(ctx context.Context, id uuid.UUID) error
	Add(ctx context.Context, subscription *model.Subscription) (*model.Subscription, error)
	ListByNewsletterId(ctx context.Context, newsletterID uuid.UUID) ([]*model.Subscription, error)
}
