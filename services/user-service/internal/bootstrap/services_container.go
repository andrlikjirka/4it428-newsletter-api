package bootstrap

import (
	"4it428-newsletter-api/services/user-service/internal/service/services"
)

type ServicesContainer struct {
	UserService services.IUserService
	AuthService services.IAuthService
}

func NewServicesContainer() *ServicesContainer {
	return &ServicesContainer{
		UserService: services.NewUserService(),
		AuthService: services.NewAuthService(),
	}
}
