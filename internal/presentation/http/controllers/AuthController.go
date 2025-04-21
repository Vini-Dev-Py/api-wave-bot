package controllers

import (
	"api-wave-bot/internal/presentation/domain"
	"api-wave-bot/internal/presentation/repository"
	"context"
	"errors"
	"time"

	"database/sql"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	UserRepo  repository.UserRepository
	JWTSecret string // segredo para assinar o token
}

type UserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CompanyID string `json:"companyId"`
}

// NewAuthController cria uma nova instância do controller com o repositório injetado.
func NewAuthController(userRepo repository.UserRepository, jwtSecret string) *AuthController {
	return &AuthController{
		UserRepo:  userRepo,
		JWTSecret: jwtSecret,
	}
}

// Login recebe email e senha, valida o usuário e retorna um objeto de domínio e um JWT.
func (c *AuthController) Login(ctx context.Context, email, password string) (string, *UserResponse, error) {
	// Busca o usuário pelo email
	user, err := c.UserRepo.FindUserByEmail(ctx, email)
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
	token, err := c.generateJWT(user)
	if err != nil {
		return "", nil, err
	}

	userResp := &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CompanyID: user.CompanyID,
	}

	// Retorna o token e o usuário
	return token, userResp, nil
}

// Função para gerar o JWT
func (c *AuthController) generateJWT(user *domain.User) (string, error) {
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
	signedToken, err := token.SignedString([]byte(c.JWTSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
