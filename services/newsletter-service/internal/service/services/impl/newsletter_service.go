package impl

import (
	"4it428-newsletter-api/libs/logger"
	"4it428-newsletter-api/libs/utils"
	errors2 "4it428-newsletter-api/services/newsletter-service/internal/service/errors"
	"4it428-newsletter-api/services/newsletter-service/internal/service/model"
	"4it428-newsletter-api/services/newsletter-service/internal/service/repositories"
	"4it428-newsletter-api/services/newsletter-service/internal/service/services"
	"context"
)

type newsletterService struct {
	repo repositories.INewsletterRepository
}

func NewNewsletterService(repo repositories.INewsletterRepository) services.NewsletterService {
	return &newsletterService{repo: repo}
}

func (n newsletterService) CreateNewsletter(ctx context.Context, newsletter *model.Newsletter) (*model.Newsletter, error) {
	createdNewsletter, err := n.repo.Add(ctx, newsletter)
	if err != nil {
		logger.Error("Failed to create newsletter", "error", err)
		return nil, err
	}
	logger.Info("Creating new newsletter", "id", newsletter.ID.String())
	return createdNewsletter, nil
}

func (n newsletterService) ListNewsletters(ctx context.Context) ([]*model.Newsletter, error) {
	newsletters, err := n.repo.List(ctx)
	if err != nil {
		return nil, err
	}
	logger.Info("Listing all newsletters")
	return newsletters, nil
}

func (n newsletterService) GetNewsletterById(ctx context.Context, id string) (*model.Newsletter, error) {
	parsedID, err := utils.ParseUUID(id)
	if err != nil {
		logger.Error("Failed to parse UUID", "id", id, "error", err)
		return nil, errors2.ErrInvalidUUID
	}

	newsletter, err := n.repo.GetById(ctx, parsedID)
	if err != nil {
		return nil, errors2.ErrNotFound
	}
	logger.Info("Getting newsletter", "id", id)
	return newsletter, nil
}

func (n newsletterService) UpdateNewsletter(ctx context.Context, id string, newsletterUpdate *model.NewsletterUpdate) (*model.Newsletter, error) {
	parsedID, err := utils.ParseUUID(id)
	if err != nil {
		logger.Error("Failed to parse UUID", "id", id, "error", err)
		return nil, err
	}

	newsletter, err := n.repo.GetById(ctx, parsedID)
	if err != nil {
		return nil, errors2.ErrNotFound
	}

	if newsletterUpdate.Title != nil {
		newsletter.Title = *newsletterUpdate.Title
	}
	if newsletterUpdate.Description != nil {
		newsletter.Description = *newsletterUpdate.Description
	}

	updatedNewsletter, err := n.repo.Update(ctx, newsletter)
	if err != nil {
		logger.Error("Failed to update newsletter", "id", id, "error", err)
		return nil, err
	}

	logger.Info("Newsletter updated successfully", "id", id)
	return updatedNewsletter, nil
}

func (n newsletterService) DeleteNewsletter(ctx context.Context, id string) error {
	parsedID, err := utils.ParseUUID(id)
	if err != nil {
		logger.Error("Failed to parse UUID", "id", id, "error", err)
		return errors2.ErrInvalidUUID
	}

	_, err = n.repo.GetById(ctx, parsedID)
	if err != nil {
		return errors2.ErrNotFound
	}

	err = n.repo.Delete(ctx, parsedID)
	if err != nil {
		logger.Error("Failed to delete newsletter", "id", id, "error", err)
		return err
	}
	logger.Info("Newsletter deleted successfully", "id", id)
	return nil
}
