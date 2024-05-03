package main

import (
	"fmt"
)

func Ola(nome, linguagem string) string {
	if nome == "" {
		nome = "Mundo"
	}
	return fmt.Sprintf("%s, %s", cumprimento(linguagem), nome)
}

var cumprimentos = map[string]string{
	"br": "Olá",
	"fr": "Bonjour",
	"es": "Hola",
}

func cumprimento(linguagem string) string {
	cumprimento, ok := cumprimentos[linguagem]
	if !ok {
		cumprimento = "Hello"
	}

	return cumprimento
}
