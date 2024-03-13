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
	UriPort   string
	DbDriver  string
	DbUri     string
	DbMaxConn string
	Msg       string

	NumBatch      int
	NumWorkers    int
	WorkerTimeout time.Duration

	MaxQueueSize int
	BatchSize    int
	BatchSleep   int
}

func NewConfig() *Config {
	production := os.Getenv("GO_ENVIRONMENT") == "production"
	if !production {
		//log.Println("loading .env file")
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
	cfg.Msg = os.Getenv("MENSAGEM")

	cfg.NumBatch = 1000
	cfg.NumWorkers = runtime.GOMAXPROCS(0) * 2
	cfg.WorkerTimeout = 1 * time.Second

	cfg.MaxQueueSize = 5000 // The size of job queue
	cfg.BatchSize = 1000    // The batch size
	cfg.BatchSleep = 2000   // Sleep time before batch

	if cfg.DbMaxConn == "" {
		cfg.DbMaxConn = "50"
	}

	// log.Println("Config.PORT", cfg.UriPort)
	// log.Println("Config.DATABASE_DRIVER", cfg.DbDriver)
	// log.Println("Config.DATABASE_URL", cfg.DbUri)
	// log.Println("Config.DATABASE_MAX_CONNECTIONS", cfg.DbMaxConn)
	// log.Println("Config.MENSAGEM", cfg.Msg)
	// log.Println("Config.NUM_BATCH", cfg.NumBatch)
	// log.Println("Config.NUM_WORKERS", cfg.NumWorkers)
	// log.Println("Config.WORKER_TIMEOUT", cfg.WorkerTimeout)
	// log.Println("Config.MAXQUEUESIZE", cfg.maxQueueSize)
	// log.Println("Config.BATCHSIZE", cfg.batchSize)
	// log.Println("Config.BATCHSLEEP", cfg.batchSleep)

	//connStr := "host=%s port=%s user=%s dbname=%s sslmode=%s"
	//connStr = fmt.Sprintf(connStr, host, port, user, dbname, sslmode)
	return &cfg
}
