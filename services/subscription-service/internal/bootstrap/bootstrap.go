package bootstrap

import (
	"4it428-newsletter-api/services/subscription-service/internal/infrastructure/aws"
	"4it428-newsletter-api/services/subscription-service/internal/infrastructure/persistence/repositories"
	"4it428-newsletter-api/services/subscription-service/internal/service/services"
	"4it428-newsletter-api/services/subscription-service/internal/service/services/impl"
	"4it428-newsletter-api/services/subscription-service/internal/transport/api/v1/handler"
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
	"log"
	"os"
)

func SetupFirestore(ctx context.Context) (*firestore.Client, error) {
	firebaseSecretPath := os.Getenv("FIREBASE_CREDENTIALS")
	if firebaseSecretPath == "" {
		firebaseSecretPath = "../../secrets/firebase-adminsdk.json"
	}
	opt := option.WithCredentialsFile(firebaseSecretPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app with firebase admin sdk")
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

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
	emailProvider services.EmailProvider,
	newsletterServiceClient services.INewsletterServiceClient,
) *ServicesContainer {
	return &ServicesContainer{
		SubscriptionService: impl.NewSubscriptionService(subscriptionRepository, emailProvider, newsletterServiceClient),
	}
}
