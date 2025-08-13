package workers

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bytedance/sonic"
	"github.com/chrismarsilva/cms.project.1million/internal/dtos"
	"github.com/chrismarsilva/cms.project.1million/internal/models"
	"github.com/chrismarsilva/cms.project.1million/internal/repositories"
	"github.com/chrismarsilva/cms.project.1million/internal/stores"
	"github.com/wagslane/go-rabbitmq"
	//amqp "github.com/rabbitmq/amqp091-go"
)

type PersonConsumerWorker struct {
	RabbitMQClient *stores.RabbitMQ
	PersonRepo     *repositories.PersonRepository
	WorkerID       int
}

func NewPersonConsumerWorker(rabbitMQClient *stores.RabbitMQ, personRepo *repositories.PersonRepository, workerID int) *PersonConsumerWorker {
	return &PersonConsumerWorker{
		RabbitMQClient: rabbitMQClient,
		PersonRepo:     personRepo,
		WorkerID:       workerID,
	}
}

func (w *PersonConsumerWorker) Start() {

	// // Set prefetchCount to 50 to allow 50 messages before Acks are returned
	// err := w.RabbitMQClient.Channel.Qos(50, 0, false)
	// if err != nil {
	// 	log.Fatalf("Failed to set QoS: %v", err)
	// }

	// q, err := w.RabbitMQClient.Channel.QueueDeclare(w.Config.RabbitMqQueue, true, false, false, false, amqp.Table{})
	// if err != nil {
	// 	log.Fatalf("Failed to declare queue: %v", err)
	// }

	// msgs, err := w.RabbitMQClient.Channel.ConsumeWithContext(ctx, w.Config.RabbitMqQueue, "", true, false, false, false, amqp.Table{})
	// if err != nil {
	// 	log.Fatalf("Failed to consume messages: %v", err)
	// }

	//var forever chan struct{}

	// go func() {
	// 	for d := range msgs {
	//     go func(msg amqp091.Delivery) {
	//
	// 		}(d)
	err := w.RabbitMQClient.Consumer.Run(func(d rabbitmq.Delivery) rabbitmq.Action {
		//log.Printf("Worker #%d received a message: %s", w.WorkerID, string(d.Body))

		var request dtos.PersonRequestDto
		err := sonic.Unmarshal([]byte(d.Body), &request)
		if err != nil {
			log.Printf("Worker #%d failed to unmarshal message: %v", w.WorkerID, err)
			return rabbitmq.NackDiscard
		}

		model := models.NewPersonModel(request.Name)
		//log.Printf("Worker #%d received a person: %s", w.WorkerID, model)

		ctx := context.Background()
		// ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		// defer cancel()

		err = w.PersonRepo.Add(ctx, *model)
		if err != nil {
			log.Printf("Worker #%d failed to add person to repository: %v", w.WorkerID, err)
			return rabbitmq.NackDiscard
		}

		//log.Printf("Worker #%d successfully added person to repository: %s\n", w.WorkerID, model.Name)
		//msg.Ack(false)
		return rabbitmq.Ack
	})
	// 	}
	// }()
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}

	// Captura de sinal para shutdown gracioso
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigchan
		log.Printf("Worker #%d shutting down consumer", w.WorkerID)
	}()

	/*
		for {
			select {
			case message := <-messages:
				log.Printf("Message: %s\n", message.Body)
			case <-sigchan:
				log.Println("Interrupt detected!")
				os.Exit(0)
			}
		}
	*/

	// log.Printf("Worker #%d is waiting for messages", w.WorkerID)
	// <-forever
	// log.Printf("Worker #%d finished processing messages", w.WorkerID)
}
