package main

import (
	"fmt"
	"log"

	"github.com/chrismarsilva/rinha-backend-2024/internals/route"
	"github.com/spf13/viper"
)

// go mod init github.com/chrismarsilva/rinha-backend-2024
// go get -u github.com/gofiber/fiber/v2
// go get -u github.com/jmoiron/sqlx
// go get -u github.com/goccy/go-json
// go get -u github.com/spf13/viper
// go mod tidy

// go run main.go

//http://127.0.0.1:3000/clientes/2/transacoes/
//http://127.0.0.1:3000/clientes/2/extrato/

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	log.Println("env.PORT", viper.Get("PORT"))
	log.Println("env.DATABASE_DRIVER", viper.Get("DATABASE_DRIVER"))
	log.Println("env.DATABASE_URL", viper.Get("DATABASE_URL"))
	log.Println("env.DATABASE_MAX_CONNECTIONS", viper.Get("DATABASE_MAX_CONNECTIONS"))
	log.Println("env.MENSAGEM", viper.Get("MENSAGEM"))
}

func main() {
	app := route.NewRoutes()
	log.Fatal(app.Listen(viper.GetString("PORT")))

	/*

		"goHexagonalBlog/internal/core/services"
		"goHexagonalBlog/internal/handlers"
		"goHexagonalBlog/internal/repositories"
		"goHexagonalBlog/internal/server"


		userRepository := repositories.NewUserRepository()
		userService := services.NewUserService(userRepository)
		userHandlers := handlers.NewUserHandlers(userService)
		server.Initialize(userHandlers)

	*/

}
