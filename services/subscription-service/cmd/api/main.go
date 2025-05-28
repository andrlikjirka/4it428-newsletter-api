package main

import (
	"4it428-newsletter-api/pkg/logger"
	"4it428-newsletter-api/services/subscription-service/internal/bootstrap"
	"4it428-newsletter-api/services/subscription-service/internal/infrastructure/clients"
	"4it428-newsletter-api/services/subscription-service/internal/infrastructure/persistence/repositories"
	"4it428-newsletter-api/services/subscription-service/internal/transport/api"
	"context"
	"errors"
	"github.com/joho/godotenv"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var version = "v0.0.0"

func main() {
	_ = godotenv.Load("../../.env")
	ctx := context.Background()
	logger.Init()

	port := os.Getenv("SUBSCRIPTION_SERVICE_PORT")
	if port == "" {
		port = "8082"
	}

	firestoreClient, err := bootstrap.SetupFirestore(ctx)
	if err != nil {
		logger.Error("initializing firestore client failed", "error", err)
		return
	}
	defer firestoreClient.Close()

	awsSesClient, err := bootstrap.SetupAwsSes(ctx)
	if err != nil {
		logger.Error("initializing AWS SES client failed", "error", err)
		return
	}

	subscriptionRepo := repositories.NewSubscriptionRepository(firestoreClient)
	newsletterServiceClient := clients.NewNewsletterServiceClient()
	services := bootstrap.NewServicesContainer(subscriptionRepo, awsSesClient, newsletterServiceClient)
	handlers := bootstrap.NewHandlersContainer(services)
	router := api.NewApiRouter(handlers, version)
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Channel to listen for interrupt or terminate signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM) // forward os signals (syscalls) into the stop channel

	// Start server in a goroutine so it doesn't block
	go func() {
		logger.Info("Server is running at http://localhost:" + port)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("Server error", err)
		}
	}()

	// Block the main thread until a stop signal is received
	<-stop
	logger.Info("Shutting down server...")

	// Gracefully shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Error("Error during server shutdown", err)
	}

	logger.Info("Server gracefully stopped")
}
