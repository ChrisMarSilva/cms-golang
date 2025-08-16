package main

import (
	"github.com/chrismarsilva/cms.project.1million/internal/stores"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/chrismarsilva/cms.project.1million/internal/workers"
)

func main() {
	cleanup := utils.InitOpenTelemetry("worker.1million")
	defer cleanup()

	//ctx := context.Background()
	logger := utils.Logger // utils.CreateLogger(ctx)

	config := utils.NewConfig(logger)

	db := stores.NewDatabase(logger, config)
	defer db.Close()

	rdb := stores.NewRedisCache(logger, config)
	defer rdb.Close()

	mq := stores.NewRabbitMQ(logger, config)
	defer mq.Close()

	//utils.Tracer = otel.Tracer("cms.worker.1million")
	//utils.Meter = otel.Meter("go-gin-service")

	forever := make(chan bool)

	for i := 0; i < config.NumConsumerWorkers; i++ {
		go func(workerID int) {
			worker := workers.NewPersonConsumerWorker(logger, config, mq, db, rdb, workerID)
			go worker.Start(workers.EventConsumer)
			go worker.Process(workers.EventConsumer)
		}(i)
	}
	logger.Info("Waiting for messages. To exit press CTRL+C")

	forever <- true
	logger.Info("Shutting down gracefully...")
}
