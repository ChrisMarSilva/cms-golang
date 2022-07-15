package main

import (
	"fmt"
	"github.com/ChrisMarSilva/cms.golang.teste.api.chi/pkg/api"
	"github.com/ChrisMarSilva/cms.golang.teste.api.chi/pkg/db"
	"log"
	"net/http"
	"os"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.api.chi
// go get -u github.com/go-chi/chi/v5
// go get -u github.com/go-pg/pg/v10
// go get -u github.com/go-pg/migrations/v8
// go mod tidy

// go run main.go

func main() {
	pgdb, err := db.NewDB()
	if err != nil {
		log.Printf("erro from database: %v\n", err)
		return
	}

	router := api.NewAPI(pgdb)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	err = http.ListenAndServe(fmt.Sprintf(":%s", port), router)
	if err != nil {
		log.Printf("erro from router: %v\n", err)
		return
	}

	log.Println("we're up and runnig!")
}
