package model

import (
	"4it428-newsletter-api/services/newsletter-service/internal/service/model"
	"github.com/google/uuid"
	"time"
)

type PostResponse struct {
	ID           uuid.UUID `json:"id"`
	NewsletterID uuid.UUID `json:"newsletter_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	HtmlContent  string    `json:"html_content"`
	Published    bool      `json:"published"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func FromPost(u *model.Post) *PostResponse {
	return &PostResponse{
		ID:           u.ID,
		NewsletterID: u.NewsletterID,
		Title:        u.Title,
		Content:      u.Content,
		HtmlContent:  u.HtmlContent,
		Published:    u.Published,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}

func FromPostList(posts []*model.Post) []*PostResponse {
	responses := make([]*PostResponse, 0, len(posts))
	for _, u := range posts {
		responses = append(responses, FromPost(u))
	}
	return responses
}
