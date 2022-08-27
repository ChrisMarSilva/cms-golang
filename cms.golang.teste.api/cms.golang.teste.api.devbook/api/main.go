package main

import (
	"log"
	"net/http"
	"strconv"
	//"encoding/base64"
    //"math/rand"

	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/config"
	"github.com/ChrisMarSilva/cms-golang-teste-dev-book-api/src/router"
)

// go mod init github.com/chrismarsilva/cms-golang-teste-dev-book-api
// go mod tidy
// go get -u github.com/gorilla/mux
// go get -u github.com/joho/godotenv
// go get -u github.com/go-sql-driver/mysql
// go get -u -v github.com/simukti/sqldb-logger
// go get -u github.com/simukti/sqldb-logger/logadapter/zerologadapter
// go get -u golang.org/x/crypto/bcrypt
// go get -u github.com/dgrijalva/jwt-go

// go run main.go

// func init() {
// 	chave := make([]byte, 64)
// 	_, err := rand.Read(chave)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	stringBase64 := base64.StdEncoding.EncodeToString(chave)
// 	log.Println(stringBase64)
// }

func main() {
	config.Carregar()
	porta := strconv.Itoa(config.Porta)
	log.Println("API - Listen port " + porta)

	r := router.Gerar()
	r.StrictSlash(true)
	log.Fatal(http.ListenAndServe(":"+porta, r))

}
