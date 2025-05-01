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
	response, err := a.authProvider.SignUp(ctx, input.Email, input.Password)
	if err != nil {
		logger.Error("AuthProvider user signup failed", "error", err)
		return err
	}
	logger.Info("User signed up with UID: " + response.LocalID)

	// Send email verification
	logger.Info("Sending verification email to: " + input.Email)
	if err := a.authProvider.SendVerificationEmail(ctx, response.IDToken); err != nil {
		logger.Error("Failed to send verification email", "error", err)
	}

	user := &model.User{
		ID:          uuid.New(),
		Email:       input.Email,
		FirebaseUID: response.LocalID,
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

func (a *authService) SignIn(ctx context.Context, email string, password string) (*auth.AuthProviderSignInResponse, error) {
	response, err := a.authProvider.SignIn(ctx, email, password)
	if err != nil {
		logger.Error("AuthProvider user sign in failed", "error", err)
		return nil, err
	}
	logger.Info("User signed in with UID: " + response.LocalID)
	return response, nil
}

func (a *authService) SocialSignIn(_ context.Context) {
	logger.Info("Social sign-in via %s with token: %s\n")
}

func (a *authService) Verify(_ context.Context) {
	logger.Info("Verifying user")
}

func (a *authService) RefreshToken(ctx context.Context, refreshToken string) (*auth.AuthProviderRefreshResponse, error) {
	response, err := a.authProvider.RefreshToken(ctx, refreshToken)
	if err != nil {
		logger.Error("AuthProvider token refresh failed", "error", err)
		return nil, err
	}
	logger.Info("Token refreshed successfully")
	return response, nil
}
