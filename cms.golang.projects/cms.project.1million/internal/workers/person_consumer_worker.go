package workers

import (
	"context"
	"log/slog"
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
	logger   *slog.Logger
	config   *utils.Config
	mq       *stores.RabbitMQ
	db       *stores.Database
	rdb      *stores.RedisCache
	workerID int
}

func NewPersonConsumerWorker(logger *slog.Logger, config *utils.Config, mq *stores.RabbitMQ, db *stores.Database, rdb *stores.RedisCache, workerID int) *PersonConsumerWorker {
	return &PersonConsumerWorker{
		logger:   logger,
		config:   config,
		mq:       mq,
		db:       db,
		rdb:      rdb,
		workerID: workerID,
	}
}

func (w *PersonConsumerWorker) Start(eventConsumer chan dtos.PersonRequestDto) {

	err := w.mq.Consumer.Run(func(d rabbitmq.Delivery) rabbitmq.Action {
		var request dtos.PersonRequestDto
		err := sonic.Unmarshal([]byte(d.Body), &request)
		if err != nil {
			w.logger.Error("failed to unmarshal message", slog.Int("worker_id", w.workerID), slog.Any("error", err))
			return rabbitmq.NackDiscard
		}

		eventConsumer <- request
		return rabbitmq.Ack
	})

	if err != nil {
		w.logger.Error("failed to start consumer", slog.Int("worker_id", w.workerID), slog.Any("error", err))
		return
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigchan
		w.logger.Info("shutting down consumer", slog.Int("worker_id", w.workerID))
	}()
}

func (w *PersonConsumerWorker) Process(eventConsumer chan dtos.PersonRequestDto) {
	ctx := context.Background()

	pipe := w.rdb.Pipeline()
	count := 0
	total := 0

	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case request, ok := <-eventConsumer:
			if !ok {
				if count > 0 {
					w.logger.Info("executing remaining commands in pipeline", slog.Int("worker_id", w.workerID), slog.Int("count", count))
					_, err := pipe.Exec(ctx) // executa em batch
					if err != nil {
						w.logger.Error("failed to execute remaining pipeline", slog.Int("worker_id", w.workerID), slog.Any("error", err))
					}
				}
				return
			}

			model := models.NewPersonModel(request.Name)

			payload, err := sonic.Marshal(model)
			if err != nil {
				w.logger.Error("failed to marshal person model", slog.Int("worker_id", w.workerID), slog.Any("error", err))
				continue
			}

			err = pipe.HSet(ctx, "persons", model.ID.String(), payload).Err()
			if err != nil {
				w.logger.Error("failed to set person in Redis", slog.Int("worker_id", w.workerID), slog.Any("error", err))
				continue
			}
			count++
			total++

			if count >= w.config.NumConsumerBatchSize {
				w.logger.Info("executing commands in pipeline", slog.Int("worker_id", w.workerID), slog.Int("count", count))

				_, err := pipe.Exec(ctx) // executa em batch
				if err != nil {
					w.logger.Error("failed to execute pipeline", slog.Int("worker_id", w.workerID), slog.Any("error", err))
					continue
				}

				count = 0
				pipe = w.rdb.Pipeline() // cria novo pipeline
			}

		case <-ticker.C:
			if count > 0 {
				w.logger.Info("executing commands in pipeline", slog.Int("worker_id", w.workerID), slog.Int("count", count))

				_, err := pipe.Exec(ctx) // executa em batch
				if err != nil {
					w.logger.Error("failed to execute pipeline", slog.Int("worker_id", w.workerID), slog.Any("error", err))
					continue
				}

				count = 0
				pipe = w.rdb.Pipeline() // cria novo pipeline
			}

		default:
			// No job available, sleep briefly to avoid spinning
			time.Sleep(10 * time.Millisecond)
		}
	}
}
