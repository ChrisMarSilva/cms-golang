package main

import (
	"github.com/chrismarsilva/rinha-backend-2024/internals/server"
)

// go mod init github.com/chrismarsilva/rinha-backend-2024
// go get -u github.com/gofiber/fiber/v3
// go get -u github.com/jmoiron/sqlx
// go get -u github.com/jackc/pgx/v5
// go get -u github.com/jackc/pgx/v5/pgxpool
// go get -u github.com/goccy/go-json
// go get -u github.com/joho/godotenv
// go get -u github.com/bytedance/sonic
// go mod tidy

// go install github.com/cosmtrek/air@latest
// air init
// air

// go run main.go
// go run ./cmd/api-server/main.go

// docker-compose down
// docker-compose up -d --build

// docker rm -f $(docker ps -a -q)
// docker run -it rinha-backend-2024-api01:latest

func main() {
	app := server.NewServer()
	app.Initialize()
}
