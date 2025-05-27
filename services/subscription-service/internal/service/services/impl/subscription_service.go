package impl

import (
	"4it428-newsletter-api/libs/logger"
	"4it428-newsletter-api/libs/utils"
	errors2 "4it428-newsletter-api/services/subscription-service/internal/service/errors"
	"4it428-newsletter-api/services/subscription-service/internal/service/model"
	"4it428-newsletter-api/services/subscription-service/internal/service/repositories"
	"4it428-newsletter-api/services/subscription-service/internal/service/services"
	"context"
)

type subscriptionService struct {
	repo repositories.ISubscriptionRepository
}

func NewSubscriptionService(repo repositories.ISubscriptionRepository) services.SubscriptionService {
	return &subscriptionService{repo: repo}
}

func (s subscriptionService) Unsubscribe(ctx context.Context, subscriptionID string) error {
	parsedSubscriptionID, err := utils.ParseUUID(subscriptionID)
	if err != nil {
		logger.Error("Failed to parse UUID", "subscriptionID", subscriptionID, "error", err)
		return errors2.ErrInvalidUUID
	}
	err = s.repo.Delete(ctx, parsedSubscriptionID)
	if err != nil {
		logger.Error("Failed to unsubscribe from newsletter", "subscriptionID", subscriptionID, "error", err)
		return err
	}
	logger.Info("Unsubscribing from newsletter", "subscriptionID", subscriptionID)
	return nil
}

func (s subscriptionService) Subscribe(ctx context.Context, subscription *model.Subscription) (*model.Subscription, error) {
	createdSubscription, err := s.repo.Add(ctx, subscription)
	if err != nil {
		logger.Error("Failed to subscribe to newsletter", "email", subscription.Email, "newsletterID", subscription.NewsletterID.String(), "error", err)
		return nil, err
	}

	//TODO send confirmation email to the user

	logger.Info("Subscribing to newsletter", "email", subscription.Email, "newsletterID", subscription.NewsletterID.String())
	return createdSubscription, nil
}

func (s subscriptionService) ListSubscriptions(ctx context.Context, newsletterID string, userID string) ([]*model.Subscription, error) {
	parsedNewsletterID, err := utils.ParseUUID(newsletterID)
	if err != nil {
		logger.Error("Failed to parse UUID", "newsletterID", newsletterID, "error", err)
		return nil, errors2.ErrInvalidUUID
	}

	//TODO check user ID is owner of the subscription

	subscriptions, err := s.repo.ListByNewsletterId(ctx, parsedNewsletterID)
	if err != nil {
		return nil, err
	}

	logger.Info("Listing subscriptions for newsletter", "newsletterID", newsletterID, "userID", userID)
	return subscriptions, nil
}

func (s subscriptionService) NotifySubscribers(ctx context.Context, newsletterID string, notification *model.Notification) error {
	parsedNewsletterID, err := utils.ParseUUID(newsletterID)
	if err != nil {
		logger.Error("Failed to parse UUID", "newsletterID", newsletterID, "error", err)
		return errors2.ErrInvalidUUID
	}
	subscriptions, err := s.repo.ListByNewsletterId(ctx, parsedNewsletterID)
	if err != nil {
		logger.Error("Failed to list subscriptions for newsletter", "newsletterID", newsletterID, "error", err)
		return err
	}

	logger.Info("Found subscriptions for newsletter", "newsletterID", newsletterID, "count", len(subscriptions))

	//TODO send notification to all subscribers via email including unsubscribe link

	logger.Info("Notifying subscribers for newsletter", "newsletterID", newsletterID, "notificationTitle", notification.Title)
	return nil
}
