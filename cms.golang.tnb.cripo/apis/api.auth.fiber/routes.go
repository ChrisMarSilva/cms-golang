package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/timeout"
)

func ConfigRoutes(app *fiber.App) *fiber.App {
	//Database
	db, err := GetDatabase()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %s", err.Error())
	}
	// defer db.Close()

	//Repository
	userRepo := NewUserRepository(db)

	//Service
	userServ := NewUserService(*userRepo)

	//Handle
	userHandler := NewUserHandler(*userServ)

	//routes

	app.Get("/", timeout.New(userHandler.Home, 5*time.Second))
	//app.Get("/metrics", monitor.New())

	routes := app.Group("/api/v1/auth")
	routes.Post("/login", userHandler.Login)
	routes.Post("/logout", userHandler.Logout)
	routes.Get("/refresh", userHandler.Refresh)
	routes.Get("/verify", userHandler.Verify)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	return app
}
