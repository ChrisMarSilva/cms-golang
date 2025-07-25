package repositories

import (
	"context"
	"fmt"
	"sync"

	"github.com/bytedance/sonic"
	"github.com/chrismarsilva/rinha-backend-2025/internal/dtos"
	"github.com/redis/go-redis/v9"
)

type PaymentRepository struct {
	RedisClient *redis.Client
}

func NewPaymentRepository(redisClient *redis.Client) *PaymentRepository {
	return &PaymentRepository{
		RedisClient: redisClient,
	}
}

func (r *PaymentRepository) StorePayment(ctx context.Context, payment dtos.PaymentDto) error {
	paymentJSON, err := sonic.Marshal(payment)
	if err != nil {
		return fmt.Errorf("failed to marshal payment data: %w", err)
	}

	//HSet: Define os campos especificados para seus respectivos valores no hash armazenado em key.
	return r.RedisClient.HSet(ctx, "payments", payment.CorrelationId.String(), paymentJSON).Err()
}

func (r *PaymentRepository) GetAllPayments(ctx context.Context) ([]*dtos.PaymentDto, error) {
	// HGetAll: Retorna todos os campos e valores do hash armazenado em key.
	paymentsData, err := r.RedisClient.HGetAll(ctx, "payments").Result()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve payments: %w", err)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	payments := make([]*dtos.PaymentDto, 0, len(paymentsData))

	for _, paymentDataJSON := range paymentsData {
		wg.Add(1)
		paymentJSON := paymentDataJSON

		go func() {
			defer wg.Done()

			var payment dtos.PaymentDto
			err := sonic.Unmarshal([]byte(paymentJSON), &payment)
			if err != nil {
				return // continue
			}

			mu.Lock()
			payments = append(payments, &payment)
			mu.Unlock()
		}()
	}

	wg.Wait()

	return payments, nil
}
