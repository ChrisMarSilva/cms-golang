package infra

// import (
// 	"bytes"
// 	"context"
// 	"errors"
// 	"fmt"
// 	"net"
// 	"net/http"
// 	"time"

// 	"github.com/sony/gobreaker"
// 	"gopkg.in/redis.v3"
// )

// type ProcessorClient struct {
// 	name   string
// 	url    string
// 	client *http.Client
// 	cb     *gobreaker.CircuitBreaker
// 	redis  *redis.Client
// }

// func NewProcessorClient(name, url string, rdb *redis.Client) *ProcessorClient {
// 	cb := gobreaker.NewCircuitBreaker(gobreaker.Settings{
// 		Name:        name,
// 		Timeout:     10 * time.Second,
// 		MaxRequests: 3,
// 		Interval:    30 * time.Second,
// 	})
// 	return &ProcessorClient{
// 		name: name, url: url,
// 		client: &http.Client{
// 			Transport: &http.Transport{
// 				MaxIdleConns:    10,
// 				IdleConnTimeout: 30 * time.Second,
// 				DialContext:     (&net.Dialer{Timeout: 2 * time.Second}).DialContext,
// 			},
// 			Timeout: 3 * time.Second,
// 		},
// 		cb:    cb,
// 		redis: rdb,
// 	}
// }

// func (p *ProcessorClient) Send(ctx context.Context, corrID string, cents int64) error {
// 	if healthy := IsHealthy(ctx, p.redis, p.name); healthy {
// 		// skip if healthy cache present
// 	} else {
// 		_, err := p.cb.Execute(func() (interface{}, error) {
// 			reqBody := []byte(`{"correlationId":"` + corrID + `", "amount":` +
// 				fmt.Sprintf("%d", cents/100.0) + `}`)
// 			req, _ := http.NewRequestWithContext(ctx, "POST", p.url+"/payments",
// 				bytes.NewReader(reqBody))
// 			req.Header.Set("Content-Type", "application/json")
// 			resp, err := p.client.Do(req)
// 			if err != nil {
// 				return nil, err
// 			}
// 			defer resp.Body.Close()
// 			if resp.StatusCode >= 500 {
// 				return nil, errors.New("server error")
// 			}
// 			return nil, nil
// 		})
// 		if err == nil {
// 			SetHealthy(ctx, p.redis, p.name, true)
// 		} else {
// 			SetHealthy(ctx, p.redis, p.name, false)
// 		}
// 	}
// 	return nil
// }
