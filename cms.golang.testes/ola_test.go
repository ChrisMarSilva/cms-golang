package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOla(t *testing.T) {
	// Arrange - Preparar o teste)
	esperado := "Olá, Chris"

	// Act - Rodar o teste
	resultado := Ola("Chris", "br")

	// Assert - Verificar as asserções
	assert.Equal(t, esperado, resultado, "Erro #1")
}

func TestOla2(t *testing.T) {
	verificaMensagemCorreta := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		assert.Equal(t, esperado, resultado)
	}

	t.Run("diz olá para as pessoas", func(t *testing.T) {
		esperado := "Olá, Chris"                        // Arrange - Preparar o teste
		resultado := Ola("Chris", "br")                 // Act - Rodar o teste
		verificaMensagemCorreta(t, resultado, esperado) // Assert - Verificar as asserções
	})

	t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		esperado := "Olá, Mundo"                        // Arrange - Preparar o teste
		resultado := Ola("", "br")                      // Act - Rodar o teste
		verificaMensagemCorreta(t, resultado, esperado) // Assert - Verificar as asserções
	})

	t.Run("em espanhol", func(t *testing.T) {
		resultado := Ola("Elodie", "es")
		esperado := "Hola, Elodie"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz olá em francês", func(t *testing.T) {
		resultado := Ola("Lauren", "fr")
		esperado := "Bonjour, Lauren"
		verificaMensagemCorreta(t, resultado, esperado)
	})

	t.Run("diz olá em ingles", func(t *testing.T) {
		resultado := Ola("Jack", "us")
		esperado := "Hello, Jack"
		verificaMensagemCorreta(t, resultado, esperado)
	})
}
