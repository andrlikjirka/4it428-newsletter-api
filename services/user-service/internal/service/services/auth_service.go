package services

import (
	"4it428-newsletter-api/libs/logger"
	"context"
)

type authService struct {
}

func NewAuthService() IAuthService {
	return &authService{}
}

//TODO: add request parameters

func (a *authService) SignUp(_ context.Context) {
	logger.Info("Signing up user")
}

func (a *authService) SignIn(_ context.Context) {
	logger.Info("Signing in user")
}

func (a *authService) SocialSignIn(_ context.Context) {
	logger.Info("Social sign-in via %s with token: %s\n")
}

func (a *authService) Logout(_ context.Context) {
	logger.Info("Logging out user")
}

func (a *authService) Verify(_ context.Context) {
	logger.Info("Verifying user")
}

func (a *authService) Refresh(_ context.Context) {
	logger.Info("Refreshing token")
}
