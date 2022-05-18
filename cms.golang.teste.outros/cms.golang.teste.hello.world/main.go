package main

// import "fmt"

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/trace"
	//"cms.golang.teste.hello.world/math"
)

//go run main.go
// go tool trace trace.out

func main() {

	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	resultado := Soma(1, 2)
	fmt.Println("resultado", resultado)
	fmt.Printf("%T", resultado)

	// resultadoFloat := math.SomaFloat(1.5, 2.5)
	// fmt.Println("resultadoFloat", resultadoFloat)
	// fmt.Printf("%T", resultadoFloat)

	// fmt.Println("Hello, Docker #1")
	// id := uuid.New()
	// fmt.Println("uuid: ", id.String())
	// message, err := Hello("")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(message)

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal("erro ao fazero get")

	}
	fmt.Print(res.StatusCode, res.Body)
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
