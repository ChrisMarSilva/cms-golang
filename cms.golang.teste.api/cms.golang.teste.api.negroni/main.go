package main

// go mod init github.com/chrismarsilva/cms.golang.teste.api.negroni
// go get -u github.com/urfave/negroni
// go mod tidy

// go run .

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(mux)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
