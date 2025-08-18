package stores

import (
	"context"
	"log/slog"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/models"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/utils"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
	//_ "github.com/lib/pq"
)

func NewDatabaseSqlX(ctx context.Context) *sqlx.DB {
	db, err := sqlx.Connect(utils.DBDriverPgx, utils.DBUri)
	if err != nil {
		slog.Error("Failed to connect to database", slog.Any("error", err))
		return nil
	}

	return db
}

func FetchAllSqlX(ctx context.Context) {

	slog.Info("Connecting to database using SQLX...")

	dbSqlx := NewDatabaseSqlX(ctx)
	defer dbSqlx.Close()

	var persons []models.PersonModel
	if err := dbSqlx.Select(&persons, utils.DBQuery); err != nil {
		slog.Error("Failed to query persons from database", slog.Any("error", err))
		return
	}

	for _, p := range persons {
		slog.Info("PersonModel", "person", p.Name)
	}

	slog.Info("Persons retrieved from database", slog.Int("count", len(persons)))
}
