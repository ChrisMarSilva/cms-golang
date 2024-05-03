package main

import "strings"

const quantidadeRepeticoes = 5

func Repetir1(caractere string) string {
	var repeticoes string
	for i := 0; i < quantidadeRepeticoes; i++ {
		repeticoes += caractere
	}

	return repeticoes
}

func Repetir2(caractere string) string {
	return strings.Repeat(caractere, quantidadeRepeticoes)
}

func Repetir3(caractere string) string {
	repeticoes := make([]byte, quantidadeRepeticoes)
	bytes := []byte(caractere)

	for i := range repeticoes {
		//repeticoes[i] = bytes
		copy(repeticoes[i*len(bytes):], bytes)
	}

	return string(repeticoes)
}

func Repetir4(caractere string) string {
	repeticoes := make([]byte, quantidadeRepeticoes)
	bytes := []byte(caractere)

	for i := range repeticoes {
		repeticoes[i] = bytes[0]
	}

	return string(repeticoes)
}
