package api

import (
	"4it428-newsletter-api/pkg/logger"
	"4it428-newsletter-api/services/subscription-service/internal/bootstrap"
	v1 "4it428-newsletter-api/services/subscription-service/internal/transport/api/v1"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func NewApiRouter(
	handlers *bootstrap.HandlersContainer,
	version string,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Get("/openapi.yaml", OpenAPIHandler)

		r.Mount("/v1", v1.NewV1Router(handlers))
	})

	// Health check routes
	r.Get("/ping", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	r.Get("/version", VersionHandler(version))

	return r
}

func OpenAPIHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./openapi.yaml")
}

func VersionHandler(version string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		_, err := w.Write([]byte(version))
		if err != nil {
			logger.Error("writing response", "error", err)
		}
	}
}
