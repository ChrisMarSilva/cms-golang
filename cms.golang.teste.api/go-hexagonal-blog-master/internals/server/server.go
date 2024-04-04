package server

import (
	"goHexagonalBlog/internal/core/ports"
	"log"

	fiber "github.com/gofiber/fiber/v2"
)

type Server struct {
	userHandlers ports.IUserHandlers
}

func NewServer(uHandlers ports.IUserHandlers) *Server {
	return &Server{userHandlers: uHandlers}
}

func (s *Server) Initialize() {
	app := fiber.New()
	v1 := app.Group("/v1")

	userRoutes := v1.Group("/user")
	userRoutes.Post("/login", s.userHandlers.Login)
	userRoutes.Post("/register", s.userHandlers.Register)

	err := app.Listen(":5000")
	if err != nil {
		log.Fatal(err)
	}
}
