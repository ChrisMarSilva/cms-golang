package utils

import (
	"log"
	"os"
	"runtime"
	"time"

	"github.com/joho/godotenv"
)

// type IConfig interface {
// }

type Config struct {
	UriPort       string
	DbDriver      string
	DbUri         string
	DbMaxConn     string
	NumBatch      int
	NumWorkers    int
	WorkerTimeout time.Duration
	Msg           string
}

func NewConfig() *Config {
	production := os.Getenv("GO_ENVIRONMENT") == "production"
	if !production {
		log.Println("loading .env file")
		if err := godotenv.Load("./../../.env"); err != nil {
			if err := godotenv.Load("./../.env"); err != nil {
				if err := godotenv.Load("./.env"); err != nil {
					if err := godotenv.Load(".env"); err != nil {
						log.Fatal("Error", err)
					}
				}
			}
		}
	}

	var cfg Config

	cfg.UriPort = ":" + os.Getenv("PORT")
	cfg.DbDriver = os.Getenv("DATABASE_DRIVER")
	cfg.DbUri = os.Getenv("DATABASE_URL")
	cfg.DbMaxConn = os.Getenv("DATABASE_MAX_CONNECTIONS")
	cfg.NumBatch = 1000
	cfg.NumWorkers = runtime.GOMAXPROCS(0) * 2
	cfg.WorkerTimeout = 1 * time.Second
	cfg.Msg = os.Getenv("MENSAGEM")

	if cfg.DbMaxConn == "" {
		cfg.DbMaxConn = "50"
	}

	log.Println("Config.PORT", cfg.UriPort)
	log.Println("Config.DATABASE_DRIVER", cfg.DbDriver)
	log.Println("Config.DATABASE_URL", cfg.DbUri)
	log.Println("Config.DATABASE_MAX_CONNECTIONS", cfg.DbMaxConn)
	log.Println("Config.NUM_BATCH", cfg.NumBatch)
	log.Println("Config.NUM_WORKERS", cfg.NumWorkers)
	log.Println("Config.WORKER_TIMEOUT", cfg.WorkerTimeout)
	log.Println("Config.MENSAGEM", cfg.Msg)

	//connStr := "host=%s port=%s user=%s dbname=%s sslmode=%s"
	//connStr = fmt.Sprintf(connStr, host, port, user, dbname, sslmode)
	return &cfg
}
