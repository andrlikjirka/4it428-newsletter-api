package services

import (
	"4it428-newsletter-api/services/newsletter-service/internal/service/model"
	"context"
)

type PostService interface {
	CreatePost(ctx context.Context, post *model.Post, newsletterID string, userID string) (*model.Post, error)
	ListPosts(ctx context.Context, newsletterID string) ([]*model.Post, error)
	GetPostById(ctx context.Context, postID string, newsletterID string) (*model.Post, error)
	UpdatePost(ctx context.Context, postID string, newsletterID string, userID string, post *model.PostUpdate) (*model.Post, error)
	DeletePost(ctx context.Context, postID string, newsletterID string, userID string) error
	PublishPost(ctx context.Context, postID string, newsletterID string, userID string) error
}
