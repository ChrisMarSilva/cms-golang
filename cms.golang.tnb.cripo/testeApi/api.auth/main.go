package main

// go mod init github.com/chrismarsilva/cms.golang.tnb.cripo.api.auth2
// go get -u github.com/gofiber/fiber/v3
// go get -u github.com/gofiber/fiber/v3/middleware/cors
// go get -u github.com/golang-jwt/jwt
// go get -u github.com/mattn/go-sqlite3
// go get -u github.com/gofiber/storage/sqlite3
// go get -u github.com/google/uuid
// go get -u github.com/goccy/go-json
// go get -u github.com/golang-migrate/migrate/v4
// go get -u github.com/golang-migrate/migrate/v4/database
// go get -u github.com/stretchr/testify/require
// go get -u github.com/stretchr/testify/suite
// go get -u github.com/stretchr/testify/assert
// go mod tidy
// go run main.go // go run .

// g0 install migrate
// go install github.com/cosmtrek/air@latest
// air init
// air

import (
	"time"
	_ "database/sql"

	_ "github.com/google/uuid"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/gofiber/fiber/v3/middleware/compress"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/csrf"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/idempotency"
	"github.com/gofiber/fiber/v3/middleware/logger"

	//"github.com/gofiber/fiber/v3/middleware/monitor"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/fiber/v3/middleware/timeout"
	"github.com/gofiber/storage/sqlite3"

	
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/hashicorp/go-multierror"
	_ "github.com/mattn/go-sqlite3"
)

var (
	secretKey = "cms-golang.tnb.cripo.api.auth2-secret-key"
	
	store     = session.New(session.Config{		
		Expiration:   24 * time.Hour,
		KeyLookup:    "cookie:session_id",
		//KeyGenerator: func() { return uuid.NewString()},
		//source:       "cookie",
		//sessionName:  "session_id",
		Storage: sqlite3.New(sqlite3.Config{Database: "./banco.auth.db"}),
	})
)

func init() {
	//store.RegisterType("jwt", jwtSessionStore{})
}

func main() {

	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})

	app.Use(compress.New(compress.Config{Level: compress.LevelBestSpeed}))
	app.Use(cors.New())
	app.Use(csrf.New())
	app.Use(adaptor.HTTPMiddleware(LogMiddleware))
	app.Get(healthcheck.DefaultLivenessEndpoint, healthcheck.NewHealthChecker())
	app.Get(healthcheck.DefaultReadinessEndpoint, healthcheck.NewHealthChecker())
	app.Use(idempotency.New()) // "X-Idempotency-Key"
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{Format: "${pid} ${status} - ${method} ${path}\n", TimeFormat: "02-Jan-2006", TimeZone: "America/New_York"}))
	log.SetLevel(log.LevelTrace)

	app.Post("/", timeout.New(HomeHandler, 5*time.Second))
	//app.Get("/metrics", monitor.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/login", LoginHandler)
	v1.Post("/logout", LogoutHandler)
	v1.Get("/refresh", RefreshHandler)
	v1.Get("/verify", VerifyHandler)

	log.Fatal(app.Listen(":3000"))
}
