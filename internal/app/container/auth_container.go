package container

import (
	"api-wave-bot/internal/infra/repository"
	"api-wave-bot/internal/presentation/http/handlers"
	"api-wave-bot/internal/presentation/http/middleware"
	"api-wave-bot/internal/presentation/http/services"
	"database/sql"

	"github.com/google/wire"
)

type AuthContainer struct {
	AuthHandler    *handlers.AuthHandler
	AuthMiddleware *middleware.AuthMiddleware
	JwtSecret      string
}

var AuthSet = wire.NewSet(
	NewUserRepository,
	NewAuthService,
	NewAuthHandler,
	NewAuthMiddleware,
	NewAuthContainer,
)

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return repository.NewUserRepository(db)
}

func NewAuthService(userRepo repository.UserRepository, jwtSecret string) *services.AuthService {
	return services.NewAuthService(userRepo, jwtSecret)
}

func NewAuthHandler(authService *services.AuthService) *handlers.AuthHandler {
	return handlers.NewAuthHandler(authService)
}

func NewAuthMiddleware(jwtSecret string) *middleware.AuthMiddleware {
	return middleware.NewAuthMiddleware(jwtSecret)
}

func NewAuthContainer(db *sql.DB, jwtSecret string) *AuthContainer {
	userRepository := NewUserRepository(db)

	authMiddleware := NewAuthMiddleware(jwtSecret)
	authService := NewAuthService(userRepository, jwtSecret)
	authHandler := NewAuthHandler(authService)

	return &AuthContainer{
		AuthHandler:    authHandler,
		AuthMiddleware: authMiddleware,
		JwtSecret:      jwtSecret,
	}
}
