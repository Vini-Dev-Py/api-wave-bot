package services

import (
	"api-wave-bot/internal/domain"
	"api-wave-bot/internal/infra/repository"
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo  repository.UserRepository
	JWTSecret string
}

// NewAuthService cria uma nova instância do AuthService.
func NewAuthService(userRepo repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		UserRepo:  userRepo,
		JWTSecret: jwtSecret,
	}
}

// Login valida as credenciais do usuário e retorna um JWT e o usuário.
func (s *AuthService) Login(ctx context.Context, email, password string) (string, *domain.User, error) {
	// Busca o usuário pelo email
	user, err := s.UserRepo.FindUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil, errors.New("usuário não encontrado")
		}
		return "", nil, err
	}

	// Valida a senha com bcrypt
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", nil, errors.New("senha inválida")
	}

	// Gerar JWT
	token, err := s.generateJWT(user)
	if err != nil {
		return "", nil, err
	}

	// Retorna o token e o usuário
	return token, user, nil
}

// Função para gerar o JWT
func (s *AuthService) generateJWT(user *domain.User) (string, error) {
	// Define os claims (informações contidas no JWT)
	claims := jwt.MapClaims{
		"id":        user.ID,
		"name":      user.Name,
		"email":     user.Email,
		"companyId": user.CompanyID,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // O token expira em 24h
	}

	// Cria o token com as informações e assina com o segredo
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.JWTSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
