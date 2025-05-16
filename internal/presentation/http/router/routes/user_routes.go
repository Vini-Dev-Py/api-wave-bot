package routes

import (
	"api-wave-bot/internal/app/container"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(r *mux.Router, userContainer *container.UserContainer, authContainer *container.AuthContainer) {
	api := r.PathPrefix("/api").Subrouter()
	api.Use(authContainer.AuthMiddleware.Middleware)

	api.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Aqui deve ficar a listagem dos meus usuários"))
	}).Methods("GET")
}
