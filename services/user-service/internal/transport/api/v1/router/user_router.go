package router

import (
	v1 "4it428-newsletter-api/services/user-service/internal/transport/api/v1/handler"
	"github.com/go-chi/chi/v5"
)

func NewUserRouter(h *v1.UserHandler) *chi.Mux {
	r := chi.NewRouter()

	// TODO: middleware

	r.Get("/", h.ListUsers)
	r.Post("/", h.CreateUser)
	r.Get("/{email}", h.GetUserByEmail)
	r.Put("/{email}", h.UpdateUser)
	r.Delete("/{email}", h.DeleteUser)

	return r
}
