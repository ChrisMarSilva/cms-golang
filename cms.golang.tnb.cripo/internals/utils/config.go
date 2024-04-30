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

	//fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),

	var cfg Config
	cfg.UriPort = ":" + os.Getenv("PORT")
	cfg.DbDriver = os.Getenv("DATABASE_DRIVER")
	cfg.DbUri = os.Getenv("DATABASE_URI")
	cfg.JwtSecretKey = os.Getenv("JWT_SECRET_KEY")

	// return Config{
	// 	PublicHost:             getEnv("PUBLIC_HOST", "http://localhost"),
	// 	Port:                   getEnv("PORT", "8080"),
	// 	DBUser:                 getEnv("DB_USER", "root"),
	// 	DBPassword:             getEnv("DB_PASSWORD", "mypassword"),
	// 	DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
	// 	DBName:                 getEnv("DB_NAME", "ecom"),
	// 	JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
	// 	JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600 * 24 * 7),
	// }

	return &cfg, nil
}

// Gets the env by key or fallbacks
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
