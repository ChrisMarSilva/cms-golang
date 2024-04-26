package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DbUrl string
	// DBHost     string `mapstructure:"POSTGRES_HOST"`
	// DBPort     string `mapstructure:"POSTGRES_PORT"`
	// DBUser     string `mapstructure:"POSTGRES_USER"`
	// DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	// DBName     string `mapstructure:"POSTGRES_DB"`
	// SSLMode    string `mapstructure:"POSTGRES_DB"`

	JwtSecret string
	// JwtExpiresIn time.Duration `mapstructure:"JWT_EXPIRED_IN"`
	// JwtMaxAge    int           `mapstructure:"JWT_MAXAGE"`
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
