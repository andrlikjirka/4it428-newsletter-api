package router

import (
	v1 "4it428-newsletter-api/services/newsletter-service/internal/transport/api/v1/handler"
	"github.com/go-chi/chi/v5"
)

func NewPostRouter(h *v1.PostHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", h.ListPosts)
	r.Post("/", h.CreatePost)
	r.Get("/{postID}", h.GetPost)
	r.Put("/{postID}", h.UpdatePost)
	r.Delete("/{postID}", h.DeletePost)

	r.Get("/{postID}/_publish", h.PublishPost)

	return r
}
