package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//const esperado = "aaaaa"

func TestRepetir1(t *testing.T) {
	esperado := "aaaaa"                  // Arrange - Preparar o teste
	resultado := Repetir1("a")           // Act - Rodar o teste
	assert.Equal(t, esperado, resultado) // Assert - Verificar as asserções
}
func TestRepetir2(t *testing.T) {
	esperado := "aaaaa"                  // Arrange - Preparar o teste
	resultado := Repetir2("a")           // Act - Rodar o teste
	assert.Equal(t, esperado, resultado) // Assert - Verificar as asserções
}
func TestRepetir3(t *testing.T) {
	esperado := "aaaaa"                  // Arrange - Preparar o teste
	resultado := Repetir3("a")           // Act - Rodar o teste
	assert.Equal(t, esperado, resultado) // Assert - Verificar as asserções
}
func TestRepetir4(t *testing.T) {
	esperado := "aaaaa"                  // Arrange - Preparar o teste
	resultado := Repetir4("a")           // Act - Rodar o teste
	assert.Equal(t, esperado, resultado) // Assert - Verificar as asserções
}

func BenchmarkRepetir1(b *testing.B) {
	//b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Repetir1("a")
	}
}
func BenchmarkRepetir2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir2("a")
	}
}
func BenchmarkRepetir3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir3("a")
	}
}
func BenchmarkRepetir4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir4("a")
	}
}
