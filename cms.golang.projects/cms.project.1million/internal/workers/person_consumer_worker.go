package workers

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bytedance/sonic"
	"github.com/chrismarsilva/cms.project.1million/internal/dtos"
	"github.com/chrismarsilva/cms.project.1million/internal/models"
	"github.com/chrismarsilva/cms.project.1million/internal/stores"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/wagslane/go-rabbitmq"
)

var (
	EventConsumer = make(chan dtos.PersonRequestDto, 50000) // Buffer size can be adjusted based on expected load
)

type PersonConsumerWorker struct {
	Config         *utils.Config
	RabbitMQClient *stores.RabbitMQ
	RedisCache     *stores.RedisCache
	WorkerID       int
}

func NewPersonConsumerWorker(config *utils.Config, rabbitMQClient *stores.RabbitMQ, redisCache *stores.RedisCache, workerID int) *PersonConsumerWorker {
	return &PersonConsumerWorker{
		Config:         config,
		RabbitMQClient: rabbitMQClient,
		RedisCache:     redisCache,
		WorkerID:       workerID,
	}
}

func (w *PersonConsumerWorker) Start(eventConsumer chan dtos.PersonRequestDto) {
	// // Set prefetchCount to 50 to allow 50 messages before Acks are returned
	// err := w.RabbitMQClient.Channel.Qos(50, 0, false) // prefetch 1000
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

		eventConsumer <- request

		// model := models.NewPersonModel(request.Name)
		// //log.Printf("Worker #%d received a person: %s", w.WorkerID, model)

		// ctx := context.Background()
		// // ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		// // defer cancel()

		// err = w.PersonRepo.Add(ctx, *model)
		// if err != nil {
		// 	log.Printf("Worker #%d failed to add person to repository: %v", w.WorkerID, err)
		// 	return rabbitmq.NackDiscard
		// }

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

func (w *PersonConsumerWorker) Process(eventConsumer chan dtos.PersonRequestDto) {
	ctx := context.Background()

	pipe := w.RedisCache.Client.Pipeline()
	count := 0
	total := 0
	// // ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	// // defer cancel()

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case request, ok := <-eventConsumer:
			if !ok {
				if count > 0 {
					log.Printf("Worker #%d executing remaining %d commands in pipeline", w.WorkerID, count)
					ctx, span := utils.Tracer.Start(ctx, "PersonConsumerWorker.Process.pipeline.Exec")
					_, err := pipe.Exec(ctx) // executa em batch
					span.End()
					if err != nil {
						log.Printf("Worker #%d failed to execute remaining pipeline: %v", w.WorkerID, err)
					}
				}
				//log.Printf("[SalvePaymentWorker %d] Job channel closed, stopping worker", w.WorkerNum)
				return
			}

			// log.Printf("Worker #%d processing request: %+v", w.WorkerID, request)

			model := models.NewPersonModel(request.Name)
			// log.Printf("Worker #%d received a person: %s", w.WorkerID, model)

			payload, err := sonic.Marshal(model)
			if err != nil {
				log.Printf("Worker #%d Failed to marshal payment: %v", w.WorkerID, err)
				continue
			}

			//pipe.Set(ctx, "persons1:"+model.ID.String(), payload, 0)
			//pipe.HSet(ctx, "persons2:"+model.ID.String(), payload, 0)
			//pipe.HSet(ctx, "persons3", payload, 0)
			//pipe.LPush(ctx, "persons4:queue", payload).Err()
			//pipe.LPush(ctx, "persons4:"+model.ID.String(), payload).Err()
			err = pipe.HSet(ctx, "persons", model.ID.String(), payload).Err()
			if err != nil {
				log.Printf("Worker #%d failed to set person in Redis: %v", w.WorkerID, err)
				continue
			}
			count++
			total++

			// w.RedisCache.Client.HSet(ctx, "persons1", model.ID.String(), payload).Err()
			// w.RedisCache.Client.HSet(ctx, "person2:"+model.ID.String(), model.ID.String(), payload).Err()
			// err = w.PersonRepo.Add(ctx, *model)
			// if err != nil {
			// 	log.Printf("Worker #%d failed to add person to repository: %v", w.WorkerID, err)
			// 	return rabbitmq.NackDiscard
			// }

			if count >= w.Config.NumConsumerBatchSize {
				log.Printf("Worker #%d executing %d commands in pipeline", w.WorkerID, count)

				ctx, span := utils.Tracer.Start(ctx, "PersonConsumerWorker.Process.pipeline.Exec")
				_, err := pipe.Exec(ctx) // executa em batch
				span.End()

				if err != nil {
					log.Printf("Worker #%d failed to execute pipeline: %v", w.WorkerID, err)
					continue
				}

				// for _, c := range cmds {
				// 	fmt.Printf("%v;", c.(*redis.StatusCmd).Val())
				// }
				count = 0
				pipe = w.RedisCache.Client.Pipeline() // cria novo pipeline
			}

			//log.Printf("Worker #%d successfully added person to repository: %s\n", w.WorkerID, model.Name)

		case <-ticker.C:
			if count > 0 {
				log.Printf("Worker #%d executing by time %d commands in pipeline", w.WorkerID, count)

				ctx, span := utils.Tracer.Start(ctx, "PersonConsumerWorker.Process.pipeline.Exec")
				_, err := pipe.Exec(ctx) // executa em batch
				span.End()
				if err != nil {
					log.Printf("Worker #%d failed to execute pipeline: %v", w.WorkerID, err)
					continue
				}

				// for _, c := range cmds {
				// 	fmt.Printf("%v;", c.(*redis.StatusCmd).Val())
				// }
				count = 0
				pipe = w.RedisCache.Client.Pipeline() // cria novo pipeline
			}

		default:
			// No job available, sleep briefly to avoid spinning
			time.Sleep(10 * time.Millisecond)
			//log.Printf("Worker #%d is waiting for messages", w.WorkerID)
		}
	}

	//log.Printf("Worker #%d finished processing messages", w.WorkerID)
}
