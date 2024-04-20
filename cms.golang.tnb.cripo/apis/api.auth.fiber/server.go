package main

import (
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	_ "github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	_ "github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

type Server struct {
	app *fiber.App
}

func NewServer() *Server {
	return &Server{
		app: fiber.New(fiber.Config{
			AppName:     "Tamo em Cripto API Auth v1.0.0",
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		}),
	}
}

func (s *Server) Initialize() {
	app := ConfigRoutes(s.app)

	app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))
	app.Use(cors.New()) // app.Use(cors.New(cors.Config{AllowOrigins: "*", AllowMethods: "*", AllowHeaders: "*", AllowCredentials: true}))
	app.Use(csrf.New())
	//app.Use(adaptor.HTTPMiddleware(s.logMiddleware))
	app.Use(s.logMiddleware)
	//app.Get(healthcheck.DefaultLivenessEndpoint, healthcheck.NewHealthChecker())
	//app.Get(healthcheck.DefaultReadinessEndpoint, healthcheck.NewHealthChecker())
	app.Use(idempotency.New())
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{Format: "${cyan}[${time}] ${white}${latency} ${red}[${status}] ${blue}[${method}] ${white}${path} Error: ${red}${error}\n", TimeFormat: "2006-01-02T15:04:05.00000", TimeZone: "America/Sao_Paulo"}))
	log.SetLevel(log.LevelTrace)

	log.Info("Server running at port: 3001")
	log.Fatal(app.Listen(":3001"))
}

// func (s *Server) logMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("log middleware")
// 		next.ServeHTTP(w, r)
// 	})
// }

func (s *Server) logMiddleware(c fiber.Ctx) error {
	log.Info("Middleware: Request received")
	return c.Next()
}
