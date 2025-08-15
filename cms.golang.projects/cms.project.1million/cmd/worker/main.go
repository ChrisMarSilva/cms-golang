package main

import (
	"log"

	"github.com/chrismarsilva/cms.project.1million/internal/stores"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/chrismarsilva/cms.project.1million/internal/workers"
)

// go run ./cmd/api/main.go
// go run ./cmd/worker/main.go

// docker-compose down
// docker-compose up -d --build
// docker compose -f docker-compose.dev.yml up -d --build

// export K6_WEB_DASHBOARD=true
// export K6_WEB_DASHBOARD_PORT=5665
// export K6_WEB_DASHBOARD_PERIOD=2s
// export K6_WEB_DASHBOARD_OPEN=true
// export K6_WEB_DASHBOARD_EXPORT='report.html'

// k6 run ./loadtest.js

func main() {
	config := utils.NewConfig()

	redisCache := stores.NewRedisCache(config)
	defer redisCache.Close()

	rabbitMQClient := stores.NewRabbitMQ(config)
	defer rabbitMQClient.Close()

	//personRepo := repositories.NewPersonRepository(redisCache)

	cleanup := utils.InitOpenTelemetry("worker.1million")
	defer cleanup()

	//utils.Tracer = otel.Tracer("cms.worker.1million")
	//utils.Meter = otel.Meter("go-gin-service")

	// sigchan := make(chan os.Signal, 1)
	// signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	forever := make(chan bool)

	// numConsumers := runtime.NumCPU() * 2   // 2 consumidores por núcleo
	// numWorkers := runtime.NumCPU() * 8     // 8 workers por núcleo
	// pipelinesPerWorker := 4                // sub-pipelines por worker
	//log.Printf("Starting %d workers...", config.NumConsumerWorkers)

	for i := 0; i < config.NumConsumerWorkers; i++ {
		go func(idx int) {
			worker := workers.NewPersonConsumerWorker(config, rabbitMQClient, redisCache, idx)
			go worker.Start(workers.EventConsumer)
			go worker.Process(workers.EventConsumer)
		}(i)
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	//<-sigchan

	forever <- true
	log.Printf("interrupted, shutting down")
}
