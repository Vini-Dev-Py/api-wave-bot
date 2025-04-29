package db

import (
	"api-wave-bot/internal/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func Connect(cfg config.Config) (*sql.DB, error) {
	psql := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName,
	)

	log.Printf("Tentando conectar ao banco de dados com a URL: %s", psql) // Adicionando log de depuração

	db, err := sql.Open("postgres", psql)
	if err != nil {
		log.Println("Erro ao abrir a conexão com o banco:", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println("Erro ao fazer ping no banco:", err)
		return nil, err
	}

	log.Println("Conexão com o banco de dados bem-sucedida!")
	return db, nil
}
