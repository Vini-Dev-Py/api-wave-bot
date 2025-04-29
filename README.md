# 🌊 api-wave-bot

API desenvolvida em Go puro para gerenciamento de **fluxos de automação no WhatsApp**, inspirada em ferramentas como **n8n** e **BotConversa**. Essa aplicação é o backend do projeto **Wave Bot**, integrado ao frontend `app-wave-bot`, baseado em React Flow.

O objetivo é permitir que empresas criem **workflows visuais** para envio de mensagens, coleta de respostas e ações automatizadas por meio do WhatsApp.

---

## ✨ Funcionalidades

- ✅ Cadastro e autenticação de usuários (JWT)
- ✅ Suporte a multiempresa (tipos de usuários: Super, Admin Sistema, Admin Empresa, Usuário)
- ✅ Criação, armazenamento e execução de fluxos de mensagens (nodes e edges estilo React Flow)
- ✅ Execução de campanhas e automações por webhook
- ✅ Integração com banco de dados PostgreSQL
- ✅ Hot reload em desenvolvimento com [air](https://github.com/cosmtrek/air)
- ✅ Totalmente containerizado com Docker e Docker Compose

---

---

## 🚀 Como rodar localmente

### 🔧 Pré-requisitos

- [Go 1.22+](https://golang.org/dl/)
- [Docker + Docker Compose](https://www.docker.com/)
- [DBeaver](https://dbeaver.io/) (opcional, para explorar o banco)
- [Air (hot reload)](https://github.com/cosmtrek/air)

---

### 📦 Clone o projeto

```bash
git clone https://github.com/vinibatista/api-wave-bot.git
cd api-wave-bot
```

---

### 🐳 Inicie com Docker

```bash
docker-compose up
```

---

### 🌀 Rode com hot reload (Air)

```bash
air
```

A API será executada por padrão em:  
**`http://localhost:3333`**

---

## 📂 Conexão com DBeaver (opcional)

- **Host**: `localhost`
- **Porta**: `5432`
- **Usuário**: `postgres`
- **Senha**: `postgres`
- **Banco de dados**: `wavebot`

---

## 🔐 Autenticação

A API utiliza JWT. Após login, envie o token nas requisições autenticadas:

```
Authorization: Bearer <seu-token>
```

---

## 🌍 Variáveis de Ambiente

Crie um arquivo `.env` na raiz com:

```
PORT=3333
DATABASE_URL=postgres://postgres:postgres@localhost:5432/wavebot?sslmode=disable
JWT_SECRET=sua_chave_super_secreta
```

---

## 📘 Documentação da API

A documentação Swagger será adicionada em breve.

---

## 🧪 Testes

> Em implementação.

---

## 🛠️ To-Do

- [ ] Executar workflows por webhooks
- [ ] Criar painel de administração multiempresa
- [ ] Suporte a agendamento e delays nos fluxos
- [ ] Logs e histórico de execuções
- [ ] Integração com outros canais além do WhatsApp

---

## 🤝 Contribuição

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests.

---

Feito com 💙 por [Vinícius Guilherme Batista](https://github.com/Vini-Dev-Py)
