package main

import (
	"log"

	"github.com/chrismarsilva/rinha-backend-2024/internals/routes"
	"github.com/chrismarsilva/rinha-backend-2024/internals/utils"
)

// go mod init github.com/chrismarsilva/rinha-backend-2024
// go get -u github.com/gofiber/fiber/v3
// go get -u github.com/jmoiron/sqlx
// go get -u github.com/jackc/pgx/v5
// go get -u github.com/jackc/pgx/v5/pgxpool
// go get -u github.com/goccy/go-json
// go get -u github.com/joho/godotenv
// go install github.com/cosmtrek/air@latest
// go mod tidy

// air
// go run main.go
// go run ./cmd/api-server/main.go

// docker-compose down
// docker-compose up -d --build
// docker rm -f $(docker ps -a -q)
// docker run -it rinha-backend-2024-api01:latest

func main() {
	cfg := utils.NewConfig()
	app := routes.NewRoutes()          // server.Initialize()
	log.Fatal(app.Listen(cfg.UriPort)) //log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
