package main

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUrl string
	JwtSecret string
}

func NewConfig(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return &Config{}, err
	}

	var cfg Config
	cfg.DbUrl = os.Getenv("DB_URL")
	cfg.JwtSecret = os.Getenv("JWT_SECRET")

	return &cfg, nil
}
