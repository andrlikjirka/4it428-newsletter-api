package bootstrap

import (
	firebaseauth "4it428-newsletter-api/services/user-service/internal/infrastructure/firebase"
	"4it428-newsletter-api/services/user-service/internal/service/auth"
	"context"
	firebase "firebase.google.com/go/v4"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/api/option"
	"log"
	"os"
)

func SetupDatabase(ctx context.Context) (*pgxpool.Pool, error) {
	// Initialize the database connection pool.
	pool, err := pgxpool.New(
		ctx,
		os.Getenv("POSTGRES_DOCKER_URL"),
	)
	if err != nil {
		return nil, err
	}
	return pool, nil
}

func SetupFirebaseAuth(ctx context.Context) (auth.IAuthProvider, error) {
	firebaseSecretPath := os.Getenv("FIREBASE_CREDENTIALS")
	if firebaseSecretPath == "" {
		firebaseSecretPath = "./secrets/firebase-adminsdk.json"
	}
	opt := option.WithCredentialsFile(firebaseSecretPath)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app with firebase admin sdk")
		return nil, err
	}

	firebaseAPIKey := os.Getenv("FIREBASE_AUTH_API_KEY")
	authProvider, err := firebaseauth.NewFirebaseAuth(context.Background(), app, firebaseAPIKey)
	if err != nil {
		log.Fatalf("failed to initialize FirebaseAuthProvider: %v", err)
		return nil, err
	}
	return authProvider, nil
}
