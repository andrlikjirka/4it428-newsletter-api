package v1

import (
	"4it428-newsletter-api/services/newsletter-service/internal/bootstrap"
	"4it428-newsletter-api/services/newsletter-service/internal/transport/api/v1/router"
	"github.com/go-chi/chi/v5"
)

func NewV1Router(h *bootstrap.HandlersContainer) chi.Router {
	r := chi.NewRouter()

	r.Mount("/newsletters", router.NewUserRouter(h.NewsletterHandler))

	return r
}
