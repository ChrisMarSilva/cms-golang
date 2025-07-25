package services

import (
	"github.com/chrismarsilva/rinha-backend-2025/internal/repositories"
)

type PaymentService struct {
	Repo *repositories.PaymentRepository
}

func NewPaymentService(repo *repositories.PaymentRepository) *PaymentService {
	return &PaymentService{
		Repo: repo,
	}
}
