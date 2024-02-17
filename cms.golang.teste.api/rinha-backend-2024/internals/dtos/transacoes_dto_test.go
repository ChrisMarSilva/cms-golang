package dtos_test

import (
	"testing"

	"github.com/chrismarsilva/rinha-backend-2024/internals/dtos"
)

// go test
// go test -v
// go test -bench=.
// go test -bench=Add
// go test -run TestTransacaoRequestDto -v
// go test -run=XXX -bench . -benchmem

func TestEntityEmpty(t *testing.T) {
	entity := dtos.TransacaoRequestDto{Valor: 1, Tipo: "d", Descricao: "teste"}
	entityEmpty := dtos.TransacaoRequestDto{}

	if entity == entityEmpty {
		t.Fatalf("entity null ")
	}
}

func TestEntityValidate(t *testing.T) {
	entity := dtos.TransacaoRequestDto{Valor: 1, Tipo: "d", Descricao: "teste"}
	if err := entity.Valido(); err != nil {
		t.Fatalf("error %t", err)
	}
}

func TestEntitiesValidate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   dtos.TransacaoRequestDto
		withErr bool
	}{
		{"OK", dtos.TransacaoRequestDto{Valor: 1, Tipo: "d", Descricao: "teste"}, false},
		{"OK", dtos.TransacaoRequestDto{Valor: 1, Tipo: "d", Descricao: "teste"}, false},
		{"OK", dtos.TransacaoRequestDto{Valor: 9223372036854775807, Tipo: "d", Descricao: "teste"}, false},
		{"ERR: Valor Min", dtos.TransacaoRequestDto{Valor: 0, Tipo: "d", Descricao: "teste"}, true},
		{"ERR: Tipo vazio", dtos.TransacaoRequestDto{Valor: 1, Tipo: "", Descricao: "teste"}, true},
		{"ERR: Tipo invalido", dtos.TransacaoRequestDto{Valor: 1, Tipo: "x", Descricao: "teste"}, true},
		{"ERR: Descricao vazio", dtos.TransacaoRequestDto{Valor: 1, Tipo: "c", Descricao: ""}, true},
		{"ERR: Descricao invalido", dtos.TransacaoRequestDto{Valor: 1, Tipo: "x", Descricao: "12345678901"}, true},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if actualErr := tt.input.Valido(); (actualErr != nil) != tt.withErr {
				t.Fatalf("expected error %t, got %s", tt.withErr, actualErr)
			}
		})
	}
}

func BenchmarkEntity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		user := dtos.TransacaoRequestDto{Valor: 1, Tipo: "d", Descricao: "teste"}
		if err := user.Valido(); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}
