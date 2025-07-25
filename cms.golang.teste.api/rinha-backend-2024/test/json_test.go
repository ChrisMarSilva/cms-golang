package main

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"

	"github.com/bytedance/sonic"
	"github.com/google/uuid"
)

// go test -bench=.
// go test -bench=. -count 5 -run=^#
// go test -run=XXX -bench . -benchmem

type PaymentDto struct {
	ID          uuid.UUID `json:"correlationId"`
	Amount      float64   `json:"amount"`
	RequestedAt time.Time `json:"requestedAt"`
}

//-------------------------------------------------------------------------------

func BenchmarkJsonMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := map[string]interface{}{"correlationId": uuid.New().String(), "amount": 10.0, "requestedAt": time.Now().UTC()}
		json.Marshal(data)
	}
}

func BenchmarkJsonDto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := PaymentDto{ID: uuid.New(), Amount: 10.0, RequestedAt: time.Now().UTC()}
		json.Marshal(data)
	}
}

//-------------------------------------------------------------------------------

func BenchmarkSonicMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := map[string]interface{}{"correlationId": uuid.New().String(), "amount": 10.0, "requestedAt": time.Now().UTC()}
		sonic.Marshal(data)
	}
}

func BenchmarkSonicDto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := PaymentDto{ID: uuid.New(), Amount: 10.0, RequestedAt: time.Now().UTC()}
		sonic.Marshal(data)
	}
}

//-------------------------------------------------------------------------------

func BenchmarkSonicBufferMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := map[string]interface{}{"correlationId": uuid.New().String(), "amount": 10.0, "requestedAt": time.Now().UTC()}
		sonic.ConfigDefault.NewEncoder(bytes.NewBuffer(nil)).Encode(data)
	}
}

func BenchmarkSonicBufferDto(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data := PaymentDto{ID: uuid.New(), Amount: 10.0, RequestedAt: time.Now().UTC()}
		sonic.ConfigDefault.NewEncoder(bytes.NewBuffer(nil)).Encode(data)
	}
}

//-------------------------------------------------------------------------------
