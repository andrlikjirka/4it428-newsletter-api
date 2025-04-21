package main

import (
	"4it428-newsletter-api/libs/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"os"
)

func main() {
	logger.Init()
	port := os.Getenv("SUBSCRIPTION_SERVICE_PORT")
	if port == "" {
		port = "8082"
	}

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/api/v1/subscriptions", basicHandler)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	logger.Info("Server is running at http://localhost:" + port)
	err := server.ListenAndServe()
	if err != nil {
		logger.Error("Failed to listen to server.", err)
	}
}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World from Subscription Service!"))
	if err != nil {
		return
	}
}
