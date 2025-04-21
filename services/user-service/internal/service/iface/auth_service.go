package iface

import "context"

type IAuthService interface {
	SignUp(ctx context.Context)
	SignIn(ctx context.Context)
	SocialSignIn(ctx context.Context)
	Logout(ctx context.Context)
	Verify(ctx context.Context)
	Refresh(ctx context.Context)
}
