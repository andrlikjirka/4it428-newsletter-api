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
	"time"
)

type NewsletterRepository struct {
	pool *pgxpool.Pool
}

func NewNewsletterRepository(pool *pgxpool.Pool) *NewsletterRepository {
	return &NewsletterRepository{pool: pool}
}

func (r *NewsletterRepository) Add(ctx context.Context, newsletter *model.Newsletter) (*model.Newsletter, error) {
	_, err := r.pool.Exec(ctx, query.InsertNewsletter, newsletter.ID, newsletter.Title, newsletter.Description, newsletter.UserID)
	if err != nil {
		return nil, fmt.Errorf("failed to insert newsletter: %w", err)
	}
	return r.GetById(ctx, newsletter.ID)
}

func (r *NewsletterRepository) List(ctx context.Context) ([]*model.Newsletter, error) {
	var newsletters []dbmodel.NewsletterEntity
	err := pgxscan.Select(ctx, r.pool, &newsletters, query.SelectNewsletters)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch newsletters: %w", err)
	}

	response := make([]*model.Newsletter, len(newsletters))
	for i, newsletter := range newsletters {
		response[i] = &model.Newsletter{
			ID:          newsletter.ID,
			Title:       newsletter.Title,
			Description: newsletter.Description,
			CreatedAt:   newsletter.CreatedAt.Time,
			UpdatedAt:   newsletter.UpdatedAt.Time,
			UserID:      newsletter.UserID,
		}
	}
	return response, nil
}

func (r *NewsletterRepository) GetById(ctx context.Context, id uuid.UUID) (*model.Newsletter, error) {
	var newsletter dbmodel.NewsletterEntity
	err := pgxscan.Get(ctx, r.pool, &newsletter, query.SelectNewsletterById, id)
	if err != nil {
		return nil, err
	}

	return &model.Newsletter{
		ID:          newsletter.ID,
		Title:       newsletter.Title,
		Description: newsletter.Description,
		CreatedAt:   newsletter.CreatedAt.Time,
		UpdatedAt:   newsletter.UpdatedAt.Time,
		UserID:      newsletter.UserID,
	}, nil
}

func (r *NewsletterRepository) Update(ctx context.Context, newsletter *model.Newsletter) (*model.Newsletter, error) {
	now := time.Now().UTC()

	_, err := r.pool.Exec(ctx, query.UpdateNewsletter, newsletter.Title, newsletter.Description, now, newsletter.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to update newsletter: %w", err)
	}

	return r.GetById(ctx, newsletter.ID)
}

func (r *NewsletterRepository) Delete(ctx context.Context, id uuid.UUID, userID uuid.UUID) error {
	commandTag, err := r.pool.Exec(ctx, query.DeleteNewsletter, id, userID)
	if err != nil {
		return err
	}

	if commandTag.RowsAffected() == 0 {
		return fmt.Errorf("newsletter not found with id: %s", id)
	}
	return nil
}
