package services

import (
	"context"
)

type EmailProvider interface {
	SendEmail(ctx context.Context, to string, subject string, textBody string, htmlBody string) error
}
