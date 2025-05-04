package repositories

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NewsletterRepository struct {
	pool *pgxpool.Pool
}

func NewNewsletterRepository(pool *pgxpool.Pool) *NewsletterRepository {
	return &NewsletterRepository{pool: pool}
}

func (r *NewsletterRepository) List(ctx context.Context) {
}

func (r *NewsletterRepository) GetById(ctx context.Context) {
}

func (r *NewsletterRepository) Create(ctx context.Context) {
}

func (r *NewsletterRepository) Update(ctx context.Context) {
}

func (r *NewsletterRepository) Delete(ctx context.Context) {
}
