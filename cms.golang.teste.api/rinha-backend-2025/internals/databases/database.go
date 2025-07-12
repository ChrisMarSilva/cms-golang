package databases

import (
	"context"
	"log"
	"log/slog"
	"time"

	"github.com/chrismarsilva/rinha-backend-2025/internals/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Conn *pgxpool.Pool

func Connect(cfg *utils.Config) *pgxpool.Pool {
	if Conn != nil {
		log.Println("getting existent connection")
		return Conn
	} else {
		log.Println("opening connections")
	}

	config, err := pgxpool.ParseConfig(cfg.DbUri)
	if err != nil {
		log.Fatalf("Erro ao analisar a string de conexão: %v", err)
	}

	config.MinConns = 49                      // Defina o número mínimo de conexões
	config.MaxConns = 50                      // Define o número máximo de conexões abertas// Defina o número máximo de conexões
	config.MaxConnLifetime = 2 * time.Hour    // Exemplo: fecha conexões após 5 minutos
	config.MaxConnIdleTime = 10 * time.Minute // Exemplo: fecha conexões ociosas após 2 minutos
	config.HealthCheckPeriod = 10 * time.Minute

	ctx := context.Background()
	Conn, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		slog.Error("Error connecting pool to database", slog.String("error", err.Error()))
		return nil
	}

	if err = Conn.Ping(ctx); err != nil {
		slog.Error("Unable to ping database", slog.String("error", err.Error()))
		return nil
	}

	return Conn
}

func Close() {
	if Conn == nil {
		return
	}
	Conn.Close()
}
