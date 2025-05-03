package bootstrap

import (
	"4it428-newsletter-api/services/user-service/internal/infrastructure/persistence/repositories"
	"4it428-newsletter-api/services/user-service/internal/service/auth"
	"4it428-newsletter-api/services/user-service/internal/service/services"
)

type ServicesContainer struct {
	UserService services.IUserService
	AuthService services.IAuthService
}

func NewServicesContainer(
	userRepository *repositories.UserRepository,
	authProvider auth.IAuthProvider,
) *ServicesContainer {

	userService := services.NewUserService(authProvider, userRepository)
	authService := services.NewAuthService(authProvider, userService)

	return &ServicesContainer{
		UserService: userService,
		AuthService: authService,
	}
}
