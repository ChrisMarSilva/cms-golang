package main

import (
	"time"    "github.com/joho/godotenv"

)func init() {
	err := godotenv.Load()
	if err != nil {
		log.Panicln(err)
	}
}


const SecretKey = "cms-golang.tnb.cripo.api.auth-secret-key" // JwtSecret

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	SSLMode        string

	JwtSecret    string        `mapstructure:"JWT_SECRET"`
	JwtExpiresIn time.Duration `mapstructure:"JWT_EXPIRED_IN"`
	JwtMaxAge    int           `mapstructure:"JWT_MAXAGE"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    config := models.Config{
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName:   os.Getenv("DB_NAME"),
        SSLMode:  os.Getenv("DB_SSLMODE"),
    }
	
	// viper.AddConfigPath(path)
	// viper.SetConfigType("env")
	// viper.SetConfigName("app")

	// viper.AutomaticEnv()

	// err = viper.ReadInConfig()
	// if err != nil {
	// 	return
	// }

	// err = viper.Unmarshal(&config)
	return
}


type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func New() *Config {
	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		JWTSecret:  os.Getenv("JWT_SECRET"),
	}
}