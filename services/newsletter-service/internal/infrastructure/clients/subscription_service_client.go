package clients

import (
	"4it428-newsletter-api/services/newsletter-service/internal/service/services"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type SubscriptionServiceClient struct{}

func NewSubscriptionServiceClient() *SubscriptionServiceClient {
	return &SubscriptionServiceClient{}
}

func (c *SubscriptionServiceClient) NotifySubscribers(ctx context.Context, newsletterID string, notification *services.Notification) error {
	baseUrl := os.Getenv("SUBSCRIPTION_SERVICE_URL")
	port := os.Getenv("SUBSCRIPTION_SERVICE_PORT")
	url := fmt.Sprintf("%s:%s/api/v1/subscriptions/_notify?newsletter_id=%s", baseUrl, port, newsletterID)

	body, err := json.Marshal(notification)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		resBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to notify subscribers: %s", string(resBody))
	}

	return nil
}
