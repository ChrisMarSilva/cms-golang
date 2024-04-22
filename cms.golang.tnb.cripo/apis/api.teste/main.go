package main

// go mod init github.com/chrismarsilva/cms.golang.tnb.cripo.api.teste
// go get -u github.com/gofiber/fiber/v3
// go get -u github.com/gofiber/fiber/v3/middleware/logger
// go run .

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":3000")
}
