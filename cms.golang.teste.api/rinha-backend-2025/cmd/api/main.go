package main

import (
	"context"
	"log"

	"github.com/chrismarsilva/rinha-backend-2025/internal/adapters"
	"github.com/chrismarsilva/rinha-backend-2025/internal/handlers"
	"github.com/chrismarsilva/rinha-backend-2025/internal/repositories"
	"github.com/chrismarsilva/rinha-backend-2025/internal/services"
	"github.com/chrismarsilva/rinha-backend-2025/internal/stores"
	"github.com/chrismarsilva/rinha-backend-2025/internal/utils"
)

// go mod init github.com/chrismarsilva/rinha-backend-2025
// go get -u github.com/gin-gonic/gin
// go get -u github.com/gin-contrib/gzip
// go get -u github.com/google/uuid
// go get -u github.com/joho/godotenv
// go get -u github.com/redis/go-redis/v9
// go get -u github.com/jackc/pgx/v5/pgxpool
// go get -u github.com/stretchr/testify
// go mod tidy

// Worker.Start                - RPopLPush("payments:queue" , "payments:processing:%d ) // Retorna e remove o último
// Worker.Start.Ok.StorePayment - HSet(     "payments"                                 ) // Add hash
// Worker.Start.Erro            - LPush(    "payments:queue"                           ) // Insere no início da lista
// Worker.Start                 - LRem(     "payments:processing:%d"                   ) // Remove as primeiras elemento
// PaymentHandler               - LPush(    "payments:queue"                           ) // Insere no início da lista
// SummaryHandler               - HGetAll(  "payments"                                 ) // Retorna todos hash

// go run ./cmd/api/main.go

func main() {
	config := utils.NewConfig()
	redisClient := stores.NewRedis(config)

	paymentRepo := repositories.NewPaymentRepository(redisClient)

	paymentSvc := services.NewPaymentService(paymentRepo)
	summarySvc := services.NewSummaryService(paymentRepo)

	handlers := handlers.New(config, paymentSvc, summarySvc)

	healthCheck := adapters.NewHealthCheckService(config, redisClient)
	healthCheck.Start()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := 0; i < config.NumWorkers; i++ {
		go func(idx int) {
			worker := adapters.NewSalvePaymentWorker(redisClient, i)
			go worker.Start(ctx, adapters.Jobs)
		}(i)

		go func(idx int) {
			worker := adapters.NewProcessPaymentWorker(redisClient, healthCheck, paymentRepo, i)
			go worker.Start(ctx)
		}(i)
	}

	if err := handlers.Listen(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
