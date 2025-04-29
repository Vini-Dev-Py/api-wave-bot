package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port   string
	DBHost string
	DBPort string
	DBUser string
	DBPass string
	DBName string
}

func Load() Config {
	// Carregar o arquivo .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis do sistema.")
	}

	// Carregar e logar as variáveis de ambiente
	port := getEnv("PORT", "8080")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPass := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "wavebot")

	// log.Printf("Variáveis de ambiente carregadas:\nPORT: %s\nDB_HOST: %s\nDB_PORT: %s\nDB_USER: %s\nDB_NAME: %s", port, dbHost, dbPort, dbUser, dbName)

	return Config{
		Port:   port,
		DBHost: dbHost,
		DBPort: dbPort,
		DBUser: dbUser,
		DBPass: dbPass,
		DBName: dbName,
	}
}

func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		log.Printf("Variável de ambiente %s não encontrada, usando valor padrão: %s", key, fallback)
		return fallback
	}
	return val
}
