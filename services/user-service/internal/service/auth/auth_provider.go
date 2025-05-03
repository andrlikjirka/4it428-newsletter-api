package auth

import (
	"context"
	"firebase.google.com/go/v4/auth"
)

type IAuthProvider interface {
	SignUp(ctx context.Context, email string, password string) (*AuthProviderSignUpResponse, error)
	SignIn(ctx context.Context, email string, password string) (*AuthProviderSignInResponse, error)
	SendVerificationEmail(ctx context.Context, idToken string) error
	RefreshToken(ctx context.Context, refreshToken string) (*AuthProviderRefreshResponse, error)
	VerifyToken(ctx context.Context, idToken string) (*auth.Token, error)
}

type AuthProviderSignUpResponse struct {
	Email        string `json:"email"`
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalID      string `json:"localId"`
}

type AuthProviderSignInResponse struct {
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
	LocalID      string `json:"localId"`
}

type AuthProviderRefreshResponse struct {
	TokenType string `json:"token_type"`
	ExpiresIn string `json:"expires_in"`
	IDToken   string `json:"id_token"`
}
