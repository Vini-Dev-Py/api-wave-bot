package routes

import (
	"api-wave-bot/internal/app/container"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterAuthRoutes configura as rotas de autenticação
func RegisterAuthRoutes(r *mux.Router, authContainer *container.AuthContainer) {
	// Rotas públicas
	r.HandleFunc("/login", authContainer.AuthHandler.LoginHandler).Methods("POST")

	// Rotas privadas (com Middleware)
	api := r.PathPrefix("/api").Subrouter()
	api.Use(authContainer.AuthMiddleware.Middleware)

	api.HandleFunc("/me", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Minha rota privada! 🔒"))
	}).Methods("GET")
}
