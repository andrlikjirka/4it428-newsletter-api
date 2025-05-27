package impl

import (
	"4it428-newsletter-api/libs/logger"
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
	//TODO delete subscription by ID from the database
	logger.Info("Unsubscribing from newsletter", "subscriptionID", subscriptionID)
	panic("implement me")
}

func (s subscriptionService) Subscribe(ctx context.Context, subscription *model.Subscription) (*model.Subscription, error) {
	//TODO save subscription in the database
	//TODO send confirmation email to the user
	logger.Info("Subscribing to newsletter", "email", subscription.Email, "newsletterID", subscription.NewsletterID.String())
	panic("implement me")
}

func (s subscriptionService) ListSubscriptions(ctx context.Context, newsletterID string, userID string) ([]*model.Subscription, error) {
	//TODO check user ID is owner of the subscription
	//TODO fetch subscriptions from the database
	logger.Info("Listing subscriptions for newsletter", "newsletterID", newsletterID, "userID", userID)
	panic("implement me")
}

func (s subscriptionService) NotifySubscribers(ctx context.Context, newsletterID string, notification *model.Notification) error {
	//TODO fetch all subscriptions for the given newsletter ID
	//TODO send notification to all subscribers via email including unsubscribe link
	logger.Info("Notifying subscribers for newsletter", "newsletterID", newsletterID, "notificationTitle", notification.Title)
	panic("implement me")
}
