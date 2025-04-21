package router

import (
	v1 "4it428-newsletter-api/services/user-service/internal/transport/api/v1/handler"
	"github.com/go-chi/chi/v5"
)

func NewAuthRouter(h *v1.AuthHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/signup", h.SignUp)
	r.Post("/signin", h.SignIn)
	r.Post("/social/signin", h.SocialSignIn) //optional: accepts provider (e.g. google/facebook) and idToken
	r.Post("/logout", h.Logout)
	r.Get("/verify", h.Verify)
	r.Post("/refresh", h.Refresh)

	return r
}
