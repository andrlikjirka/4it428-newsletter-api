package bootstrap

import (
	firebaseauth "4it428-newsletter-api/services/user-service/internal/infrastructure/firebase"
	"4it428-newsletter-api/services/user-service/internal/infrastructure/persistence/repositories"
	"4it428-newsletter-api/services/user-service/internal/service/auth"
	"4it428-newsletter-api/services/user-service/internal/service/services"
	"4it428-newsletter-api/services/user-service/internal/transport/api/v1/handler"
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/api/option"
	"log"
	"os"
)

// SETUP METHODS:

func SetupDatabase(ctx context.Context) (*pgxpool.Pool, error) {
	dbURL := os.Getenv("POSTGRES_URL")
	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return nil, err
	}

	// force a connection to validate config and availability
	if err := pool.Ping(ctx); err != nil {
		pool.Close() // clean up if ping fails
		return nil, err
	}

	return pool, nil
}

func SetupFirebaseAuth(ctx context.Context) (auth.IAuthProvider, error) {
	firebaseSecretPath := os.Getenv("FIREBASE_CREDENTIALS")
	if firebaseSecretPath == "" {
		firebaseSecretPath = "../../secrets/firebase-adminsdk.json"
	}
	opt := option.WithCredentialsFile(firebaseSecretPath)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app with firebase admin sdk")
		return nil, err
	}

	firebaseAPIKey := os.Getenv("FIREBASE_AUTH_API_KEY")
	authProvider, err := firebaseauth.NewFirebaseAuth(ctx, app, firebaseAPIKey)
	if err != nil {
		log.Fatalf("failed to initialize FirebaseAuthProvider: %v", err)
		return nil, err
	}
	return authProvider, nil
}

// CONTAINERS FOR EASIER DI:

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
