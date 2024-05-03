package main

// Soma calcula o valor total dos números em um array
func Soma1(numeros []int) int {
	soma := 0
	for i := 0; i < len(numeros); i++ {
		soma += numeros[i]
	}

	return soma
}

func Soma2(numeros []int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}

	return soma
}

// SomaTudo calcula as respectivas somas de cada slice recebido
func SomaTudo1(numerosParaSomar ...[]int) []int {
	quantidadeDeNumeros := len(numerosParaSomar)
	somas := make([]int, quantidadeDeNumeros)

	for i, numeros := range numerosParaSomar {
		somas[i] = Soma2(numeros)
	}

	return somas
}

func SomaTudo2(numerosParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numerosParaSomar {
		somas = append(somas, Soma2(numeros))
	}

	return somas
}

func SomaTodoOResto(numerosParaSomar ...[]int) []int {
	var somas []int
	for _, numeros := range numerosParaSomar {
		if len(numeros) == 0 {
			somas = append(somas, 0)
		} else {
			final := numeros[1:]
			somas = append(somas, Soma2(final))
		}
	}

	return somas
}
