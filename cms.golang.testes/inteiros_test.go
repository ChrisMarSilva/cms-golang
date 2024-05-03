package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdicionador(t *testing.T) {
	// Arrange - Preparar o teste
	esperado := 4

	// Act - Rodar o teste
	resultado := Adiciona(2, 2)

	// Assert - Verificar as asserções
	assert.Equal(t, esperado, resultado)
}

func ExampleAdiciona() {
	soma := Adiciona(1, 5)
	fmt.Println(soma)
	// Output: 6
}
