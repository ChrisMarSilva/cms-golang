package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimetro(t *testing.T) {
	esperado := 40.0                              // Arrange - Preparar o teste
	resultado := Perimetro(Retangulo{10.0, 10.0}) // Act - Rodar o teste
	assert.Equal(t, esperado, resultado)          // Assert - Verificar as asserções
}
func TestArea(t *testing.T) {
	esperado := 72.0
	retangulo := Retangulo{12.0, 6.0}
	resultado := retangulo.Area()
	assert.Equal(t, esperado, resultado)
}

func TestFormas1(t *testing.T) {
	verificaArea := func(t *testing.T, forma IForma, esperado float64) {
		t.Helper()
		resultado := forma.Area()
		assert.Equal(t, esperado, resultado)
	}

	t.Run("retângulos", func(t *testing.T) {
		retangulo := Retangulo{12.0, 6.0}
		verificaArea(t, retangulo, 72.0)
	})

	t.Run("círculos", func(t *testing.T) {
		circulo := Circulo{10}
		verificaArea(t, circulo, 314.1592653589793)
	})
}

func TestFormas2(t *testing.T) {
	testesArea := []struct {
		nome     string
		forma    IForma
		esperado float64
	}{
		{nome: "Retângulo", forma: Retangulo{Largura: 12, Altura: 6}, esperado: 72.0},
		{nome: "Círculo", forma: Circulo{Raio: 10}, esperado: 314.1592653589793},
		{nome: "Triângulo", forma: Triangulo{Base: 12, Altura: 6}, esperado: 36.0},
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			resultado := tt.forma.Area()
			assert.Equal(t, tt.esperado, resultado)
		})
	}
}
