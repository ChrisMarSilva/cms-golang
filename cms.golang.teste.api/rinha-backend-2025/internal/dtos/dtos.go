package dtos

import (
	"time"

	"github.com/google/uuid"
)

type ProcessorStatusDto struct {
	Service string
	URL     string
}

type PaymentDto struct {
	CorrelationId uuid.UUID `json:"correlationId"`
	Amount        float64   `json:"amount"`
	RequestedAt   time.Time `json:"requestedAt"`
	Processor     string    `json:"processor"` // "default" ou "fallback" ou "out"
}

type PaymentRequestDto struct {
	CorrelationId uuid.UUID `json:"correlationId"`
	Amount        float64   `json:"amount"`
}

type PaymentResponseDto struct {
	CorrelationID uuid.UUID `json:"correlationId"`
	Amount        float64   `json:"amount"`
	RequestedAt   time.Time `json:"requestedAt"`
}

type HealthCheckResponseDto struct {
	Failing         bool `json:"failing"`
	MinResponseTime int  `json:"minResponseTime"`
}

type SummaryResponseDto struct {
	Default  SummaryItemResponseDto `json:"default"`
	Fallback SummaryItemResponseDto `json:"fallback"`
}

type SummaryItemResponseDto struct {
	TotalRequests int     `json:"totalRequests"`
	TotalAmount   float64 `json:"totalAmount"`
}

// // Backend

// type PaymentRequestDto struct {
// 	CorrelationId     uuid.UUID `json:"correlationId"`
// 	Amount float64   `json:"amount"`
// }

// type PaymentResponseDto struct {
// 	CorrelationId          uuid.UUID `json:"correlationId"`
// 	Amount      float64   `json:"amount"`
// 	RequestedAt time.Time `json:"requestedAt"`
// }

// type SummaryResponseDto struct {
// 	Default  SummaryDetailResponseDto `json:"default"`
// 	Fallback SummaryDetailResponseDto `json:"fallback"`
// }

// type SummaryDetailResponseDto struct {
// 	TotalRequests int     `json:"totalRequests"`
// 	TotalAmount   float64 `json:"totalAmount"`
// }

// type HealthResponseDto struct {
// 	Default  HealthDetailResponseDto `json:"default"`
// 	Fallback HealthDetailResponseDto `json:"fallback"`
// }

// type HealthDetailResponseDto struct {
// 	Failing         bool `json:"failing"`
// 	MinResponseTime int  `json:"minResponseTime"`
// }

// // PaymentProcessor

// type PaymentProcessorHealthResponse struct {
// 	Failing         bool `json:"failing"`
// 	MinResponseTime int  `json:"minResponseTime"`
// }

// type PaymentProcessorPaymentRequest struct {
// 	CorrelationID uuid.UUID `json:"correlationId"`
// 	Amount        float64   `json:"amount"`
// 	RequestedAt   time.Time `json:"requestedAt"`
// }

// type PaymentProcessorPaymentResponse struct {
// 	Message string `json:"message"`
// }

// type PaymentProcessorSummaryResponse struct {
// 	TotalRequests     int     `json:"totalRequests"`
// 	TotalAmount       float64 `json:"totalAmount"`
// 	TotalFee          float64 `json:"totalFee,omitempty"`
// 	FeePerTransaction float64 `json:"feePerTransaction,omitempty"`
// }
