package routes

import (
	"os"

	"github.com/chrismarsilva/rinha-backend-2024/internals/databases"
	"github.com/chrismarsilva/rinha-backend-2024/internals/handlers"
	"github.com/chrismarsilva/rinha-backend-2024/internals/repositories"
	"github.com/chrismarsilva/rinha-backend-2024/internals/services"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"

	//"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	_ "github.com/lib/pq"
	_ "github.com/spf13/viper"
)

func NewRoutes() *fiber.App {

	app := fiber.New(fiber.Config{
		//ServerHeader: "Fiber",
		AppName: "Rinha Backend 2024 v1.0.0",
		// Prefork:       true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
		BodyLimit:   12 * 1024 * 1024,
	})

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{AllowOrigins: "*", AllowMethods: "*", AllowHeaders: "*", AllowCredentials: true}))
	app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))

	//app.Use(logger.New(logger.Config{Format: "${cyan}[${time}] ${white}${latency} ${red}[${status}] ${blue}[${method}] ${white}${path} Error: ${red}${error}\n", TimeFormat: "2006-01-02T15:04:05.00000", TimeZone: "America/Sao_Paulo"}))
	//app.Use(logger.New(logger.Config{Format: "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path} Error: ${error}\n", TimeFormat: "2006-01-02T15:04:05.00000", TimeZone: "America/Sao_Paulo"}))
	//app.Use(logger.New(logger.Config{Format: "[${time}]: ${ip} ${status} ${latency} ${method} ${path} Error: ${error}\n", TimeFormat: "2006-01-02T15:04:05.00000", TimeZone: "America/Sao_Paulo"}))

	setupConfigRoutes(app)

	return app
}

func setupConfigRoutes(app fiber.Router) {

	//Database
	driverDbWriter := databases.DatabasePostgres{}
	driverDbWriter.StartDbWriter()
	writer := driverDbWriter.GetDatabaseWriter()

	driverDbReader := databases.DatabasePostgres{}
	driverDbReader.StartDbReader()
	reader := driverDbReader.GetDatabaseReader()

	//Repository
	clientRepo := repositories.NewClientRepository(writer, reader)
	clientTransactionRepo := repositories.NewClientTransactionRepository(writer, reader)

	//Service
	clientServ := services.NewClientService(*clientRepo, *clientTransactionRepo)

	//Handle
	clientHandler := handlers.NewClientHandler(*clientServ)

	routes := app.Group("/clientes")
	routes.Post(":id/transacoes", clientHandler.CreateTransaction)
	routes.Get("/:id/extrato", clientHandler.GetExtract)

	app.Get("/msg", Menssage)
	app.Use(NotFound)
}

func Menssage(c fiber.Ctx) error {
	//return c.SendString(viper.GetString("MENSAGEM"))
	return c.SendString(os.Getenv("MENSAGEM"))
}

func NotFound(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotFound)
}
