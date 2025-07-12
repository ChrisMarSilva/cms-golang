package utils

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	UriPort string
	//DbDriver  string
	DbUri string // DatabaseURL
	//DbMaxConn string
	//   RedisAddr       string
	//   DefaultProcURL  string
	//   FallbackProcURL string
	//   QueueCapacity   int
	//   WorkerCount     int
}

func NewConfig() *Config {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error", err)
	}

	// 	 capVal, err := strconv.Atoi(os.Getenv("QUEUE_CAPACITY"))
	//   if err != nil {
	//     log.Fatal("QUEUE_CAPACITY inválido")
	//   }
	//   wkVal, err := strconv.Atoi(os.Getenv("WORKER_COUNT"))
	//   if err != nil {
	//     log.Fatal("WORKER_COUNT inválido")
	//   }

	var cfg Config
	cfg.UriPort = fmt.Sprintf(":%v", os.Getenv("URI_PORT")) // getenv("PORT", "8080"),
	//cfg.DbDriver = os.Getenv("DATABASE_DRIVER")
	cfg.DbUri = os.Getenv("DATABASE_URL")
	//cfg.DbMaxConn = os.Getenv("DATABASE_MAX_CONNECTIONS")

	// log.Println("Config.PORT", cfg.UriPort)
	// log.Println("Config.DATABASE_DRIVER", cfg.DbDriver)
	// log.Println("Config.DATABASE_URL", cfg.DbUri)
	// log.Println("Config.DATABASE_MAX_CONNECTIONS", cfg.DbMaxConn)
	slog.Info("config", slog.Any("c=", cfg))

	return &cfg
}

func getEnvStr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func getEnvInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return def
}
