package services

import (
	"4it428-newsletter-api/libs/logger"
	"4it428-newsletter-api/services/user-service/internal/service/auth"
	"4it428-newsletter-api/services/user-service/internal/service/model"
	"context"
	"github.com/google/uuid"
)

type authService struct {
	authProvider auth.IAuthProvider
	userService  IUserService
}

func NewAuthService(authProvider auth.IAuthProvider, service IUserService) IAuthService {
	return &authService{
		authProvider: authProvider,
		userService:  service,
	}
}

func (a *authService) SignUp(ctx context.Context, input *model.SignUpInput) error {
	providerResponse, err := a.authProvider.CreateUser(ctx, input.Email, input.Password)
	if err != nil {
		logger.Error("AuthProvider user signup failed", "error", err)
		return err
	}
	logger.Info("User signed up with UID: " + providerResponse.LocalID)

	user := &model.User{
		ID:          uuid.New(),
		Email:       input.Email,
		FirebaseUID: providerResponse.LocalID,
		FirstName:   input.FirstName,
		LastName:    input.LastName,
	}
	if err := a.userService.CreateUser(ctx, user); err != nil {
		logger.Error("Failed to create user locally.", "error", err)
		return err
	}
	logger.Info("Local user record created successfully for email: " + input.Email)
	return nil
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
