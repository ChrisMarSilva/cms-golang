package routes

import (
	"github.com/chrismarsilva/rinha-backend-2024/internals/databases"
	"github.com/chrismarsilva/rinha-backend-2024/internals/handlers"
	"github.com/chrismarsilva/rinha-backend-2024/internals/models"
	"github.com/chrismarsilva/rinha-backend-2024/internals/repositories"
	"github.com/chrismarsilva/rinha-backend-2024/internals/services"
	"github.com/chrismarsilva/rinha-backend-2024/internals/utils"
	"github.com/gofiber/fiber/v3"
)

func ConfigRoutes(app *fiber.App, cfg *utils.Config) *fiber.App {

	//Database
	driverDb := databases.DatabasePostgres{}
	driverDb.StartDbConnPgx(cfg)
	db := driverDb.GetDatabaseConnPgx()

	// driverDb := databases.DatabasePostgres{}
	// driverDb.StartDbConn(cfg)
	// db := driverDb.GetDatabaseConn()

	// driverDbWriter := databases.DatabasePostgres{}
	// driverDbWriter.StartDbWriter(cfg)
	// writer := driverDbWriter.GetDatabaseWriter()

	// driverDbReader := databases.DatabasePostgres{}
	// driverDbReader.StartDbReader(cfg)
	// reader := driverDbReader.GetDatabaseReader()

	//Repository
	clientRepo := repositories.NewClientRepository(db)
	clientTransactionRepo := repositories.NewClientTransactionRepository(db)

	//Service
	clientServ := services.NewClientService(db, *clientRepo, *clientTransactionRepo)

	//Handle
	clientHandler := handlers.NewClientHandler(*clientServ)

	batchChannelCreateTransaction := make(chan models.Job, cfg.MaxQueueSize)
	go clientHandler.ProcessBatch(batchChannelCreateTransaction, clientHandler.CreateTransactionBatch)
	//time.Sleep(3 * time.Second) // wait for db is up

	routes := app.Group("/clientes")
	//routes.Post(":id/transacoes", clientHandler.CreateTransaction)

	routes.Post(":id/transacoes", func(c fiber.Ctx) error {
		return clientHandler.CreateTransactionBatch(c, batchChannelCreateTransaction)
	})

	routes.Get("/:id/extrato", clientHandler.GetExtract)

	app.Get("/msg", func(c fiber.Ctx) error {
		return c.SendString(cfg.Msg)
	})

	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	return app
}
