package bootstrap

import (
	"4it428-newsletter-api/services/user-service/internal/service/iface"
	"4it428-newsletter-api/services/user-service/internal/service/impl"
)

type ServicesContainer struct {
	UserService iface.IUserService
	AuthService iface.IAuthService
}

func NewServicesContainer() *ServicesContainer {
	return &ServicesContainer{
		UserService: impl.NewUserService(),
		AuthService: impl.NewAuthService(),
	}
}
