package main

import (
	"log/slog"

	"github.com/chrismarsilva/cms.project.1million/internal/stores"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/chrismarsilva/cms.project.1million/internal/workers"
)

func main() {
	config := utils.NewConfig()

	redisCache := stores.NewRedisCache(config)
	defer redisCache.Close()

	rabbitMQClient := stores.NewRabbitMQ(config)
	defer rabbitMQClient.Close()

	cleanup := utils.InitOpenTelemetry("worker.1million")
	defer cleanup()

	//utils.Tracer = otel.Tracer("cms.worker.1million")
	//utils.Meter = otel.Meter("go-gin-service")

	forever := make(chan bool)

	for i := 0; i < config.NumConsumerWorkers; i++ {
		go func(workerID int) {
			worker := workers.NewPersonConsumerWorker(config, rabbitMQClient, redisCache, workerID)
			go worker.Start(workers.EventConsumer)
			go worker.Process(workers.EventConsumer)
		}(i)
	}
	slog.Info("Waiting for messages. To exit press CTRL+C")

	forever <- true
	slog.Info("Shutting down gracefully...")
}
