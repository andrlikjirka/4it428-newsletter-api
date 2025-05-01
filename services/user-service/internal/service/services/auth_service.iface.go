package services

import (
	"4it428-newsletter-api/services/user-service/internal/service/model"
	"context"
)

type IAuthService interface {
	SignUp(ctx context.Context, input *model.SignUpInput) error
	SignIn(ctx context.Context)
	SocialSignIn(ctx context.Context)
	Logout(ctx context.Context)
	Verify(ctx context.Context)
	Refresh(ctx context.Context)
}
