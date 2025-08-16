package main

// go mod init github.com/chrismarsilva/cms.project.1million
// go get -u github.com/gin-gonic/gin
// go get -u github.com/gin-contrib/gzip
// go get -u github.com/google/uuid
// go get -u github.com/joho/godotenv
// go get -u github.com/redis/go-redis/v9
// go get -u github.com/bytedance/sonic
// go get -u github.com/rabbitmq/amqp091-go
// go get -u github.com/wagslane/go-rabbitmq
// go get -u go.opentelemetry.io/otel
// go get -u go.opentelemetry.io/otel/trace
// go get -u go.opentelemetry.io/otel/metric
// go get -u go.opentelemetry.io/otel/sdk/resource
// go get -u go.opentelemetry.io/otel/sdk/trace
// go get -u go.opentelemetry.io/otel/sdk/metric
// go get -u go.opentelemetry.io/otel/sdk/log
// go get -u go.opentelemetry.io/otel/attribute
// go get -u go.opentelemetry.io/otel/exporters/otlp/otlptrace
// go get -u go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc
// go get -u go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc
// go get -u go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp
// go get -u go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc
// go get -u go.opentelemetry.io/otel/exporters/jaeger
// go get -u go.opentelemetry.io/otel/exporters/prometheus
// go get -u go.opentelemetry.io/otel/exporters/stdout/stdoutlog
// go get -u go.opentelemetry.io/otel/log/global
// go get -u go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin
// go get -u go.opentelemetry.io/contrib/bridges/otelslog
// go get -u github.com/prometheus/client_golang/prometheus
// go get -u github.com/prometheus/client_golang/prometheus/promhttp
// go get -u github.com/zsais/go-gin-prometheus
// go get -u github.com/jackc/pgx/v5
// go get -u github.com/jackc/pgx/v5/pgxpool
// go mod tidy

// go get -u "github.com/cosmtrek/air@latest"
// air init
// air

import (
	"log/slog"

	"github.com/chrismarsilva/cms.project.1million/internal/handlers"
	"github.com/chrismarsilva/cms.project.1million/internal/repositories"
	"github.com/chrismarsilva/cms.project.1million/internal/services"
	"github.com/chrismarsilva/cms.project.1million/internal/stores"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/chrismarsilva/cms.project.1million/internal/workers"
)

func main() {
	cleanup := utils.InitOpenTelemetry("api.1million")
	defer cleanup()

	logger := utils.NewLogger() // utils.Logger
	logger.Info("App iniciado")
	logger.Warn("Algo pode dar errado", "context", "exemplo")
	logger.Error("Erro crítico ocorreu", "error", "detalhe do erro")

	config := utils.NewConfig(logger)

	db := stores.NewDatabase(logger, config)
	defer db.Close()

	redisCache := stores.NewRedisCache(logger, config)
	defer redisCache.Close()

	rabbitMQClient := stores.NewRabbitMQ(logger, config)
	defer rabbitMQClient.Close()

	repo := repositories.NewPersonRepository(logger, db, redisCache)
	svc := services.NewPersonService(logger, repo)
	handler := handlers.NewPersonHandler(logger, svc)

	for i := 0; i < config.NumPublisherWorkers; i++ {
		go func(workerID int) {
			worker := workers.NewPersonPublisherWorker(logger, config, rabbitMQClient, workerID)
			go worker.Start(workers.EventPublisher)
		}(i)
	}

	router := handlers.NewRouter(logger, config, handler)
	// handler.RegisterRoutes(router)

	if err := router.Listen(); err != nil {
		logger.Error("Failed to start server", slog.Any("error", err))
	}
}
