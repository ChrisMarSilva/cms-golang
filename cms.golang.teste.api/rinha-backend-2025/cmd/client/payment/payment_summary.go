package payment

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/chrismarsilva/rinha-backend-2025/cmd/client/dtos"
	"github.com/chrismarsilva/rinha-backend-2025/cmd/client/http"
)

const paymentSummaryTimeout = 200 * time.Millisecond

func PaymentSummary(defaultURL string, fallbackURL string) {
	//log.Println("== Payments Summary - INI ==")

	var wg sync.WaitGroup

	wg.Add(1)
	go func(wwLocal *sync.WaitGroup) {
		defer wwLocal.Done()
		var start time.Time = time.Now()

		res, err := fetchSummary(defaultURL)
		if err != nil {
			log.Fatalf("== Payments Summary - default - Fetch summary failed: %v", err)
		}

		log.Printf("== Payments Summary - default - time=%v totalRequests=%d totalAmount=%.2f totalFee=%.2f feePerTransaction=%.2f", time.Since(start), res.TotalRequests, res.TotalAmount, res.TotalFee, res.FeePerTransaction)
	}(&wg)

	wg.Add(1)
	go func(wwLocal *sync.WaitGroup) {
		defer wwLocal.Done()
		var start time.Time = time.Now()
		res, err := fetchSummary(fallbackURL)
		if err != nil {
			log.Fatalf("== Payments Summary - fallback - Fetch summary failed: %v", err)
		}
		log.Printf("== Payments Summary - fallback - time=%v totalRequests=%d totalAmount=%.2f totalFee=%.2f feePerTransaction=%.2f", time.Since(start), res.TotalRequests, res.TotalAmount, res.TotalFee, res.FeePerTransaction)
	}(&wg)

	wg.Wait()
	//log.Println("== Payments Summary - FIM ==")
}

func fetchSummary(url string) (*dtos.SummaryResponse, error) {

	//?from=0&size=0
	// if from != "" {
	// 	url = url + "/from:" + from
	// }
	// if to != "" {
	// 	url = url + "/to:" + to
	// }

	ctx, cancel := context.WithTimeout(context.Background(), paymentSummaryTimeout)
	defer cancel()

	data, err := http.DoGet(ctx, url+"/admin/payments-summary")
	if err != nil {
		return nil, fmt.Errorf("failed: %w", err)
	}

	var s dtos.SummaryResponse
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, fmt.Errorf("parsing summary: %w", err)
	}

	return &s, nil
}
