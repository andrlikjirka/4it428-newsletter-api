package bootstrap

import (
	"4it428-newsletter-api/services/subscription-service/internal/infrastructure/aws"
	"4it428-newsletter-api/services/subscription-service/internal/infrastructure/persistence/repositories"
	"4it428-newsletter-api/services/subscription-service/internal/service/services"
	"4it428-newsletter-api/services/subscription-service/internal/service/services/impl"
	"4it428-newsletter-api/services/subscription-service/internal/transport/api/v1/handler"
	"cloud.google.com/go/firestore"
	"context"
	"google.golang.org/api/option"
	"os"
)

func SetupFirestore(ctx context.Context) (*firestore.Client, error) {
	credentials := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	projectID := os.Getenv("GOOGLE_PROJECT_ID")
	client, err := firestore.NewClient(ctx, projectID, option.WithCredentialsFile(credentials))
	if err != nil {
		return nil, err
	}
	return client, nil
}

func SetupAwsSes(ctx context.Context) (*aws.SESClient, error) {
	awsRegion := os.Getenv("AWS_REGION")
	sender := os.Getenv("SES_SENDER_EMAIL")
	client, err := aws.NewSESClient(ctx, sender, awsRegion)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// CONTAINERS FOR EASIER DI:

type HandlersContainer struct {
	SubscriptionHandler *handler.SubscriptionHandler
}

func NewHandlersContainer(s *ServicesContainer) *HandlersContainer {
	return &HandlersContainer{
		SubscriptionHandler: handler.NewSubscriptionHandler(s.SubscriptionService),
	}
}

type ServicesContainer struct {
	SubscriptionService services.SubscriptionService
}

func NewServicesContainer(
	subscriptionRepository *repositories.SubscriptionRepository,
	awsSesClient *aws.SESClient,
	newsletterServiceClient services.INewsletterServiceClient,
) *ServicesContainer {
	return &ServicesContainer{
		SubscriptionService: impl.NewSubscriptionService(subscriptionRepository, awsSesClient, newsletterServiceClient),
	}
}
