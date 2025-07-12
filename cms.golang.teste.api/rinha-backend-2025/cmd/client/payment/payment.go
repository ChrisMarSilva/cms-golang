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
	"github.com/google/uuid"
)

var paymentTimeout = 10 * time.Second // const paymentTimeout = 200 * time.Millisecond

func Payment(defaultURL string, fallbackURL string) {
	// var mLocal sync.Mutex
	// mLocal.Lock()
	// mLocal.Unlock()

	//log.Println("== Process Payment - INI ==")

	tot := 1_000 // 1 - 10 - 100 - 1_000 - 10_000 - 100_000 - 1_000_000
	var start time.Time = time.Now()
	var wg sync.WaitGroup
	var isDefaultActive bool = true // Reset default processor on error

	for i := 1; i <= tot; i++ {
		if i%100 == 0 {
			//isDefaultActive = !isDefaultActive
			//log.Println("== Process Payment - Nro(", i, ") - Request(", payment.CorrelationID, "-", payment.Amount, "-", payment.RequestedAt.Format(time.RFC3339Nano), ") ==")
			log.Println("== Process Payment - Nro(", i, ") - isDefaultActive(", isDefaultActive, ")  - paymentTimeout(", paymentTimeout, ") ==")
			//time.Sleep(time.Millisecond * 10)
			//paymentTimeout = paymentTimeout + (1 * time.Second)
		}

		wg.Add(1)
		go func(wwLocal *sync.WaitGroup, iIdx int, isDefaultActiveLocal bool) {
			defer wwLocal.Done()

			payment := dtos.PaymentRequest{
				CorrelationID: uuid.New(),
				Amount:        float64(iIdx),
				RequestedAt:   time.Now().UTC(), // 2025-01-01T01:01:06
			}
			_, err := processPayment(defaultURL, fallbackURL, payment, &isDefaultActiveLocal)
			if err != nil {
				isDefaultActiveLocal = true // Reset default processor on error
				//paymentTimeout = paymentTimeout + (1 * time.Second)
				log.Fatalf("== Process Payment - error: %v", err)
			}

			// log.Println("Response:", resp.Message)
		}(&wg, i, isDefaultActive)
	}

	wg.Wait()
	log.Println("== Process Payment(", tot, ") - Tempo(", time.Since(start), ") - isDefaultActive(", isDefaultActive, ")  - paymentTimeout(", paymentTimeout, ") ==")
}

func processPayment(defaultURL string, fallbackURL string, request dtos.PaymentRequest, isDefaultActive *bool) (*dtos.PaymentResponse, error) {
	payload, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("json marshal: %w", err)
	}

	if *isDefaultActive {
		// 1st try default

		ctx1, cancel := context.WithTimeout(context.Background(), paymentTimeout)
		defer cancel()

		res, err := http.DoPost(ctx1, defaultURL+"/payments", payload)
		if err == nil {
			//log.Println("== Process Payment - Processed by default processor")
			return res, nil
		}

		log.Printf("== Process Payment - Default failed: %v; trying fallback...\n", err)
	}

	//paymentTimeout = paymentTimeout + (1 * time.Second)
	*isDefaultActive = false // Disable default processor on error

	// 2nd try fallback

	ctx2, cancel2 := context.WithTimeout(context.Background(), paymentTimeout)
	defer cancel2()

	res2, err2 := http.DoPost(ctx2, fallbackURL+"/payments", payload)
	if err2 != nil {
		return nil, fmt.Errorf("fallback failed: %w", err2)
	}

	//log.Println("== Process Payment - Processed by fallback processor")
	return res2, nil
}
