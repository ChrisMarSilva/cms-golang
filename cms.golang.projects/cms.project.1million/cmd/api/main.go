package main

import (
	"log"

	"github.com/chrismarsilva/cms.project.1million/internal/handlers"
	"github.com/chrismarsilva/cms.project.1million/internal/repositories"
	"github.com/chrismarsilva/cms.project.1million/internal/services"
	"github.com/chrismarsilva/cms.project.1million/internal/stores"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/chrismarsilva/cms.project.1million/internal/workers"
)

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
// go get -u go.opentelemetry.io/otel/attribute
// go get -u go.opentelemetry.io/otel/exporters/otlp/otlptrace
// go get -u go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc
// go get -u go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc
// go get -u go.opentelemetry.io/otel/exporters/jaeger
// go get -u go.opentelemetry.io/otel/exporters/prometheus
// go get -u go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin
// go get -u github.com/prometheus/client_golang/prometheus
// go get -u github.com/prometheus/client_golang/prometheus/promhttp
// go get -u github.com/zsais/go-gin-prometheus
// go mod tidy

// go run ./cmd/api/main.go
// go run ./cmd/worker/main.go

// go get -u "github.com/cosmtrek/air@latest"
// air init
// air

// docker-compose down
// docker-compose up -d --build
// docker compose -f docker-compose.dev.yml up -d --build

// export K6_WEB_DASHBOARD=true
// export K6_WEB_DASHBOARD_PORT=5665
// export K6_WEB_DASHBOARD_PERIOD=2s
// export K6_WEB_DASHBOARD_OPEN=true
// export K6_WEB_DASHBOARD_EXPORT='report.html'

// k6 run ./loadtest.js

// REDIS    ✗ http_req_duration...........: avg=656.49ms min=1.53ms med=574.78ms max=5.25s    p(90)=1.26s   p(95)=1.56s
// RABBITMQ ✗ http_req_duration...........: avg=755.47ms min=0s      med=410.78ms max=19.55s   p(90)=1.47s p(95)=1.97s
// REDIS     { expected_response:true }....: avg=652.34ms min=1.53ms med=571.81ms max=5.25s    p(90)=1.25s   p(95)=1.55s
// RABBITMQ  { expected_response:true }....: avg=664.89ms min=814.1µs med=426.22ms max=12.08s   p(90)=1.45s p(95)=1.88s
// REDIS    http_req_failed................: 0.59%  ✓ 2142       ✗ 357158
// RABBITMQ http_req_failed................: 5.51%  ✓ 18608     ✗ 318956
// REDIS    http_req_waiting...............: avg=656.05ms min=1.52ms med=574.42ms max=5.25s    p(90)=1.26s   p(95)=1.56s
// RABBITMQ http_req_waiting...............: avg=754.81ms min=0s      med=410.43ms max=19.55s   p(90)=1.47s p(95)=1.97s
// REDIS    iteration_duration.............: avg=1.67s    min=1s     med=1.58s    max=6.48s    p(90)=2.29s   p(95)=2.59s
// RABBITMQ iteration_duration.............: avg=1.78s    min=1s      med=1.43s    max=20.57s   p(90)=2.51s p(95)=3.01s
// REDIS    iterations.....................: 359300 498.551426/s
// RABBITMQ iterations.....................: 337564 468.71404/s

func main() {
	config := utils.NewConfig()

	redisCache := stores.NewRedisCache(config)
	defer redisCache.Close()

	rabbitMQClient := stores.NewRabbitMQ(config)
	defer rabbitMQClient.Close()

	cleanup := utils.InitOpenTelemetry("api.1million")
	defer cleanup()

	//utils.Tracer = otel.Tracer("cms.api.1million")
	//utils.Meter = otel.Meter("go-gin-service")

	personRepo := repositories.NewPersonRepository(redisCache)
	personSvc := services.NewPersonService(personRepo, rabbitMQClient)
	personHandler := handlers.NewPersonHandler(personSvc)

	for i := 0; i < config.NumPublisherWorkers; i++ {
		go func(idx int) {
			worker := workers.NewPersonPublisherWorker(config, rabbitMQClient, idx)
			go worker.Start(workers.EventPublisher)
		}(i)
	}

	router := handlers.NewRouter(config, personHandler)
	log.Fatal(router.Listen())
}
