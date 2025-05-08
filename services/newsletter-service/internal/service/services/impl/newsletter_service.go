package impl

import (
	"4it428-newsletter-api/services/newsletter-service/internal/service/repositories"
	"4it428-newsletter-api/services/newsletter-service/internal/service/services"
)

type newsletterService struct {
	repo repositories.INewsletterRepository
}

func NewNewsletterService(repo repositories.INewsletterRepository) services.NewsletterService {
	return &newsletterService{repo: repo}
}
