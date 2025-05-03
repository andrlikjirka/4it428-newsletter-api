package firebase_auth

import (
	"4it428-newsletter-api/libs/logger"
	"4it428-newsletter-api/services/user-service/internal/service/auth"
	"bytes"
	"context"
	"encoding/json"
	firebase "firebase.google.com/go/v4"
	firebaseAuth "firebase.google.com/go/v4/auth"
	"fmt"
	"io"
	"net/http"
)

type FirebaseAuthProvider struct {
	apiKey string
	client *http.Client
	admin  *firebaseAuth.Client
}

func NewFirebaseAuth(ctx context.Context, app *firebase.App, apiKey string) (*FirebaseAuthProvider, error) {
	adminClient, err := app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize firebase auth client: %w", err)
	}

	return &FirebaseAuthProvider{
		apiKey: apiKey,
		client: &http.Client{},
		admin:  adminClient,
	}, nil
}

func (f *FirebaseAuthProvider) SignUp(ctx context.Context, email string, password string) (*auth.AuthProviderSignUpResponse, error) {
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

func (f *FirebaseAuthProvider) SignIn(ctx context.Context, email string, password string) (*auth.AuthProviderSignInResponse, error) {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:signInWithPassword?key=%s", f.apiKey)
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

	// Check if the response status is a 400 (INVALID_LOGIN_CREDENTIALS)
	if response.StatusCode == http.StatusBadRequest {
		return nil, fmt.Errorf("invalid email or password")
	}
	// If no errors, parse and return the result
	var result auth.AuthProviderSignInResponse
	if err := json.Unmarshal(resBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (f *FirebaseAuthProvider) SendVerificationEmail(ctx context.Context, idToken string) error {
	url := fmt.Sprintf("https://identitytoolkit.googleapis.com/v1/accounts:sendOobCode?key=%s", f.apiKey)
	payload := map[string]string{
		"requestType": "VERIFY_EMAIL",
		"idToken":     idToken,
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := f.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send verification email: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error sending verification email: %s", string(resBody))
	}
	return nil
}

func (f *FirebaseAuthProvider) RefreshToken(ctx context.Context, refreshToken string) (*auth.AuthProviderRefreshResponse, error) {
	url := fmt.Sprintf("https://securetoken.googleapis.com/v1/token?key=%s", f.apiKey)
	payload := map[string]interface{}{
		"grant_type":    "refresh_token",
		"refresh_token": refreshToken,
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

	// Check for the response status
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to refresh token: %s", string(resBody))
	}

	var result auth.AuthProviderRefreshResponse
	if err := json.Unmarshal(resBody, &result); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &result, nil
}

func (f *FirebaseAuthProvider) VerifyToken(ctx context.Context, idToken string) (*firebaseAuth.Token, error) {
	token, err := f.admin.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("failed to verify ID token: %w", err)
	}
	return token, nil
}

func (f *FirebaseAuthProvider) DeleteUser(ctx context.Context, uid string) error {
	err := f.admin.DeleteUser(ctx, uid)
	if err != nil {
		logger.Error("Error deleting user in firebase auth", "error", err)
		return err
	}
	logger.Info("User deleted from firebase auth with uid: " + uid)
	return nil
}
