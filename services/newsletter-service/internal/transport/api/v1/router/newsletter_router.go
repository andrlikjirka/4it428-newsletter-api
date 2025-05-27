package router

import (
	v1 "4it428-newsletter-api/services/newsletter-service/internal/transport/api/v1/handler"
	"github.com/go-chi/chi/v5"
)

func NewNewsletterRouter(h *v1.NewsletterHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", h.ListNewsletters)
	r.Post("/", h.CreateNewsletter)
	r.Get("/{id}", h.GetNewsletter)
	r.Put("/{id}", h.UpdateNewsletter)
	r.Delete("/{id}", h.DeleteNewsletter)

	return r
}
