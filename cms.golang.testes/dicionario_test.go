package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBusca(t *testing.T) {
	// Arrange - Preparar o teste)
	dicionario := Dicionario{"teste": "isso é apenas um teste"}

	// Act - Rodar o teste
	resultado, _ := dicionario.Busca("teste")
	esperado := "isso é apenas um teste"

	// Assert - Verificar as asserções
	comparaStrings(t, resultado, esperado)
}

func TestBuscas(t *testing.T) {
	dicionario := Dicionario{"teste": "isso é apenas um teste"}

	t.Run("palavra conhecida", func(t *testing.T) {
		resultado, _ := dicionario.Busca("teste")
		esperado := "isso é apenas um teste"
		comparaStrings(t, resultado, esperado)
	})

	t.Run("palavra desconhecida", func(t *testing.T) {
		_, erro := dicionario.Busca("desconhecida")
		comparaErro(t, erro, ErrNaoEncontrado)
	})
}

func comparaStrings(t *testing.T, resultado, esperado string) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s', dado '%s'", resultado, esperado, "teste")
	}
	assert.Equal(t, esperado, resultado)
}

func comparaErro(t *testing.T, resultado, esperado error) {
	t.Helper()
	if resultado != esperado {
		t.Errorf("resultado erro '%s', esperado '%s'", resultado, esperado)
	}
}

func TestAdiciona(t *testing.T) {
	dicionario := Dicionario{}
	dicionario.Adiciona("teste", "isso é apenas um teste")

	esperado := "isso é apenas um teste"
	resultado, err := dicionario.Busca("teste")
	if err != nil {
		t.Fatal("não foi possível encontrar a palavra adicionada:", err)
	}

	if esperado != resultado {
		t.Errorf("resultado '%s', esperado '%s'", resultado, esperado)
	}
}

func TestAdiciona2(t *testing.T) {
	dicionario := Dicionario{}
	palavra := "teste"
	definicao := "isso é apenas um teste"

	dicionario.Adiciona(palavra, definicao)

	comparaDefinicao(t, dicionario, palavra, definicao)
}

func comparaDefinicao(t *testing.T, dicionario Dicionario, palavra, definicao string) {
	t.Helper()

	resultado, err := dicionario.Busca(palavra)
	if err != nil {
		t.Fatal("deveria ter encontrado palavra adicionada:", err)
	}

	if definicao != resultado {
		t.Errorf("resultado '%s',  esperado '%s'", resultado, definicao)
	}
}

func TestAdicionas(t *testing.T) {
	t.Run("palavra nova", func(t *testing.T) {
		dicionario := Dicionario{}
		palavra := "teste"
		definicao := "isso é apenas um teste"

		err := dicionario.Adiciona(palavra, definicao)

		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, definicao)
	})

	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario := Dicionario{palavra: definicao}
		err := dicionario.Adiciona(palavra, "teste novo")

		comparaErro(t, err, ErrPalavraExistente)
		comparaDefinicao(t, dicionario, palavra, definicao)
	})

	t.Run("palavra existente", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		novaDefinicao := "nova definição"
		dicionario := Dicionario{palavra: definicao}
		err := dicionario.Atualiza(palavra, novaDefinicao)

		comparaErro(t, err, nil)
		comparaDefinicao(t, dicionario, palavra, novaDefinicao)
	})

	t.Run("palavra nova", func(t *testing.T) {
		palavra := "teste"
		definicao := "isso é apenas um teste"
		dicionario := Dicionario{}

		err := dicionario.Atualiza(palavra, definicao)

		comparaErro(t, err, ErrPalavraInexistente)
	})
}

func TestUpdate(t *testing.T) {
	palavra := "teste"
	definicao := "isso é apenas um teste"
	dicionario := Dicionario{palavra: definicao}
	novaDefinicao := "nova definição"

	dicionario.Atualiza(palavra, novaDefinicao)

	comparaDefinicao(t, dicionario, palavra, novaDefinicao)
}

func TestDeleta(t *testing.T) {
	palavra := "teste"
	dicionario := Dicionario{palavra: "definição de teste"}

	dicionario.Deleta(palavra)

	_, err := dicionario.Busca(palavra)
	if err != ErrNaoEncontrado {
		t.Errorf("espera-se que '%s' seja deletado", palavra)
	}
}
