package services

import (
	"4it428-newsletter-api/services/user-service/internal/service/auth"
	"4it428-newsletter-api/services/user-service/internal/service/model"
	"context"
)

type IAuthService interface {
	SignUp(ctx context.Context, input *model.SignUpInput) (*auth.AuthProviderSignUpResponse, error)
	SignIn(ctx context.Context, email, password string) (*auth.AuthProviderSignInResponse, error)
	SocialSignIn(ctx context.Context)
	Verify(ctx context.Context, idToken string) (map[string]interface{}, error)
	RefreshToken(ctx context.Context, refreshToken string) (*auth.AuthProviderRefreshResponse, error)
}
