package main

import (
	"api-wave-bot/internal/presentation/config"
	"api-wave-bot/internal/presentation/db"
	"api-wave-bot/internal/presentation/http/router"
	"log"
	"net/http"
)

func main() {
	// Carrega as configurações da aplicação
	cfg := config.Load()

	// Conecta no banco de dados
	conn, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Erro ao conectar com banco de dados: %v", err)
	}
	defer conn.Close()

	// Configura as rotas usando o roteador que criamos
	r := router.SetupRouter(conn)

	// Inicia o servidor HTTP
	log.Printf("Servidor rodando em http://localhost:%s", cfg.Port)
	err = http.ListenAndServe(":"+cfg.Port, r) // Usa o router configurado
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
