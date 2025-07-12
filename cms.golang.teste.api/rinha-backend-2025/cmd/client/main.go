package main

// go run ./cmd/client/main.go

import (
	"github.com/chrismarsilva/rinha-backend-2025/cmd/client/health"
	"github.com/chrismarsilva/rinha-backend-2025/cmd/client/payment"
)

const (
	defaultURL  = "http://localhost:8001"
	fallbackURL = "http://localhost:8002"
)

func main() {
	health.HealthCheck(defaultURL, fallbackURL)
	payment.Payment(defaultURL, fallbackURL)
	health.HealthCheck(defaultURL, fallbackURL)
	payment.PaymentSummary(defaultURL, fallbackURL)
}
