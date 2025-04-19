package main

import (
	"api-wave-bot/internal/presentation/config"
	"api-wave-bot/internal/presentation/db"
	"log"
	"net/http"
)

func main() {
	cfg := config.Load()

	conn, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Erro ao conectar com banco de dados: %v", err)
	}
	defer conn.Close()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Wave Bot rodando! 🚀"))
	})

	log.Printf("Servidor rodando em http://localhost:%s", cfg.Port)
	err = http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
