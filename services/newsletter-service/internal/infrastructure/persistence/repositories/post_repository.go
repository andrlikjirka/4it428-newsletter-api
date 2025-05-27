package repositories

import (
	dbmodel "4it428-newsletter-api/services/newsletter-service/internal/infrastructure/persistence/model"
	"4it428-newsletter-api/services/newsletter-service/internal/infrastructure/persistence/query"
	"4it428-newsletter-api/services/newsletter-service/internal/service/model"
	"context"
	"fmt"
	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepository struct {
	pool *pgxpool.Pool
}

func NewPostRepository(pool *pgxpool.Pool) *PostRepository {
	return &PostRepository{pool: pool}
}

func (r *PostRepository) Add(ctx context.Context, post *model.Post) (*model.Post, error) {
	_, err := r.pool.Exec(ctx, query.InsertPost, post.ID, post.NewsletterID, post.Title, post.Content, post.HtmlContent)
	if err != nil {
		return nil, fmt.Errorf("failed to insert post: %w", err)
	}
	return r.GetById(ctx, post.ID, post.NewsletterID)
}

func (r *PostRepository) List(ctx context.Context, newsletterID uuid.UUID) ([]*model.Post, error) {
	var posts []dbmodel.PostEntity
	err := pgxscan.Select(ctx, r.pool, &posts, query.SelectPosts, newsletterID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch post: %w", err)
	}

	response := make([]*model.Post, len(posts))
	for i, post := range posts {
		response[i] = &model.Post{
			ID:           post.ID,
			NewsletterID: post.NewsletterID,
			Title:        post.Title,
			Content:      post.Content,
			HtmlContent:  post.HtmlContent,
			Published:    post.Published,
			CreatedAt:    post.CreatedAt.Time,
			UpdatedAt:    post.UpdatedAt.Time,
		}
	}
	return response, nil
}

func (r *PostRepository) GetById(ctx context.Context, postID uuid.UUID, newsletterID uuid.UUID) (*model.Post, error) {
	var post dbmodel.PostEntity
	err := pgxscan.Get(ctx, r.pool, &post, query.SelectPostById, postID, newsletterID)
	if err != nil {
		return nil, err
	}

	return &model.Post{
		ID:           post.ID,
		NewsletterID: post.NewsletterID,
		Title:        post.Title,
		Content:      post.Content,
		HtmlContent:  post.HtmlContent,
		Published:    post.Published,
		CreatedAt:    post.CreatedAt.Time,
		UpdatedAt:    post.UpdatedAt.Time,
	}, nil
}

func (r *PostRepository) GetByIdAndUserId(ctx context.Context, postID uuid.UUID, newsletterID uuid.UUID, userID uuid.UUID) (*model.Post, error) {
	var post dbmodel.PostEntity
	err := pgxscan.Get(ctx, r.pool, &post, query.SelectPostByIdAndUserId, postID, newsletterID, userID)
	if err != nil {
		return nil, err
	}

	return &model.Post{
		ID:           post.ID,
		NewsletterID: post.NewsletterID,
		Title:        post.Title,
		Content:      post.Content,
		HtmlContent:  post.HtmlContent,
		Published:    post.Published,
		CreatedAt:    post.CreatedAt.Time,
		UpdatedAt:    post.UpdatedAt.Time,
	}, nil
}

func (r *PostRepository) Update(ctx context.Context, post *model.Post) (*model.Post, error) {
	_, err := r.pool.Exec(ctx, query.UpdatePost, post.Title, post.Content, post.HtmlContent, post.Published, post.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update post: %w", err)
	}

	return r.GetById(ctx, post.ID, post.NewsletterID)
}

func (r *PostRepository) Delete(ctx context.Context, postID uuid.UUID, newsletterID uuid.UUID) error {
	commandTag, err := r.pool.Exec(ctx, query.DeletePost, postID, newsletterID)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("post not found with id: %s", postID)
	}
	return nil
}
