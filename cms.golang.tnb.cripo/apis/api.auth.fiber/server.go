package main

import (
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	app *fiber.App
}

func NewServer() *Server {
	return &Server{
		app: fiber.New(fiber.Config{
			AppName:     "Tamo em Cripto - API Auth - v1.0.0",
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
	}
}

func (s *Server) Initialize() {
	s.app = ConfigRoutes(s.app)

	s.app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))

	s.app.Use(cors.New())
	// app.Use(cors.New(cors.Config{AllowOrigins: "*", AllowMethods: "*", AllowHeaders: "*", AllowCredentials: true}))

	s.app.Use(healthcheck.New()) //live /ready
	// Or extend your config for customization
	// app.Use(healthcheck.New(healthcheck.Config{
	//     LivenessProbe: func(c *fiber.Ctx) bool {
	//         return true
	//     },
	//     LivenessEndpoint: "/live",
	//     ReadinessProbe: func(c *fiber.Ctx) bool {
	//         return serviceA.Ready() && serviceB.Ready() && ...
	//     },
	//     ReadinessEndpoint: "/ready",
	// }))

	//app.Use(idempotency.New())
	s.app.Use(idempotency.New(idempotency.Config{
		Lifetime:  30 * time.Minute,
		KeyHeader: "X-Idempotency-Key",
		Storage:   storage,
	}))

	s.app.Use(recover.New())

	// app.Use(logger.New())
	s.app.Use(logger.New(
		logger.Config{
			// 	Next:          nil,
			// 	Done:          nil,
			Format:     "${cyan}[${time}] ${pid} ${locals:requestid} ${white}${latency} ${red}[${status}] ${blue}[${method}] ${white}${path} Error: ${red}${error}\n",
			TimeFormat: "2006-01-02T15:04:05.00000",
			TimeZone:   "America/Sao_Paulo"},
		// TimeInterval:  500 * time.Millisecond,
	))

	// log.SetLevel(log.LevelDebug)

	log.Info("Server running at port: 3001")
	log.Fatal(s.app.Listen(":3001"))
}
