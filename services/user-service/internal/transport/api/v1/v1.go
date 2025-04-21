package v1

import (
	"4it428-newsletter-api/services/user-service/internal/bootstrap"
	"4it428-newsletter-api/services/user-service/internal/transport/api/v1/router"
	"github.com/go-chi/chi/v5"
)

func NewV1Router(h *bootstrap.HandlersContainer) chi.Router {
	r := chi.NewRouter()

	r.Mount("/users", router.NewUserRouter(h.UserHandler))
	r.Mount("/auth", router.NewAuthRouter(h.AuthHandler))

	return r
}
