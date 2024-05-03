package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//const numeros = [5]int{1, 2, 3, 4, 5}
//const esperado = 15

func TestSoma1(t *testing.T) {
	// Arrange - Preparar o teste
	esperado := 15
	numeros := []int{1, 2, 3, 4, 5}

	// Act - Rodar o teste
	resultado := Soma1(numeros)

	// Assert - Verificar as asserções
	assert.Equal(t, esperado, resultado)
}
func TestSoma2(t *testing.T) {
	// Arrange - Preparar o teste
	esperado := 15
	numeros := []int{1, 2, 3, 4, 5}

	// Act - Rodar o teste
	resultado := Soma2(numeros)

	// Assert - Verificar as asserções
	assert.Equal(t, esperado, resultado)
}
func TestSoma(t *testing.T) {
	t.Run("coleção de 5 números", func(t *testing.T) {
		esperado := 15
		numeros := []int{1, 2, 3, 4, 5}
		resultado := Soma1(numeros)
		assert.Equal(t, esperado, resultado)
	})

	t.Run("coleção de qualquer tamanho", func(t *testing.T) {
		esperado := 6
		numeros := []int{1, 2, 3}
		resultado := Soma2(numeros)
		assert.Equal(t, esperado, resultado)
	})
}

func TestSomaTudo1(t *testing.T) {
	resultado := SomaTudo1([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}
	// if !reflect.DeepEqual(resultado, esperado) {
	// 	t.Errorf("resultado %v esperado %v", resultado, esperado)
	// }
	assert.Equal(t, esperado, resultado)
}
func TestSomaTudo2(t *testing.T) {
	resultado := SomaTudo2([]int{1, 2}, []int{0, 9})
	esperado := []int{3, 9}
	assert.Equal(t, esperado, resultado)
}

func TestSomaTodoOResto(t *testing.T) {
	resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
	esperado := []int{2, 9}
	// if !reflect.DeepEqual(resultado, esperado) {
	//     t.Errorf("resultado %v, esperado %v", resultado, esperado)
	// }
	assert.Equal(t, esperado, resultado)
}

func TestSomaTodoOResto2(t *testing.T) {
	verificaSomas := func(t *testing.T, resultado, esperado []int) {
		t.Helper()
		// if !reflect.DeepEqual(resultado, esperado) {
		//     t.Errorf("resultado %v, esperado %v", resultado, esperado)
		// }
		assert.Equal(t, esperado, resultado)
	}

	t.Run("faz as somas de alguns slices", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{1, 2}, []int{0, 9})
		esperado := []int{2, 9}
		verificaSomas(t, resultado, esperado)
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		resultado := SomaTodoOResto([]int{}, []int{3, 4, 5})
		esperado := []int{0, 9}
		verificaSomas(t, resultado, esperado)
	})
}

func BenchmarkSoma1(b *testing.B) {
	numeros := []int{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Soma1(numeros)
	}
}
func BenchmarkSoma2(b *testing.B) {
	numeros := []int{1, 2, 3, 4, 5}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Soma2(numeros)
	}
}
func BenchmarkSomaTudo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SomaTudo1([]int{1, 2}, []int{0, 9})
	}
}
func BenchmarkSomaTudo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SomaTudo2([]int{1, 2}, []int{0, 9})
	}
}
