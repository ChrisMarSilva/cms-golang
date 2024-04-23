package main

// go mod init github.com/chrismarsilva/cms.golang.tnb.cripo.api.auth
// go get -u github.com/gofiber/fiber/v2
// go get -u github.com/gofiber/fiber/v2/middleware/adaptor
// go get -u github.com/gofiber/fiber/v2/middleware/compress
// go get -u github.com/gofiber/fiber/v2/middleware/cors
// go get -u github.com/gofiber/fiber/v2/middleware/healthcheck
// go get -u github.com/gofiber/fiber/v2/middleware/idempotency
// go get -u github.com/gofiber/fiber/v2/middleware/logger
// go get -u github.com/gofiber/fiber/v2/middleware/monitor
// go get -u github.com/gofiber/fiber/v2/middleware/recover
// go get -u github.com/gofiber/fiber/v2/middleware/session
// go get -u github.com/gofiber/fiber/v2/middleware/timeout
// go get -u github.com/gofiber/fiber/v2/middleware/requestid
// go get -u github.com/gofiber/fiber/v2/log
// go get -u github.com/mattn/go-sqlite3
// go get -u github.com/gofiber/storage/sqlite3
// go get -u github.com/google/uuid
// go get -u github.com/goccy/go-json
// go get -u go.uber.org/zap
// go get -u github.com/gofiber/contrib/fiberzap/v2
// go get -u github.com/gofiber/contrib/swagger
// go get -u github.com/swaggo/fiber-swagger
// go get -u github.com/gofiber/contrib/jwt
// go get -u github.com/golang-jwt/jwt/v5
// go get -u golang.org/x/crypto/bcrypt


// go get -u github.com/golang-migrate/migrate/v4
// go get -u github.com/golang-migrate/migrate/v4/database
// go get -u github.com/stretchr/testify/require
// go get -u github.com/stretchr/testify/suite
// go get -u github.com/stretchr/testify/assert
// go mod tidy
// go run .

// go install migrate
// go install github.com/cosmtrek/air@latest
// air init
// air

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3"
)

var (
	storage = sqlite3.New(sqlite3.Config{Database: "./banco.db"})

	store = session.New(session.Config{
		Expiration: 24 * time.Hour,
		KeyLookup:  "cookie:session_id",
		Storage:    storage,
	})
)

func init() {
	// config, err := initializers.LoadConfig(".")
	// if err != nil {
	// 	log.Fatalln("Failed to load environment variables! \n", err.Error())
	// }
	// initializers.ConnectDB(&config)
}

func main() {
	app := NewServer()
	app.Initialize()
}

/*
 http://localhost:1323/swagger/index.html,


func main() {
    logger, _ := zap.NewProduction()
    ctx := context.Background()
    s := NewServiceA(logger)
    ctx = zax.Set(ctx, logger, []zap.Field{zap.String("trace_id", "my-trace-id")})
    s.funcA(ctx)
}

type ServiceA struct {
logger *zap.Logger
}

func NewServiceA(logger *zap.Logger) *ServiceA {
    return &ServiceA{
        logger: logger,
    }
}

func (s *ServiceA) funcA(ctx context.Context) {
    s.logger.Info("func A") // it does not contain trace_id, you need to add it manually
    zax.Get(ctx).Info("func A") // it will logged with "trace_id" = "my-trace-id"
}

*/
