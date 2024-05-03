package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarteira(t *testing.T) {
	// Arrange - Preparar o teste
	esperado := Bitcoin(10)
	carteira := Carteira{}

	// Act - Agir - Rodar o teste (Executar a ação que está sendo testada)
	carteira.Depositar(Bitcoin(10))
	resultado := carteira.Saldo()

	fmt.Printf("O endereço do saldo no teste é %v \n", &carteira.saldo)

	// Assert - Verificar as asserções
	assert.Equal(t, esperado, resultado)
}

func TestCarteiras(t *testing.T) {
	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))
	})

	t.Run("Retirar com saldo suficiente", func(t *testing.T) {
		carteira := Carteira{Bitcoin(20)}
		erro := carteira.Retirar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))
		confirmaErroInexistente(t, erro)
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{saldoInicial}
		erro := carteira.Retirar(Bitcoin(100))
		confirmaSaldo(t, carteira, saldoInicial)
		confirmaErro(t, erro, ErroSaldoInsuficiente)
	})
}

func confirmaSaldo(t *testing.T, carteira Carteira, esperado Bitcoin) {
	t.Helper()
	resultado := carteira.Saldo()
	assert.Equal(t, esperado, resultado)
}

func confirmaErroInexistente(t *testing.T, resultado error) {
	t.Helper()
	// if resultado != nil {
	// 	t.Fatal("erro inesperado recebido")
	// }
	assert.Nil(t, resultado)
}

func confirmaErro(t *testing.T, resultado error, esperado error) {
	t.Helper()
	// if resultado == nil {
	// 	t.Fatal("esperava um erro, mas nenhum ocorreu")
	// }
	// if resultado != esperado {
	// 	t.Errorf("erro resultado %s, erro esperado %s", resultado, esperado)
	// }
	assert.NotNil(t, resultado)
	assert.Equal(t, esperado, resultado)
}
