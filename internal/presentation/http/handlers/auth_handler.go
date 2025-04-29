package handlers

import (
	"api-wave-bot/internal/presentation/http/services"
	"encoding/json"
	"net/http"
)

// AuthHandler é responsável por gerenciar as requisições de autenticação.
type AuthHandler struct {
	Service *services.AuthService
}

// NewAuthHandler cria uma nova instância de AuthHandler com o serviço injetado.
func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{
		Service: service,
	}
}

// LoginHandler é responsável por lidar com a requisição de login.
func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Chama o serviço para validar e gerar o JWT
	token, user, err := h.Service.Login(r.Context(), credentials.Email, credentials.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	// Mapeia apenas os dados seguros do usuário
	type UserResponse struct {
		ID        string `json:"id"`
		Name      string `json:"name"`
		Email     string `json:"email"`
		CompanyID string `json:"companyId"`
	}

	response := struct {
		Token string       `json:"token"`
		User  UserResponse `json:"user"`
	}{
		Token: token,
		User: UserResponse{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CompanyID: user.CompanyID,
		},
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
