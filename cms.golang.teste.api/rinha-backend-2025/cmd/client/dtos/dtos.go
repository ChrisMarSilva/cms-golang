package dtos

import (
	"time"

	"github.com/google/uuid"
)

type HealthResponse struct {
	Failing         bool `json:"failing"`
	MinResponseTime int  `json:"minResponseTime"`
}

type PaymentRequest struct {
	CorrelationID uuid.UUID `json:"correlationId"`
	Amount        float64   `json:"amount"`
	RequestedAt   time.Time `json:"requestedAt"`
}

type PaymentResponse struct {
	Message string `json:"message"`
}

type SummaryResponse struct {
	TotalRequests     int64   `json:"totalRequests"`
	TotalAmount       float64 `json:"totalAmount"`
	TotalFee          float64 `json:"totalFee,omitempty"`
	FeePerTransaction float64 `json:"feePerTransaction,omitempty"`
}
