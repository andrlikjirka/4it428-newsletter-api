package model

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type NewsletterEntity struct {
	ID          uuid.UUID          `db:"id"`
	Title       string             `db:"title"`
	Description string             `db:"description"`
	CreatedAt   pgtype.Timestamptz `db:"created_at"`
	UpdatedAt   pgtype.Timestamptz `db:"updated_at"`
	Posts       []PostEntity       `db:"-" json:"posts,omitempty"`
	UserID      uuid.UUID          `db:"user_id"`
}
