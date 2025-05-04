package services

import "4it428-newsletter-api/services/newsletter-service/internal/service/repositories"

type newsletterService struct {
	repo repositories.INewsletterRepository
}

func NewNewsletterService(repo repositories.INewsletterRepository) INewsletterService {
	return &newsletterService{repo: repo}
}
