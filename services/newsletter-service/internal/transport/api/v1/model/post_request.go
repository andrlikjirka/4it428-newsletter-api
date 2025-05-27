package model

import (
	"4it428-newsletter-api/services/newsletter-service/internal/service/model"
	"github.com/google/uuid"
)

type CreatePostRequest struct {
	Title       string `json:"title" validate:"required,min=3,max=255"`
	Content     string `json:"content" validate:"required"`
	HtmlContent string `json:"html_content" validate:"required"`
}

type UpdatePostRequest struct {
	Title       *string `json:"title,omitempty" validate:"omitempty,min=3,max=255"`
	Content     *string `json:"content,omitempty" validate:"omitempty"`
	HtmlContent *string `json:"html_content,omitempty" validate:"omitempty"`
}

func (n *CreatePostRequest) ToPost() *model.Post {
	return &model.Post{
		ID:          uuid.New(),
		Title:       n.Title,
		Content:     n.Content,
		HtmlContent: n.HtmlContent,
	}
}

func (n *UpdatePostRequest) ToPostUpdate() *model.PostUpdate {
	return &model.PostUpdate{
		Title:       n.Title,
		Content:     n.Content,
		HtmlContent: n.HtmlContent,
	}
}
