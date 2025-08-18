package stores

import (
	"context"
	"log/slog"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/config_sqlc/tutorial"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/utils"
	"github.com/jackc/pgx/v5"
)

func NewDatabaseSqlC(ctx context.Context) *pgx.Conn {
	conn, err := pgx.Connect(ctx, utils.DBUri)
	if err != nil {
		slog.Error("Failed to connect to database", slog.Any("error", err))
		return nil
	}

	return conn
}

func FetchAllSqlC(ctx context.Context) {
	slog.Info("Connecting to database using SQLC...")

	dbSqlc := NewDatabaseSqlC(ctx)
	defer dbSqlc.Close(ctx)

	queries := tutorial.New(dbSqlc)

	persons, err := queries.GetAll(ctx)
	if err != nil {
		slog.Error("Failed to query persons from database", slog.Any("error", err))
		return
	}

	for _, p := range persons {
		slog.Info("PersonModel", "person", p.Name)
	}

	slog.Info("Persons retrieved from database", slog.Int("count", len(persons)))
}
