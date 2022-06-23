package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/trace"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.performance.trace
// go mod tidy

// go run main.go
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

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		log.Fatal("erro ao fazero get")

	}
	fmt.Print(res.StatusCode, res.Body)
}

func Soma(a int, b int) int {
	return a + b
}
