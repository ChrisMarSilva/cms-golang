package main

import (
	"log"

	"github.com/chrismarsilva/rinha-backend-2025/internals/databases"
	"github.com/chrismarsilva/rinha-backend-2025/internals/handlers"
	"github.com/chrismarsilva/rinha-backend-2025/internals/repositories"
	"github.com/chrismarsilva/rinha-backend-2025/internals/usecases"
	"github.com/chrismarsilva/rinha-backend-2025/internals/utils"
)

// go mod init github.com/chrismarsilva/rinha-backend-2025
// go get -u github.com/gin-gonic/gin
// go get -u github.com/gin-contrib/gzip
// go get -u github.com/google/uuid
// go get -u github.com/joho/godotenv
// go get -u github.com/jackc/pgx/v5
// go get -u github.com/jackc/pgx/v5/pgxpool
// go mod tidy

// go run ./cmd/api/main.go

func main() {
	log.Println("Starting Rinha Backend 2025...")

	cfg := utils.NewConfig()

	db := databases.Connect(cfg)
	defer databases.Close()

	// redis := infra.NewRedis(cfg)

	repos := repositories.New(db)
	useCases := usecases.New(repos)
	handlers := handlers.New(useCases, cfg)

	if err := handlers.Listen(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
