package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/chrismarsilva/rinha-backend-2025/cmd/client/dtos"
)

func DoGet(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Rinha-Token", "123")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("status %d: %s", resp.StatusCode, string(body))
	}

	return io.ReadAll(resp.Body)
}

func DoPost(ctx context.Context, url string, payload []byte) (*dtos.PaymentResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(payload))
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("X-Rinha-Token", "123")

	client := http.Client{} // Timeout: time.Second * 10
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request to %s failed: %w", url, err)
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}

	//log.Println("== Status Code:", resp.StatusCode, "-", resp.Status, "- Tempo:", time.Since(start), "==")
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("non-2xx status %d: %s", resp.StatusCode, string(body))
	}

	var response dtos.PaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("decoding response: %w", err)
	}
	//log.Println("== Response:", response.Message, "==")

	return &response, nil
}
