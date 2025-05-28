package router

import (
	v1 "4it428-newsletter-api/services/subscription-service/internal/transport/api/v1/handler"
	"github.com/go-chi/chi/v5"
)

func NewSubscriptionRouter(h *v1.SubscriptionHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/{subscriptionID}/_unsubscribe", h.Unsubscribe)
	r.Post("/_notify", h.NotifySubscribers)
	r.Get("/", h.ListSubscriptions)
	r.Post("/", h.Subscribe)

	return r
}
