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
	fmt.Println("resultado Soma:", resultado)
	//fmt.Printf("%T\n", resultado)

	resultadoFloat := math.SomaFloat(1.5, 2.5)
	fmt.Println("resultado Float:", resultadoFloat)
	// fmt.Printf("%T\n", resultadoFloat)

	fmt.Println("Hello, Docker #1")

	id := uuid.New()
	fmt.Println("resultado uuid: ", id.String())

	message, err := Hello("Chris")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

}

func Soma(a int, b int) (retorno int) {
	retorno = a + b
	return
}

func Hello(name string) (message string, erro error) {
	message = ""
	if name == "" {
		erro = errors.New("empty name")
		return
	}
	message = fmt.Sprintf("Hi, %v. Welcome!", name)
	return
}
