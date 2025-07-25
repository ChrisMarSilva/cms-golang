package adapters

import (
	"context"
	"log"
	"time"

	"github.com/bytedance/sonic"
	"github.com/chrismarsilva/rinha-backend-2025/internal/dtos"
	"github.com/redis/go-redis/v9"
)

var (
	Jobs = make(chan dtos.PaymentRequestDto, 1000)
)

type SalvePaymentWorker struct {
	RedisClient *redis.Client
	WorkerNum   int
}

func NewSalvePaymentWorker(redisClient *redis.Client, num int) *SalvePaymentWorker {
	return &SalvePaymentWorker{
		RedisClient: redisClient,
		WorkerNum:   num,
	}
}

func (w *SalvePaymentWorker) Start(ctx context.Context, jobs chan dtos.PaymentRequestDto) {
	//log.Printf("[SalvePaymentWorker %d] Started", w.WorkerNum)

	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				//log.Printf("[SalvePaymentWorker %d] Job channel closed, stopping worker", w.WorkerNum)
				return
			}

			//log.Printf("[SalvePaymentWorker %d] Processing job: %v", w.WorkerNum, job)

			payment := dtos.PaymentDto{CorrelationId: job.CorrelationId, Amount: job.Amount, RequestedAt: time.Now().UTC()}

			data, err := sonic.Marshal(payment)
			if err != nil {
				log.Printf("[SalvePaymentWorker %d] Failed to marshal payment: %v", w.WorkerNum, err)
				continue
			}

			err = w.RedisClient.LPush(ctx, "payments:queue", data).Err()
			if err != nil {
				log.Printf("[SalvePaymentWorker %d] Failed to enqueue payment: %v", w.WorkerNum, err)
				continue
			}
		default:
			// No job available, sleep briefly to avoid spinning
			time.Sleep(10 * time.Millisecond)
		}
	}
}
