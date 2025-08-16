package stores

import (
	"context"
	"log/slog"
	"os"

	"github.com/chrismarsilva/cms.project.1million/internal/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	logger *slog.Logger
	config *utils.Config
	Conn   *pgxpool.Pool
}

func NewDatabase(logger *slog.Logger, config *utils.Config) *Database {
	pgxConfig, err := pgxpool.ParseConfig(config.DatabaseUrl)
	if err != nil {
		logger.Error("Unable to parse database config", slog.Any("error", err))
		os.Exit(1)
	}

	pgxConfig.MaxConns = 10000    // Defina o número máximo de conexões
	pgxConfig.MinConns = 100      // Defina o número mínimo de conexões
	pgxConfig.MaxConnIdleTime = 0 // Desative o tempo máximo de inatividade da conexão

	// conn, err := pgxpool.New(context.Background(), cfg.DbUri)
	// conn, err := pgxpool.Connect(context.Background(), cfg.DbUri)
	conn, err := pgxpool.NewWithConfig(context.Background(), pgxConfig)
	if err != nil {
		logger.Error("Unable to connect to database", slog.Any("error", err))
		os.Exit(1)
	}

	if err = conn.Ping(context.Background()); err != nil {
		logger.Error("Unable to ping database", slog.Any("error", err))
		os.Exit(1)
	}

	// 	conn.Config().MaxConns = 1000     // Define o número máximo de conexões abertas// Defina o número máximo de conexões
	// 	conn.Config().MinConns = 100      // Defina o número mínimo de conexões
	// 	conn.Config().MaxConnIdleTime = 0 // Define o tempo máximo de ociosidade para 5 minutos// Desative o tempo máximo de inatividade da conexão
	// 	conn.Config().MaxConnLifetime = 0 // Define o tempo máximo de vida (lifetime) para 30 minutos// Desative o tempo máximo de espera para adquirir uma conexão
	// 	conn.Config().MaxConnIdleTime = time.Minute * 3

	return &Database{
		logger: logger,
		config: config,
		Conn:   conn,
	}
}

func (d *Database) Close() {
	if d.Conn != nil {
		d.Conn.Close()
	}
}
