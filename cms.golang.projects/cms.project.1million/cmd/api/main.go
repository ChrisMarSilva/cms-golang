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

	repo := repositories.NewPersonRepository(redisCache)
	svc := services.NewPersonService(repo)
	handler := handlers.NewPersonHandler(svc)

	for i := 0; i < config.NumPublisherWorkers; i++ {
		go func(workerID int) {
			worker := workers.NewPersonPublisherWorker(config, rabbitMQClient, workerID)
			go worker.Start(workers.EventPublisher)
		}(i)
	}

	router := handlers.NewRouter(config, handler)
	log.Fatal(router.Listen())
}
