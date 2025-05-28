package services

import (
	"context"
)

type Notification struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	HtmlContent string `json:"html_content"`
}

type ISubscriptionServiceClient interface {
	NotifySubscribers(ctx context.Context, newsletterID string, notification *Notification) error
}
