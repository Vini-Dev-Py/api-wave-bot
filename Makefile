# Makefile

# Variáveis
APP_NAME=api-wave-bot
DOCKER_COMPOSE=docker-compose
DOCKER_SERVICE=api

# Comandos

# Rodar o projeto com Docker
dev:
	$(DOCKER_COMPOSE) up --build

# Parar e remover containers
stop:
	$(DOCKER_COMPOSE) down

# Reiniciar o Docker (down e up)
restart:
	$(DOCKER_COMPOSE) down && $(DOCKER_COMPOSE) up --build

# Entrar no container 'api'
sh:
	$(DOCKER_COMPOSE) exec $(DOCKER_SERVICE) sh

# Subir apenas o banco de dados (em background)
db-up:
	$(DOCKER_COMPOSE) up -d db

# Entrar no container do banco de dados
db-sh:
	$(DOCKER_COMPOSE) exec db psql -U postgres -d $(APP_NAME)

# Rodar o Go mod tidy para limpar dependências
tidy:
	go mod tidy

# Instalar a dependência do pq (driver do PostgreSQL)
get-lib:
	go get github.com/lib/pq

# Rodar o projeto localmente (sem Docker)
run-local:
	go run main.go

# Rodar o projeto no Docker com hot reload (usando o air)
air:
	$(DOCKER_COMPOSE) exec $(DOCKER_SERVICE) air
