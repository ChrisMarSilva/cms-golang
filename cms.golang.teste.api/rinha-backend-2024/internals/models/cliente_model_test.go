package models_test

import (
	"testing"

	"github.com/chrismarsilva/rinha-backend-2024/internals/models"
)

// go test
// go test -v
// go test -bench=.
// go test -bench=Add
// go test -run TestUClienteRepoGetAll -v
// go test -run=XXX -bench . -benchmem

func TestEntityClientEmpty(t *testing.T) {
	entity := models.Cliente{Limite: 1, Saldo: 1000}
	entityEmpty := models.Cliente{}

	if entity == entityEmpty {
		t.Fatalf("entity null ")
	}
}

func BenchmarkClientEntity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = models.Cliente{Limite: 1, Saldo: 1000}
	}
}
