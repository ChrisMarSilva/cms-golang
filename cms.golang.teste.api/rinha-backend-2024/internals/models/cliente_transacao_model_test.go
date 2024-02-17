package models_test

import (
	"testing"
	"time"

	"github.com/chrismarsilva/rinha-backend-2024/internals/models"
)

// go test
// go test -v
// go test -bench=.
// go test -bench=Add
// go test -run TestUClienteRepoGetAll -v
// go test -run=XXX -bench . -benchmem

func TestEntityTransactionEmpty(t *testing.T) {
	entity := models.ClienteTransacao{IdTransacao: 1, IdCliente: 2, Valor: 100, Tipo: "c", Descricao: "teste1", DtHrRegistro: time.Now()}
	entityEmpty := models.ClienteTransacao{}

	if entity == entityEmpty {
		t.Fatalf("entity null ")
	}
}

func BenchmarkTransactionEntity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = models.ClienteTransacao{IdTransacao: 1, IdCliente: 2, Valor: 100, Tipo: "c", Descricao: "teste1", DtHrRegistro: time.Now()}
	}
}
