package routes

import (
	"github.com/chrismarsilva/rinha-backend-2024/internals/databases"
	"github.com/chrismarsilva/rinha-backend-2024/internals/handlers"
	"github.com/chrismarsilva/rinha-backend-2024/internals/repositories"
	"github.com/chrismarsilva/rinha-backend-2024/internals/services"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func NewRoutes() *fiber.App {

	app := fiber.New(fiber.Config{
		//ServerHeader: "Fiber",
		AppName: "Rinha Backend 2024 v1.0.0",
		// Prefork:       true,
		// CaseSensitive: true,
		// StrictRouting: true,
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{AllowOrigins: "*", AllowMethods: "*", AllowHeaders: "*", AllowCredentials: true}))
	app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))

	app.Use(logger.New(logger.Config{Format: "${cyan}[${time}] ${white}${latency} ${red}[${status}] ${blue}[${method}] ${white}${path} Error: ${red}${error}\n", TimeFormat: "2006-01-02T15:04:05.00000", TimeZone: "America/Sao_Paulo"}))
	//app.Use(logger.New(logger.Config{Format: "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path} Error: ${error}\n", TimeFormat: "2006-01-02T15:04:05.00000", TimeZone: "America/Sao_Paulo"}))
	//app.Use(logger.New(logger.Config{Format: "[${time}]: ${ip} ${status} ${latency} ${method} ${path} Error: ${error}\n", TimeFormat: "2006-01-02T15:04:05.00000", TimeZone: "America/Sao_Paulo"}))

	setupConfigRoutes(app)

	return app
}

func setupConfigRoutes(app fiber.Router) {

	//Database
	driverDB := databases.DatabasePostgres{}
	driverDB.StartDB()
	db := driverDB.GetDatabase()

	//Repository
	clientRepo := repositories.NewClientRepository(db)
	clientTransactionRepo := repositories.NewClientTransactionRepository(db)

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
	return c.SendString(viper.GetString("MENSAGEM"))
}

func NotFound(c fiber.Ctx) error {
	return c.SendStatus(fiber.StatusNotFound)
}
