package workers

import (
	"context"
	"log/slog"
	"time"

	"github.com/bytedance/sonic"
	"github.com/chrismarsilva/cms.project.1million/internal/dtos"
	"github.com/chrismarsilva/cms.project.1million/internal/stores"
	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/wagslane/go-rabbitmq"
)

var (
	EventPublisher = make(chan dtos.PersonRequestDto, 1000)
)

type PersonPublisherWorker struct {
	logger   *slog.Logger
	config   *utils.Config
	mq       *stores.RabbitMQ
	workerID int
}

func NewPersonPublisherWorker(logger *slog.Logger, config *utils.Config, mq *stores.RabbitMQ, workerID int) *PersonPublisherWorker {
	return &PersonPublisherWorker{
		logger:   logger,
		config:   config,
		mq:       mq,
		workerID: workerID,
	}
}

func (w *PersonPublisherWorker) Start(eventPublisher chan dtos.PersonRequestDto) {
	ctx := context.Background()

	for {
		select {
		case request, ok := <-eventPublisher:
			if !ok {
				//w.logger.Info("[SalvePaymentWorker %d] Job channel closed, stopping worker", w.WorkerNum)
				return
			}

			//w.logger.Info("[SalvePaymentWorker %d] Processing job: %v", w.WorkerNum, job)

			//request := event
			payload, err := sonic.Marshal(request)
			if err != nil {
				w.logger.Error("Failed to marshal person request", slog.Any("error", err))
				continue
			}

			ctx, span := utils.Tracer.Start(ctx, "PersonPublisherWorker.Start")
			err = w.mq.Publisher.PublishWithContext(
				ctx,
				payload,
				[]string{w.config.RabbitMqQueue},
				rabbitmq.WithPublishOptionsContentType("application/json"),
				rabbitmq.WithPublishOptionsMandatory,
				rabbitmq.WithPublishOptionsPersistentDelivery,
				rabbitmq.WithPublishOptionsExchange(""),
			)
			span.End()
			if err != nil {
				w.logger.Error("Failed to publish person request", slog.Any("error", err))
				continue
			}

		default:
			// No job available, sleep briefly to avoid spinning
			time.Sleep(10 * time.Millisecond)
		}
	}
}
