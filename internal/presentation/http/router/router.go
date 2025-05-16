package router

import (
	"api-wave-bot/internal/app/container"
	"api-wave-bot/internal/presentation/http/router/routes"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRouter configura as rotas principais da aplicação
func SetupRouter(authContainer *container.AuthContainer, userContainer *container.UserContainer) *mux.Router {
	router := mux.NewRouter()

	// Health Check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Wave Bot rodando! 🚀"))
	}).Methods("GET")

	// Rotas de autenticação
	routes.RegisterAuthRoutes(router, authContainer)
	routes.RegisterUserRoutes(router, userContainer, authContainer)

	return router
}
