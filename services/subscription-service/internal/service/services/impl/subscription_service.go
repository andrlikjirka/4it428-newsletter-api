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
	repo                    repositories.ISubscriptionRepository
	ses                     services.EmailProvider
	newsletterServiceClient services.INewsletterServiceClient
}

func NewSubscriptionService(
	repo repositories.ISubscriptionRepository,
	ses services.EmailProvider,
	newsletterServiceClient services.INewsletterServiceClient,
) services.SubscriptionService {
	return &subscriptionService{repo: repo, ses: ses, newsletterServiceClient: newsletterServiceClient}
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
	_, err := s.newsletterServiceClient.GetNewsletter(ctx, subscription.NewsletterID)
	if err != nil {
		logger.Error("Failed to get newsletter", "newsletterID", subscription.NewsletterID, "error", err)
		return nil, errors2.ErrNewsletterNotFound
	}

	logger.Info("Subscribing to newsletter", "subscriptionID", subscription.ID)

	createdSubscription, err := s.repo.Add(ctx, subscription)
	if err != nil {
		logger.Error("Failed to subscribe to newsletter", "email", subscription.Email, "newsletterID", subscription.NewsletterID, "error", err)
		return nil, err
	}

	err = s.ses.SendEmail(ctx, subscription.Email, "Subscription Confirmation",
		"Thank you for subscribing to our newsletter!",
		"<h1>Thank you for subscribing to our newsletter!</h1>")
	if err != nil {
		logger.Error("Failed to send confirmation email", "email", subscription.Email, "error", err)
		return nil, err
	}

	logger.Info("Subscribing to newsletter", "email", subscription.Email, "newsletterID", subscription.NewsletterID)
	return createdSubscription, nil
}

func (s subscriptionService) ListSubscriptions(ctx context.Context, newsletterID string, userID string) ([]*model.Subscription, error) {
	parsedNewsletterID, err := utils.ParseUUID(newsletterID)
	if err != nil {
		logger.Error("Failed to parse UUID", "newsletterID", newsletterID, "error", err)
		return nil, errors2.ErrInvalidUUID
	}

	newsletter, err := s.newsletterServiceClient.GetNewsletter(ctx, parsedNewsletterID)
	if err != nil {
		logger.Error("Failed to get newsletter", "newsletterID", newsletterID, "error", err)
		return nil, errors2.ErrNewsletterNotFound
	}
	logger.Info("Got newsletter for listing subscriptions", "newsletter", newsletter)
	if newsletter.UserID != userID {
		logger.Error("User is not authorized to list subscriptions for this newsletter", "newsletterID", newsletterID, "userID", userID)
		return nil, errors2.ErrUnauthorized
	}

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

	for _, sub := range subscriptions {
		htmlContent := notification.HtmlContent + utils.GenerateUnsubscribeLink(sub.ID)
		err = s.ses.SendEmail(ctx, sub.Email, notification.Title, notification.Content, htmlContent)
		if err != nil {
			logger.Error("Failed to send notification email", "email", sub.Email, "error", err)
			return err
		}
		logger.Info("Sent notification email", "email", sub.Email, "newsletterID", newsletterID)
	}

	logger.Info("Notifying subscribers for newsletter", "newsletterID", newsletterID, "notificationTitle", notification.Title)
	return nil
}
