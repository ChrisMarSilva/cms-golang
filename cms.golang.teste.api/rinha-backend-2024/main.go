package main

import (
	"log"
	"os"

	"github.com/chrismarsilva/rinha-backend-2024/internals/routes"
	"github.com/joho/godotenv"
)

// go mod init github.com/chrismarsilva/rinha-backend-2024
// go get -u github.com/gofiber/fiber/v3
// go get -u github.com/jmoiron/sqlx
// go get -u github.com/goccy/go-json
// go get -u github.com/joho/godotenv
// go mod tidy

// go run main.go
// go run ./cmd/api-server/main.go

//http://127.0.0.1:3000/clientes/2/transacoes/
//http://127.0.0.1:3000/clientes/2/extrato/

func init() {
	production := os.Getenv("GO_ENVIRONMENT") == "production"

	if !production {
		log.Println("loading .env file")

		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error c")
		}
	}

	log.Println("env.PORT", os.Getenv("PORT"))
	log.Println("env.DATABASE_DRIVER", os.Getenv("DATABASE_DRIVER"))
	log.Println("env.DATABASE_URL", os.Getenv("DATABASE_URL"))
	log.Println("env.DATABASE_MAX_CONNECTIONS", os.Getenv("DATABASE_MAX_CONNECTIONS"))
	log.Println("env.MENSAGEM", os.Getenv("MENSAGEM"))
}

func main() {
	app := routes.NewRoutes() //server.Initialize()
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

// docker-compose down
// docker-compose up -d --build

// docker rm -f $(docker ps -a -q)
// docker run -it rinha-backend-2024-api01:latest
