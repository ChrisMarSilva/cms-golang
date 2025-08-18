package stores

import (
	"context"
	"log/slog"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/models"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewDatabasePgxPool(ctx context.Context) *pgxpool.Pool {
	pgxConfig, err := pgxpool.ParseConfig(utils.DBUri)
	if err != nil {
		slog.Error("Unable to parse database config", slog.Any("error", err))
		return nil
	}

	conn, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		slog.Error("Unable to connect to database", slog.Any("error", err))
		return nil
	}

	if err = conn.Ping(ctx); err != nil {
		slog.Error("Unable to ping database", slog.Any("error", err))
		return nil
	}

	return conn
}

func GetAllPgxPool(ctx context.Context, db *pgxpool.Pool) (pgx.Rows, error) {
	return db.Query(ctx, utils.DBQuery)
}

func FetchAllPgxPool(ctx context.Context) {
	slog.Info("Connecting to database using pgx pool...")

	dbPgxPool := NewDatabasePgxPool(ctx)
	defer dbPgxPool.Close()

	rows, err := GetAllPgxPool(ctx, dbPgxPool) // dbPgxPool.Query(ctx, utils.DBQuery)
	if err != nil {
		slog.Error("Failed to query persons from database", slog.Any("error", err))
		return
	}
	defer rows.Close()

	// persons := make([]*models.PersonModel, 0)

	for rows.Next() {
		var person models.PersonModel
		if err := rows.Scan(&person.ID, &person.Name, &person.CreatedAt); err != nil {
			slog.Error("Failed to scan person data", slog.Any("error", err))
			return
		}

		slog.Info(person.Name)
		//persons = append(persons, &person)
	}

	if err := rows.Err(); err != nil {
		slog.Error("Failed to iterate over persons", slog.Any("error", err))
		return
	}

	//slog.Info("Persons retrieved from database", slog.Int("count", len(persons)))
}
