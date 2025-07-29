package adapters

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/bytedance/sonic"
	"github.com/chrismarsilva/rinha-backend-2025/internal/dtos"
	"github.com/chrismarsilva/rinha-backend-2025/internal/repositories"
	"github.com/redis/go-redis/v9"
)

type ProcessPaymentWorker struct {
	RedisClient *redis.Client
	HealthCheck *HealthCheckService
	Repo        *repositories.PaymentRepository
	WorkerNum   int
}

func NewProcessPaymentWorker(redisClient *redis.Client, healthCheck *HealthCheckService, paymentRepo *repositories.PaymentRepository, num int) *ProcessPaymentWorker {
	return &ProcessPaymentWorker{
		RedisClient: redisClient,
		HealthCheck: healthCheck,
		Repo:        paymentRepo,
		WorkerNum:   num,
	}
}

func (w *ProcessPaymentWorker) Start(ctx context.Context) {
	httpClient := &http.Client{
		Timeout: 3 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:        1000,
			DisableCompression:  true,
			MaxIdleConnsPerHost: 100,
			MaxConnsPerHost:     10000,
			IdleConnTimeout:     120 * time.Second,
			DisableKeepAlives:   false,
			DialContext: (&net.Dialer{
				Timeout:   2 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
		},
	}

	processingQueue := fmt.Sprintf("payments:processing:%d", w.WorkerNum)

	for {

		//RPopLPush: Retorna e remove atomicamente o último elemento da lista
		result, err := w.RedisClient.RPopLPush(ctx, "payments:queue", processingQueue).Result()
		if err != nil {
			if err == redis.Nil {
				time.Sleep(100 * time.Millisecond) // 1 * time.Second
				continue
			}

			log.Printf("[ProcessPaymentWorker %d] Redis error: %v", w.WorkerNum, err)
			time.Sleep(1 * time.Second)
			continue
		}

		var payment dtos.PaymentDto
		err = sonic.Unmarshal([]byte(result), &payment)
		if err != nil {
			log.Printf("[ProcessPaymentWorker %d] Failed to unmarshal payment: %v", w.WorkerNum, err)
			//LRem: Remove as primeiras ocorrências de elementos iguais a elementda lista armazenada
			w.RedisClient.LRem(ctx, processingQueue, 1, result)
			continue
		}

		processor := w.HealthCheck.GetCurrent()

		if !w.processPayment(ctx, payment, processor, httpClient) {
			// LPush: Insere todos os valores especificados no início da lista
			w.RedisClient.LPush(ctx, "payments:queue", result)
			//continue
		}

		//LRem: Remove as primeiras ocorrências de elementos iguais a elementda lista armazenada
		w.RedisClient.LRem(ctx, processingQueue, 1, result)
	}
}

func (w *ProcessPaymentWorker) processPayment(ctx context.Context, payment dtos.PaymentDto, processor dtos.ProcessorStatusDto, httpClient *http.Client) bool {
	if w.Repo == nil {
		log.Printf("[ProcessPaymentWorker %d] CRITICAL: Repo is nil!", w.WorkerNum)
		return false
	}

	if processor.Service == "out" {
		log.Printf("[ProcessPaymentWorker %d] CRITICAL: Processor is down, skipping payment processing!", w.WorkerNum)
		return false
	}

	if processor.URL == "" {
		log.Printf("[ProcessPaymentWorker %d] CRITICAL: Processor URL is empty!", w.WorkerNum)
		return false
	}

	body := dtos.PaymentResponseDto{CorrelationID: payment.CorrelationId, Amount: payment.Amount, RequestedAt: payment.RequestedAt}
	payload, _ := sonic.Marshal(body)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, processor.URL, bytes.NewReader(payload))
	if err != nil {
		log.Printf("[ProcessPaymentWorker %d] Failed to create request: %v", w.WorkerNum, err)
		return false
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("[ProcessPaymentWorker %d] Processor %s HTTP error: %v", w.WorkerNum, processor.Service, err)
		return false
	}

	if resp.Body != nil {
		defer resp.Body.Close()
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("[ProcessPaymentWorker %d] Processor %s returned status: %v %s", w.WorkerNum, processor.Service, resp.StatusCode, string(body))
		return false
	}

	payment.Processor = processor.Service

	if err := w.Repo.StorePayment(ctx, payment); err != nil {
		log.Printf("[ProcessPaymentWorker %d] CRITICAL: Payment accepted by processor but failed to save in Redis: %v", w.WorkerNum, err)
	}

	return true
}
