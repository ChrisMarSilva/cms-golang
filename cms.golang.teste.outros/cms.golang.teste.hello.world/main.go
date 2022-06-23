package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/ChrisMarSilva/cms.golang.teste.hello.world/math"
	"github.com/google/uuid"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.hello.world
// go mod tidy

// go run main.go

func main() {

	resultado := Soma(1, 2)
	fmt.Println("resultado", resultado)
	fmt.Printf("%T", resultado)

	resultadoFloat := math.SomaFloat(1.5, 2.5)
	fmt.Println("resultadoFloat", resultadoFloat)
	fmt.Printf("%T", resultadoFloat)

	fmt.Println("Hello, Docker #1")
	id := uuid.New()
	fmt.Println("uuid: ", id.String())
	message, err := Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

}

func Soma(a int, b int) int {
	return a + b
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
