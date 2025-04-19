# Dockerfile
FROM golang:1.24-alpine

RUN apk add --no-cache git curl

WORKDIR /app

COPY .env .env

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["air", "-c", ".air.toml"]
