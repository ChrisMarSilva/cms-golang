package utils

import (
	// "runtime"

	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	GinMode     string
	UriPort     string
	RedisAddr   string
	RedisPwd    string
	UrlDefault  string
	UrlFallback string
	NumWorkers  int
}

func NewConfig() *Config {
	godotenv.Load(".env")
	// if err := godotenv.Load(".env"); err != nil {
	// 	log.Fatal("Error", err)
	// }

	var cfg Config
	cfg.GinMode = GetEnvStrOrSetDefault("GIN_MODE", "release")
	cfg.UriPort = GetEnvStrOrSetDefault("URI_PORT", "9999")
	cfg.UrlDefault = GetEnvStrOrSetDefault("PROCESSOR_DEFAULT_URL", "http://localhost:8001")
	cfg.UrlFallback = GetEnvStrOrSetDefault("PROCESSOR_FALLBACK_URL", "http://localhost:8002")
	cfg.RedisAddr = GetEnvStrOrSetDefault("REDIS_ADDR", "localhost:6379")
	cfg.RedisPwd = GetEnvStrOrSetDefault("REDIS_PWD", "123")
	cfg.NumWorkers = GetEnvIntOrSetDefault("NUM_WORKERS", 2)

	// runtime.GOMAXPROCS(runtime.NumCPU())
	// utils.NumWorkers = 1 // runtime.NumCPU()

	// slog.Info("config", slog.Any("c=", cfg))
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
