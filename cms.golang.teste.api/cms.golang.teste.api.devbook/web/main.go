package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-web/src/config"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-web/src/router"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-web/src/utils"
)

// go mod init github.com/chrismarsilva/cms-golang-teste-dev-book-web
// go get -u github.com/gorilla/mux
// go get -u github.com/joho/godotenv
// go mod tidy

// go run main.go

func main() {

	utils.CarregarTemplates()

	config.Carregar()
	porta := strconv.Itoa(config.Porta)
	log.Println("WEB - Listen port " + porta)

	r := router.Gerar()
	r.StrictSlash(false)
	log.Fatal(http.ListenAndServe(":"+porta, r))

}
