package bootstrap

import (
	"4it428-newsletter-api/services/subscription-service/internal/infrastructure/persistence/repositories"
	"4it428-newsletter-api/services/subscription-service/internal/service/services"
	"4it428-newsletter-api/services/subscription-service/internal/service/services/impl"
	"4it428-newsletter-api/services/subscription-service/internal/transport/api/v1/handler"
)

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
) *ServicesContainer {
	return &ServicesContainer{
		SubscriptionService: impl.NewSubscriptionService(subscriptionRepository),
	}
}
