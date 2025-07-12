package health

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

const healthTimeout = 6 * time.Second

func HealthCheck(defaultURL string, fallbackURL string) {
	//log.Println("== Health Check - INI ==")

	var wg sync.WaitGroup

	wg.Add(1)
	go func(wwLocal *sync.WaitGroup) {
		defer wwLocal.Done()
		var start time.Time = time.Now()

		res, err := fetchHealth(defaultURL)
		if err != nil {
			log.Fatalf("== Health Check - default - error: %v\n", err)
		}

		log.Printf("== Health Check - default - time=%v, failing=%v, minResponseTime=%dms\n", time.Since(start), res.Failing, res.MinResponseTime)
	}(&wg)

	wg.Add(1)
	go func(wwLocal *sync.WaitGroup) {
		defer wwLocal.Done()
		var start time.Time = time.Now()

		res, err := fetchHealth(fallbackURL)
		if err != nil {
			log.Fatalf("== Health Check - fallback error: %v\n", err)
		}

		log.Printf("== Health Check - fallback - time=%v, failing=%v, minResponseTime=%dms\n", time.Since(start), res.Failing, res.MinResponseTime)
	}(&wg)

	wg.Wait()
	//log.Println("== Health Check - FIM ==")
}

func fetchHealth(url string) (*dtos.HealthResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), healthTimeout)
	defer cancel()

	data, err := http.DoGet(ctx, url+"/payments/service-health")
	if err != nil {
		return nil, fmt.Errorf("failed: %w", err)
	}

	var h dtos.HealthResponse
	if err := json.Unmarshal(data, &h); err != nil {
		return nil, fmt.Errorf("parsing health %w", err)
	}

	return &h, nil
}
