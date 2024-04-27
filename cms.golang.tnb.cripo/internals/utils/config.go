package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	UriPort string

	DbDriver string
	DbUri    string
	// DBHost     string `mapstructure:"POSTGRES_HOST"`
	// DBPort     string `mapstructure:"POSTGRES_PORT"`
	// DBUser     string `mapstructure:"POSTGRES_USER"`
	// DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	// DBName     string `mapstructure:"POSTGRES_DB"`
	// SSLMode    string `mapstructure:"POSTGRES_DB"`

	JwtSecretKey string
	// JwtExpiresIn time.Duration `mapstructure:"JWT_EXPIRED_IN"`
	// JwtMaxAge    int           `mapstructure:"JWT_MAXAGE"`
}

func NewConfig(path string) (*Config, error) {
	err := godotenv.Load(path)
	if err != nil {
		return &Config{}, err
	}

	var cfg Config
	cfg.UriPort = ":" + os.Getenv("PORT")
	cfg.DbDriver = os.Getenv("DATABASE_DRIVER")
	cfg.DbUri = os.Getenv("DATABASE_URI")
	cfg.JwtSecretKey = os.Getenv("JWT_SECRET_KEY")

	return &cfg, nil
}
