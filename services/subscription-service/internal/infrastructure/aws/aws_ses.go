package aws

import (
	"4it428-newsletter-api/pkg/logger"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

type SESClient struct {
	client *ses.Client
	sender string
}

func NewSESClient(ctx context.Context, sender string, region string) (*SESClient, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %w", err)
	}
	return &SESClient{
		client: ses.NewFromConfig(cfg),
		sender: sender,
	}, nil
}

func (s *SESClient) SendEmail(ctx context.Context, to string, subject string, textBody string, htmlBody string) error {
	input := &ses.SendEmailInput{
		Source: aws.String(s.sender),
		Destination: &types.Destination{
			ToAddresses: []string{to},
		},
		Message: &types.Message{
			Subject: &types.Content{
				Data: aws.String(subject),
			},
			Body: &types.Body{
				Text: &types.Content{
					Data: aws.String(textBody),
				},
				Html: &types.Content{
					Data: aws.String(htmlBody),
				},
			},
		},
	}

	result, err := s.client.SendEmail(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}
	logger.Info("Email sent", "to", to)
	logger.Debug("SES SendEmail result: %v", result)
	return nil
}
