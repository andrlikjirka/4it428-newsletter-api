package impl

import (
	"4it428-newsletter-api/pkg/logger"
	"4it428-newsletter-api/services/user-service/internal/service/auth"
	"4it428-newsletter-api/services/user-service/internal/service/errors"
	"4it428-newsletter-api/services/user-service/internal/service/model"
	"4it428-newsletter-api/services/user-service/internal/service/services"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type authService struct {
	authProvider auth.IAuthProvider
	userService  services.UserService
}

func NewAuthService(authProvider auth.IAuthProvider, service services.UserService) services.AuthService {
	return &authService{
		authProvider: authProvider,
		userService:  service,
	}
}

func (a *authService) SignUp(ctx context.Context, input *model.SignUpInput) (*auth.AuthProviderSignUpResponse, error) {
	response, err := a.authProvider.SignUp(ctx, input.Email, input.Password)
	if err != nil {
		logger.Error("AuthProvider user signup failed", "error", err)
		return nil, err
	}
	logger.Info("User signed up with UID: " + response.LocalID)

	// Send verification email
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
		return nil, err
	}
	logger.Info("Local user record created successfully for email: " + input.Email)
	return response, nil
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

func (a *authService) RefreshToken(ctx context.Context, refreshToken string) (*auth.AuthProviderRefreshResponse, error) {
	response, err := a.authProvider.RefreshToken(ctx, refreshToken)
	if err != nil {
		logger.Error("AuthProvider token refresh failed", "error", err)
		return nil, err
	}
	logger.Info("Token refreshed successfully")
	return response, nil
}

func (a *authService) Verify(ctx context.Context, idToken string) (map[string]interface{}, error) {
	token, err := a.authProvider.VerifyToken(ctx, idToken)
	if err != nil {
		logger.Error("Token verification failed", "error", err)
		return nil, err
	}
	logger.Info("Token successfully verified", "email", token.Claims["email"])

	userEmail, ok := token.Claims["email"].(string)
	if !ok {
		return nil, fmt.Errorf("user_id not found in token claims")
	}

	// fetch user from DB
	user, err := a.userService.GetUserByEmail(ctx, userEmail)
	if err != nil {
		return nil, errors.ErrUserNotFound
	}

	// Compose enriched claims
	claims := map[string]interface{}{
		"user_id":        user.ID,
		"email":          user.Email,
		"email_verified": token.Claims["email_verified"],
		//"role":    user.Role, // e.g., "editor"
	}

	return claims, nil
}
