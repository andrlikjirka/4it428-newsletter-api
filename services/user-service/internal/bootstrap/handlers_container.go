package bootstrap

import "4it428-newsletter-api/services/user-service/internal/transport/api/v1/handler"

type HandlersContainer struct {
	UserHandler *handler.UserHandler
	AuthHandler *handler.AuthHandler
}

func NewHandlersContainer(s *ServicesContainer) *HandlersContainer {
	return &HandlersContainer{
		UserHandler: handler.NewUserHandler(s.UserService),
		AuthHandler: handler.NewAuthHandler(s.AuthService),
	}
}
