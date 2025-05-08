package bootstrap

import (
	"4it428-newsletter-api/services/newsletter-service/internal/infrastructure/persistence/repositories"
	"4it428-newsletter-api/services/newsletter-service/internal/service/services"
	"4it428-newsletter-api/services/newsletter-service/internal/service/services/impl"
	"4it428-newsletter-api/services/newsletter-service/internal/transport/api/v1/handler"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
)

func SetupDatabase(ctx context.Context) (*pgxpool.Pool, error) {
	dbURL := os.Getenv("POSTGRES_URL")
	fmt.Println(dbURL)
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, err
	}

	// force a connection to validate config and availability
	if err := pool.Ping(ctx); err != nil {
		pool.Close() // clean up if ping fails
		return nil, err
	}

	return pool, nil
}

// CONTAINERS FOR EASIER DI:

type HandlersContainer struct {
	NewsletterHandler *handler.NewsletterHandler
}

func NewHandlersContainer(s *ServicesContainer) *HandlersContainer {
	return &HandlersContainer{
		NewsletterHandler: handler.NewNewsletterHandler(s.NewsletterService),
	}
}

type ServicesContainer struct {
	NewsletterService services.NewsletterService
}

func NewServicesContainer(
	newsletterRepository *repositories.NewsletterRepository,
) *ServicesContainer {
	return &ServicesContainer{
		NewsletterService: impl.NewNewsletterService(newsletterRepository),
	}
}
