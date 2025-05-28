package clients

import (
	errors2 "4it428-newsletter-api/services/subscription-service/internal/service/errors"
	"4it428-newsletter-api/services/subscription-service/internal/service/services"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type NewsletterServiceClient struct{}

func NewNewsletterServiceClient() *NewsletterServiceClient {
	return &NewsletterServiceClient{}
}

func (c *NewsletterServiceClient) GetNewsletter(ctx context.Context, newsletterID string) (*services.Newsletter, error) {
	baseUrl := os.Getenv("NEWSLETTER_SERVICE_URL")
	port := os.Getenv("NEWSLETTER_SERVICE_PORT")
	url := fmt.Sprintf("%s:%s/api/v1/newsletters/%s", baseUrl, port, newsletterID)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return nil, errors2.ErrNewsletterNotFound
	}

	var newsletter services.Newsletter
	if err := json.Unmarshal(resBody, &newsletter); err != nil {
		return nil, err
	}
	return &newsletter, nil
}
