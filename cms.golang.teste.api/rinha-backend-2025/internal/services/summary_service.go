package services

import (
	"context"
	"sync"
	"time"

	"github.com/chrismarsilva/rinha-backend-2025/internal/dtos"
	"github.com/chrismarsilva/rinha-backend-2025/internal/repositories"
)

type SummaryService struct {
	Repo *repositories.PaymentRepository
}

func NewSummaryService(repo *repositories.PaymentRepository) *SummaryService {
	return &SummaryService{
		Repo: repo,
	}
}

func (s *SummaryService) GetSummary(ctx context.Context, from, to *time.Time) (dtos.SummaryResponseDto, error) {
	payments, err := s.Repo.GetAllPayments(ctx)
	if err != nil {
		return dtos.SummaryResponseDto{}, err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var defaultCount, fallbackCount int
	var defaultAmount, fallbackAmount float64

	for _, p := range payments {
		payment := p
		wg.Add(1)

		go func() {
			defer wg.Done()

			if from != nil && !from.IsZero() && payment.RequestedAt.Before(*from) {
				return // continue
			}

			if to != nil && !to.IsZero() && payment.RequestedAt.After(*to) {
				return // continue
			}

			mu.Lock()
			defer mu.Unlock()

			switch payment.Processor {
			case "default":
				defaultCount++
				defaultAmount += payment.Amount
			case "fallback":
				fallbackCount++
				fallbackAmount += payment.Amount
			}
		}()
	}

	wg.Wait()

	data := dtos.SummaryResponseDto{
		Default:  dtos.SummaryItemResponseDto{TotalRequests: defaultCount, TotalAmount: defaultAmount},
		Fallback: dtos.SummaryItemResponseDto{TotalRequests: fallbackCount, TotalAmount: fallbackAmount},
	}

	return data, nil
}
