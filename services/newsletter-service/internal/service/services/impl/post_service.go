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

type postService struct {
	postRepo       repositories.IPostRepository
	newsletterRepo repositories.INewsletterRepository
}

func NewPostService(postRepo repositories.IPostRepository, newsletterRepo repositories.INewsletterRepository) services.PostService {
	return &postService{
		postRepo:       postRepo,
		newsletterRepo: newsletterRepo,
	}
}

func (p postService) CreatePost(ctx context.Context, post *model.Post, newsletterID string, userID string) (*model.Post, error) {
	parsedNewsletterID, err := utils.ParseUUID(newsletterID)
	if err != nil {
		logger.Error("Failed to parse UUID", "newsletterID", newsletterID, "error", err)
		return nil, errors2.ErrInvalidUUID
	}

	parsedUserID, err := utils.ParseUUID(userID)
	if err != nil {
		logger.Error("Failed to parse UUID", "userID", userID, "error", err)
		return nil, errors2.ErrInvalidUUID
	}

	_, err = p.newsletterRepo.GetByIdAndUserId(ctx, parsedNewsletterID, parsedUserID)
	if err != nil {
		logger.Error("Failed to get newsletter by ID", "newsletterID", newsletterID, "error", err)
		return nil, errors2.ErrNotFound
	}

	post.NewsletterID = parsedNewsletterID

	createdPost, err := p.postRepo.Add(ctx, post)
	if err != nil {
		logger.Error("Failed to create post", "error", err)
		return nil, err
	}
	logger.Info("Creating new post", "postID", post.ID.String())
	return createdPost, nil
}

func (p postService) ListPosts(ctx context.Context, newsletterID string) ([]*model.Post, error) {
	parsedNewsletterID, err := utils.ParseUUID(newsletterID)
	if err != nil {
		logger.Error("Failed to parse UUID", "newsletterID", newsletterID, "error", err)
		return nil, errors2.ErrInvalidUUID
	}

	_, err = p.newsletterRepo.GetById(ctx, parsedNewsletterID)
	if err != nil {
		logger.Error("Failed to get newsletter by ID", "newsletterID", newsletterID, "error", err)
		return nil, errors2.ErrNotFound
	}

	posts, err := p.postRepo.List(ctx, parsedNewsletterID)
	if err != nil {
		return nil, err
	}
	logger.Info("Listing all posts")
	return posts, nil
}

func (p postService) GetPostById(ctx context.Context, postID string, newsletterID string) (*model.Post, error) {
	parsedPostID, err := utils.ParseUUID(postID)
	if err != nil {
		logger.Error("Failed to parse UUID", "postID", postID, "error", err)
		return nil, errors2.ErrInvalidUUID
	}
	parsedNewsletterID, err := utils.ParseUUID(newsletterID)
	if err != nil {
		logger.Error("Failed to parse UUID", "newsletterID", newsletterID, "error", err)
		return nil, errors2.ErrInvalidUUID
	}

	post, err := p.postRepo.GetById(ctx, parsedPostID, parsedNewsletterID)
	if err != nil {
		return nil, errors2.ErrPostNotFound
	}

	logger.Info("Getting post", "postID", postID)
	return post, nil
}

func (p postService) UpdatePost(ctx context.Context, postID string, newsletterID string, userID string, postUpdate *model.PostUpdate) (*model.Post, error) {
	parsedPostID, err := utils.ParseUUID(postID)
	if err != nil {
		logger.Error("Failed to parse UUID", "postID", postID, "error", err)
		return nil, err
	}
	parsedNewsletterID, err := utils.ParseUUID(newsletterID)
	if err != nil {
		logger.Error("Failed to parse UUID", "newsletterID", newsletterID, "error", err)
		return nil, err
	}
	parsedUserID, err := utils.ParseUUID(userID)
	if err != nil {
		logger.Error("Failed to parse UUID", "userID", userID, "error", err)
		return nil, errors2.ErrInvalidUUID
	}

	post, err := p.postRepo.GetByIdAndUserId(ctx, parsedPostID, parsedNewsletterID, parsedUserID)
	if err != nil {
		logger.Error("Failed to get post by ID", "postID", postID, "error", err)
		return nil, errors2.ErrPostNotFound
	}

	if postUpdate.Title != nil {
		post.Title = *postUpdate.Title
	}
	if postUpdate.Content != nil {
		post.Content = *postUpdate.Content
	}
	if postUpdate.HtmlContent != nil {
		post.HtmlContent = *postUpdate.HtmlContent
	}

	updatedPost, err := p.postRepo.Update(ctx, post)
	if err != nil {
		logger.Error("Failed to update post", "postID", postID, "error", err)
		return nil, err
	}

	logger.Info("Post updated successfully", "postID", postID)
	return updatedPost, nil
}

func (p postService) DeletePost(ctx context.Context, postID string, newsletterID string, userID string) error {
	parsedPostID, err := utils.ParseUUID(postID)
	if err != nil {
		logger.Error("Failed to parse UUID", "postID", postID, "error", err)
		return errors2.ErrInvalidUUID
	}
	parsedNewsletterID, err := utils.ParseUUID(newsletterID)
	if err != nil {
		logger.Error("Failed to parse UUID", "newsletterID", newsletterID, "error", err)
		return errors2.ErrInvalidUUID
	}
	parsedUserID, err := utils.ParseUUID(userID)
	if err != nil {
		logger.Error("Failed to parse UUID", "userID", userID, "error", err)
		return errors2.ErrInvalidUUID
	}

	_, err = p.postRepo.GetByIdAndUserId(ctx, parsedPostID, parsedNewsletterID, parsedUserID)
	if err != nil {
		return errors2.ErrPostNotFound
	}

	err = p.postRepo.Delete(ctx, parsedPostID, parsedNewsletterID)
	if err != nil {
		logger.Error("Failed to delete post", "postID", postID, "error", err)
		return err
	}
	logger.Info("Post deleted successfully", "postID", postID)
	return nil
}

func (p postService) PublishPost(ctx context.Context, postID string, newsletterID string, userID string) error {
	parsedPostID, err := utils.ParseUUID(postID)
	if err != nil {
		logger.Error("Failed to parse UUID", "postID", postID, "error", err)
		return errors2.ErrInvalidUUID
	}
	parsedNewsletterID, err := utils.ParseUUID(newsletterID)
	if err != nil {
		logger.Error("Failed to parse UUID", "newsletterID", newsletterID, "error", err)
		return errors2.ErrInvalidUUID
	}
	parsedUserID, err := utils.ParseUUID(userID)
	if err != nil {
		logger.Error("Failed to parse UUID", "userID", userID, "error", err)
		return errors2.ErrInvalidUUID
	}

	post, err := p.postRepo.GetByIdAndUserId(ctx, parsedPostID, parsedNewsletterID, parsedUserID)
	if err != nil {
		logger.Error("Failed to get post by ID", "postID", postID, "error", err)
		return errors2.ErrPostNotFound
	}

	if post.Published {
		logger.Error("Post already published", "postID", postID)
		return errors2.ErrAlreadyPublished
	}

	//TODO: call subscription service to notify subscribers

	post.Published = true
	_, err = p.postRepo.Update(ctx, post)

	if err != nil {
		logger.Error("Failed to publish post", "postID", postID, "error", err)
		return err
	}
	logger.Info("Post published successfully", "postID", postID)
	return nil
}
