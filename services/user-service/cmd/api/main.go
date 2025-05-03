package main

import (
	"4it428-newsletter-api/libs/logger"
	"4it428-newsletter-api/services/user-service/internal/bootstrap"
	"4it428-newsletter-api/services/user-service/internal/infrastructure/persistence/repositories"
	"4it428-newsletter-api/services/user-service/internal/transport/api"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var version = "v0.0.0"

func main() {
	ctx := context.Background()
	logger.Init()

	port := os.Getenv("USER_SERVICE_PORT")
	if port == "" {
		port = "8083"
	}

	db, err := bootstrap.SetupDatabase(ctx)
	if err != nil {
		logger.Error("initializing database failed", "error", err)
	}
	defer db.Close()

	authProvider, err := bootstrap.SetupFirebaseAuth(ctx)
	if err != nil {
		logger.Error("initializing firebase auth failed", "error", err)
		return
	}

	userRepo := repositories.NewUserRepository(db)
	services := bootstrap.NewServicesContainer(userRepo, authProvider)
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

	db.Close()
	logger.Info("Database pool closed")

	logger.Info("Server gracefully stopped")
}
