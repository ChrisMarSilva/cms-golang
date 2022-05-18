package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	StringConexaoBanco = ""
	Porta              = 0
	SecretKey          []byte
)

func Carregar() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// os.Getenv("DB_USUARIO")
	// os.Getenv("DB_SENHA")
	// os.Getenv("DB_NOME")

	StringConexaoBanco = "root:senha@tcp(localhost:3306)/database?parseTime=true&loc=Local"

	Porta, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Porta = 3000
	}

	SecretKey = []byte(os.Getenv("API_SECRET_KEY"))

}
