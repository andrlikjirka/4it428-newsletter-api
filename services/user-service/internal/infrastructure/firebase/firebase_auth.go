package firebase

import (
	"4it428-newsletter-api/services/user-service/internal/service/auth"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type FirebaseAuth struct {
	apiKey string
	client *http.Client
}

func NewFirebaseAuth(apiKey string) *FirebaseAuth {
	return &FirebaseAuth{
		apiKey: apiKey,
		client: &http.Client{},
	}
}

func (f *FirebaseAuth) CreateUser(ctx context.Context, email string, password string) (*auth.AuthProviderSignUpResponse, error) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signUp?key=%s", f.apiKey)
	payload := map[string]interface{}{
		"email":             email,
		"password":          password,
		"returnSecureToken": true,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	response, err := f.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var result auth.AuthProviderSignUpResponse
	if err := json.Unmarshal(resBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}
