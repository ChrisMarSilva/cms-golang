package main

import (
	"log"

	"github.com/chrismarsilva/cms.project.1million/internal/repositories"
	"github.com/chrismarsilva/cms.project.1million/internal/stores"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/chrismarsilva/cms.project.1million/internal/workers"
)

func main() {
	config := utils.NewConfig()
	redisClient := stores.NewRedis(config)
	rabbitMQClient := stores.NewRabbitMQConnection(config)
	defer rabbitMQClient.CloseConnection()

	personRepo := repositories.NewPersonRepository(redisClient)

	// sigchan := make(chan os.Signal, 1)
	// signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	forever := make(chan bool)

	log.Printf("Starting %d workers...", config.NumWorkers)

	for i := 0; i < config.NumWorkers; i++ {
		go func(idx int) {
			worker := workers.NewPersonConsumerWorker(rabbitMQClient, personRepo, idx)
			go worker.Start()
		}(i)
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	//<-sigchan

	forever <- true
	log.Printf("interrupted, shutting down")
}
