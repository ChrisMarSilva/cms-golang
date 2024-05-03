package utils

import (
	"os"
	"strconv"

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
	//getEnv("PUBLIC_HOST", "http://localhost"),
	// fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
	cfg.UriPort = ":" + getEnv("PORT", "8080")                                       // os.Getenv("PORT")
	cfg.DbDriver = getEnv("DATABASE_DRIVER", "sqlite3")                              // os.Getenv("DATABASE_DRIVER")
	cfg.DbUri = getEnv("DATABASE_URI", "./banco.db")                                 // os.Getenv("DATABASE_URI")
	cfg.JwtSecretKey = getEnv("JWT_SECRET", "cms_tamo_em_cripo_api_auth_secret_key") // os.Getenv("JWT_SECRET_KEY")
	// JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600 * 24 * 7),

	return &cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
