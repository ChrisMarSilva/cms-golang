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

// go install migrate
// go install github.com/cosmtrek/air@latest
// air init
// air

import (
	"time"

	"github.com/gofiber/fiber/v3/middleware/session"
	"github.com/gofiber/storage/sqlite3"
)

var (
	secretKey = "cms-golang.tnb.cripo.api.auth-secret-key"
	store     = session.New(session.Config{
		Expiration: 24 * time.Hour,
		KeyLookup:  "cookie:session_id",
		Storage:    sqlite3.New(sqlite3.Config{Database: "./banco.db"}),
	})
)

func main() {
	app := NewServer()
	app.Initialize()
}

/*

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/google/uuid"
)







go get -u github.com/uber-go/zap
go get -u github.com/prometheus/client_golang/prometheus
go get -u github.com/opentracing/opentracing-go


import (
"github.com/uber-go/zap"
"github.com/uber-go/zap/zapcore"
)

var logger *zap.Logger

func init() {
encoderConfig := zap.NewJSONEncoderConfig()
encoderConfig.EncodeTime = zap.TimeEncoderRFC3339
encoderConfig.EncodeLevel = zap.LevelEncoderLowerCase
fileWriter, _ := os.OpenFile("api.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

logger = zap.New(
	zap.NewStdOutWriter(zap.LevelInfo),
	zap.NewStreamWriter(fileWriter, zap.LevelWarn),
	zap.NewEncoderConfig(&encoderConfig),
)
}

logger.Info("Requisição recebida", zap.String("método", c.Request.Method), zap.String("caminho", c.Request.URL.Path))
logger.Debug("Consultando banco de dados", zap.String("consulta", consulta))
logger.Info("Resposta enviada", zap.Int("status", c.Status), zap.Int("tamanho", len(resposta)))



requestCounter := prometheus.NewCounter(prometheus.CounterOpts{
Name:      "api_requests_total",
Help:      "Número total de requisições recebidas pela API",
Namespace: "api",
})


logger, _ := zap.NewProduction()
defer logger.Sync() // flushes buffer, if any
sugar := logger.Sugar()
sugar.Infow("failed to fetch URL",
// Structured context as loosely typed key-value pairs.
"url", url,
"attempt", 3,
"backoff", time.Second,
)
sugar.Infof("Failed to fetch URL: %s", url)


logger, _ := zap.NewProduction()
defer logger.Sync()
logger.Info("failed to fetch URL",
// Structured context as strongly typed Field values.
zap.String("url", url),
zap.Int("attempt", 3),
zap.Duration("backoff", time.Second),
)

requestDurationHistogram := prometheus.NewHistogram(prometheus.HistogramOpts{
Name:      "api_request_duration_seconds",
Help:      "Tempo de resposta das requisições da API (em segundos)",
Namespace: "api",
Buckets: []float64{0.01, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
})

requestErrorsCounter := prometheus.NewCounter(prometheus.CounterOpts{
Name:      "api_request_errors_total",
Help:      "Número total de erros nas requisições da API",
Namespace: "api",
})

prometheus.DefaultRegister



*/
