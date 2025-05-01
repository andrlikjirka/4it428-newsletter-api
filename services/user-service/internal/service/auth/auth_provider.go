package auth

import "context"

type IAuthProvider interface {
	CreateUser(ctx context.Context, email string, password string) (*AuthProviderSignUpResponse, error)
}
