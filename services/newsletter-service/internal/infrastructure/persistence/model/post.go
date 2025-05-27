package model

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type PostEntity struct {
	ID           uuid.UUID          `db:"id"`
	NewsletterID uuid.UUID          `db:"newsletter_id"`
	Title        string             `db:"title"`
	Content      string             `db:"content"`
	HtmlContent  string             `db:"html_content"`
	Published    bool               `db:"published"`
	CreatedAt    pgtype.Timestamptz `db:"created_at"`
	UpdatedAt    pgtype.Timestamptz `db:"updated_at"`
}
