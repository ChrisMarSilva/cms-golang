package stores

import (
	"context"
	"log/slog"
	"strconv"
	"time"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/config_sqlc/tutorial"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/models"
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

func CloseDatabaseSqlC(ctx context.Context, db *pgx.Conn) {
	if db != nil {
		db.Close(ctx)
	}
}

func GetAllSqlC(ctx context.Context, queries *tutorial.Queries) ([]tutorial.TbPerson, error) {
	return queries.GetAll(ctx)
}

func FetchAllSqlC(ctx context.Context) {
	slog.Info("Connecting to database using SQLC...")

	db := NewDatabaseSqlC(ctx)
	defer CloseDatabaseSqlC(ctx, db)

	queries := tutorial.New(db)

	// persons, err := queries.GetAll(ctx)
	persons, err := GetAllSqlC(ctx, queries)
	if err != nil {
		slog.Error("Failed to query persons from database", slog.Any("error", err))
		return
	}

	for _, p := range persons {
		slog.Info("PersonModel", "person", p.Name)
	}

	slog.Info("Persons retrieved from database", slog.Int("count", len(persons)))
}

func ClearSqlC(ctx context.Context) {
	db := NewDatabaseSqlC(ctx)
	defer CloseDatabaseSqlC(ctx, db)

	queries := tutorial.New(db)
	queries.Clear(ctx)
}

func AddOneSqlC(ctx context.Context, count int) time.Duration {
	db := NewDatabaseSqlC(ctx)
	defer CloseDatabaseSqlC(ctx, db)
	queries := tutorial.New(db)

	start := time.Now()
	for i := 0; i < count; i++ {
		row := models.NewPerson("SqlC AddOne #" + strconv.Itoa(i))
		queries.Create(ctx, tutorial.CreateParams{ID: row.ID, Name: row.Name, CreatedAt: row.CreatedAt})
	}

	return time.Since(start)
}

func AddManySqlC(ctx context.Context, count int) time.Duration {
	db := NewDatabaseSqlC(ctx)
	defer CloseDatabaseSqlC(ctx, db)

	start := time.Now()

	rows := make([]*tutorial.CreateParams, 0, count)
	for i := 0; i < count; i++ {
		row := models.NewPerson("SqlC AddMany #" + strconv.Itoa(i))
		rows = append(rows, &tutorial.CreateParams{ID: row.ID, Name: row.Name, CreatedAt: row.CreatedAt})
	}

	tx, _ := db.Begin(ctx)
	defer tx.Rollback(ctx)

	queries := tutorial.New(db)
	qtx := queries.WithTx(tx)

	for _, row := range rows {
		qtx.Create(ctx, *row)
	}

	tx.Commit(ctx)

	return time.Since(start)
}
