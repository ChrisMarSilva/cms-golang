package server

import (
	"log"

	"github.com/chrismarsilva/rinha-backend-2024/internals/server/routes"
	"github.com/chrismarsilva/rinha-backend-2024/internals/utils"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

type Server struct {
	app *fiber.App
	cfg *utils.Config
}

func NewServer() *Server {
	return &Server{
		cfg: utils.NewConfig(),
		app: fiber.New(fiber.Config{
			AppName:     "Rinha Backend 2024 v1.0.0",
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
	}
}

func (s *Server) Initialize() {
	app := routes.ConfigRoutes(s.app, s.cfg)

	app.Use(recover.New())
	app.Use(cors.New(cors.Config{AllowOrigins: "*", AllowMethods: "*", AllowHeaders: "*", AllowCredentials: true}))
	app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))
	app.Use(logger.New(logger.Config{Format: "${cyan}[${time}] ${white}${latency} ${red}[${status}] ${blue}[${method}] ${white}${path} Error: ${red}${error}\n", TimeFormat: "2006-01-02T15:04:05.00000", TimeZone: "America/Sao_Paulo"}))
	//app.Use(logger.New(logger.Config{Format: "${cyan}[${time}] ${white}${pid} ${red}${status} ${blue}[${method}] ${white}${path} Error: ${error}\n", TimeFormat: "2006-01-02T15:04:05.00000", TimeZone: "America/Sao_Paulo"}))
	//app.Use(logger.New(logger.Config{Format: "[${time}]: ${ip} ${status} ${latency} ${method} ${path} Error: ${error}\n", TimeFormat: "2006-01-02T15:04:05.00000", TimeZone: "America/Sao_Paulo"}))

	log.Printf("Server running at port: %v", s.cfg.UriPort)
	log.Fatal(app.Listen(s.cfg.UriPort))
}
