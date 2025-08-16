package utils

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	GinMode string

	UriPort string
	NameApi string

	NumConsumerWorkers    int
	NumConsumerBatchSize  int
	NumPublisherWorkers   int
	NumPublisherBatchSize int

	DatabaseDriver string
	DatabaseUrl    string
	//DatabaseMaxConn int

	RedisAddr string
	RedisPwd  string

	RabbitMqUrl   string
	RabbitMqQueue string
	// RabbitMqHost  string
	// RabbitMqPort  int
	// RabbitMqUser  string
	// RabbitMqPass  string
	// RabbitMqVhost string
}

func NewConfig(logger *slog.Logger) *Config {
	godotenv.Load(".env")

	var cfg Config

	cfg.GinMode = GetEnvStrOrSetDefault("GIN_MODE", "release")

	cfg.UriPort = GetEnvStrOrSetDefault("URI_PORT", "9999")
	cfg.NameApi = GetEnvStrOrSetDefault("NAME_API", "456")

	cfg.NumConsumerWorkers = GetEnvIntOrSetDefault("NUM_CONSUMER_WORKERS", 100)
	cfg.NumConsumerBatchSize = GetEnvIntOrSetDefault("NUM_CONSUMER_BATCH_SIZE", 1000)
	cfg.NumPublisherWorkers = GetEnvIntOrSetDefault("NUM_PUBLISHER_WORKERS", 100)
	cfg.NumPublisherBatchSize = GetEnvIntOrSetDefault("NUM_PUBLISHER_BATCH_SIZE", 1000)

	cfg.DatabaseDriver = GetEnvStrOrSetDefault("DATABASE_DRIVER", "postgres")
	cfg.DatabaseUrl = GetEnvStrOrSetDefault("DATABASE_URL", "host=localhost port=5432 dbname=postgres user=postgres password=postgres sslmode=disable")

	cfg.RedisAddr = GetEnvStrOrSetDefault("REDIS_ADDR", "localhost:6379")
	cfg.RedisPwd = GetEnvStrOrSetDefault("REDIS_PWD", "123")

	cfg.RabbitMqUrl = GetEnvStrOrSetDefault("RABBIT_MQ_URL", "amqp://guest:guest@localhost:5672/")
	cfg.RabbitMqQueue = GetEnvStrOrSetDefault("RABBIT_MQ_DEFAULT_QUEUE", "queue.person")
	// cfg.RabbitMqHost = GetEnvStrOrSetDefault("RABBIT_MQ_HOST", "localhost")
	// cfg.RabbitMqPort = GetEnvIntOrSetDefault("RABBIT_MQ_PORT", 5672)
	// cfg.RabbitMqUser = GetEnvStrOrSetDefault("RABBIT_MQ_USERNAME", "guest")
	// cfg.RabbitMqPass = GetEnvStrOrSetDefault("RABBIT_MQ_PASSWORD", "guest")
	// cfg.RabbitMqVhost = GetEnvStrOrSetDefault("RABBIT_MQ_VHOST", "")

	logger.Info("config", slog.Any("c=", cfg))

	return &cfg
}

func GetEnvStrOrSetDefault(key, def string) string {
	if valueStr := os.Getenv(key); valueStr != "" {
		return valueStr
	}

	os.Setenv(key, def)
	return def
}

func GetEnvIntOrSetDefault(key string, def int) int {
	if valueStr := os.Getenv(key); valueStr != "" {
		if valueInt, err := strconv.Atoi(valueStr); err == nil {
			return valueInt
		}
	}

	os.Setenv(key, strconv.Itoa(def))
	return def
}
