package handler

import (
	"4it428-newsletter-api/services/newsletter-service/internal/service/services"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type NewsletterHandler struct {
	newsletterService services.NewsletterService
}

func NewNewsletterHandler(s services.NewsletterService) *NewsletterHandler {
	return &NewsletterHandler{newsletterService: s}
}

func (h *NewsletterHandler) ListNewsletters(w http.ResponseWriter, r *http.Request) {
}

func (h *NewsletterHandler) GetNewsletter(w http.ResponseWriter, r *http.Request) {
}

func (h *NewsletterHandler) CreateNewsletter(w http.ResponseWriter, r *http.Request) {
}

func (h *NewsletterHandler) UpdateNewsletter(w http.ResponseWriter, r *http.Request) {
}

func (h *NewsletterHandler) DeleteNewsletter(w http.ResponseWriter, r *http.Request) {
}
