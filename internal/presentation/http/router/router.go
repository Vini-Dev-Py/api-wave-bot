package router

import (
	"api-wave-bot/internal/presentation/http/controllers"
	"api-wave-bot/internal/presentation/http/handlers"
	"api-wave-bot/internal/presentation/repository"
	"database/sql"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// SetupRouter configura as rotas da aplicação
func SetupRouter(conn *sql.DB) *mux.Router {
	r := mux.NewRouter()

	userRepository := repository.NewUserRepository(conn)
	authController := controllers.NewAuthController(userRepository, os.Getenv("JWT_SECRET"))
	authHandler := handlers.NewAuthHandler(authController)

	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Wave Bot rodando! 🚀"))
	}).Methods("GET")

	// Defina a rota para o login
	r.HandleFunc("/login", authHandler.LoginHandler).Methods("POST")

	// Defina outras rotas aqui

	return r
}
