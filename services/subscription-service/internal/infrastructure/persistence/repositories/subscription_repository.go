package repositories

import (
	"4it428-newsletter-api/pkg/logger"
	dbmodel "4it428-newsletter-api/services/subscription-service/internal/infrastructure/persistence/model"
	"4it428-newsletter-api/services/subscription-service/internal/service/model"
	"cloud.google.com/go/firestore"
	"context"
	"github.com/google/uuid"
)

type SubscriptionRepository struct {
	client *firestore.Client
}

func NewSubscriptionRepository(client *firestore.Client) *SubscriptionRepository {
	return &SubscriptionRepository{client: client}
}

func (r *SubscriptionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	iter := r.client.Collection("subscriptions").Where("ID", "==", id.String()).Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		_, err = doc.Ref.Delete(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *SubscriptionRepository) Add(ctx context.Context, subscription *model.Subscription) (*model.Subscription, error) {
	logger.Info("Adding subscription to Firestore", "subscriptionID", subscription.ID, "email", subscription.Email, "newsletterID", subscription.NewsletterID)
	entity := &dbmodel.SubscriptionEntity{
		ID:           subscription.ID.String(),
		Email:        subscription.Email,
		NewsletterID: subscription.NewsletterID.String(),
	}
	docRef := r.client.Collection("subscriptions").NewDoc()
	_, err := docRef.Set(ctx, entity)
	if err != nil {
		return nil, err
	}
	return subscription, nil
}

func (r *SubscriptionRepository) ListByNewsletterId(ctx context.Context, newsletterID uuid.UUID) ([]*model.Subscription, error) {
	iter := r.client.Collection("subscriptions").Where("NewsletterID", "==", newsletterID.String()).Documents(ctx)
	defer iter.Stop()

	var subs []*model.Subscription
	for {
		doc, err := iter.Next()
		if err != nil {
			break
		}
		var entity dbmodel.SubscriptionEntity
		if err := doc.DataTo(&entity); err != nil {
			continue
		}
		sub := &model.Subscription{
			ID:           uuid.MustParse(entity.ID),
			Email:        entity.Email,
			NewsletterID: uuid.MustParse(entity.NewsletterID),
		}
		subs = append(subs, sub)
	}
	return subs, nil
}
