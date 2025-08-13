package workers

import (
	"context"
	"log"
	"time"

	"github.com/bytedance/sonic"
	"github.com/chrismarsilva/cms.project.1million/internal/dtos"
	"github.com/chrismarsilva/cms.project.1million/internal/stores"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/wagslane/go-rabbitmq"
	//amqp "github.com/rabbitmq/amqp091-go"
)

var (
	Jobs = make(chan dtos.PersonRequestDto, 1000)
)

type PersonPublisherWorker struct {
	Config         *utils.Config
	RabbitMQClient *stores.RabbitMQ
	WorkerID       int
}

func NewPersonPublisherWorker(config *utils.Config, rabbitMQClient *stores.RabbitMQ, workerID int) *PersonPublisherWorker {
	return &PersonPublisherWorker{
		Config:         config,
		RabbitMQClient: rabbitMQClient,
		WorkerID:       workerID,
	}
}

func (w *PersonPublisherWorker) Start(jobs chan dtos.PersonRequestDto) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				//log.Printf("[SalvePaymentWorker %d] Job channel closed, stopping worker", w.WorkerNum)
				return
			}

			//log.Printf("[SalvePaymentWorker %d] Processing job: %v", w.WorkerNum, job)

			request := job
			payload, err := sonic.Marshal(request)
			if err != nil {
				log.Printf("[SalvePaymentWorker %d] Failed to marshal payment: %v", w.WorkerID, err)
				continue
			}

			ctx := context.Background()
			// ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			// defer cancel()

			err = w.RabbitMQClient.Publisher.PublishWithContext(
				ctx,
				payload,
				[]string{w.Config.RabbitMqQueue},
				rabbitmq.WithPublishOptionsContentType("application/json"),
				rabbitmq.WithPublishOptionsMandatory,
				rabbitmq.WithPublishOptionsPersistentDelivery,
				rabbitmq.WithPublishOptionsExchange(""),
			)
			if err != nil {
				log.Printf("[SalvePaymentWorker %d] Failed to enqueue payment: %v", w.WorkerID, err)
				continue
			}
		default:
			// No job available, sleep briefly to avoid spinning
			time.Sleep(10 * time.Millisecond)
		}
	}
}
