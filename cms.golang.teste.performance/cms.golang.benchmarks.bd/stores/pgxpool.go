package stores

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/models"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/utils"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	pgxslog "github.com/timtoronto634/pgx-slog"
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

	pgxConfig.ConnConfig.Tracer = pgxslog.NewTracer(slog.Default())

	return conn
}

func CloseDatabasePgxPool(ctx context.Context, db *pgxpool.Pool) {
	if db != nil {
		db.Close()
	}
}

func FetchAllPgxPool(ctx context.Context) {
	slog.Info("Connecting to database using pgx pool...")

	dbPgxPool := NewDatabasePgxPool(ctx)
	defer CloseDatabasePgxPool(ctx, dbPgxPool)

	rows, err := dbPgxPool.Query(ctx, utils.DBQuery)
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

func ClearPgxPool(ctx context.Context) {
	db := NewDatabasePgxPool(ctx)
	defer CloseDatabasePgxPool(ctx, db)

	db.Exec(ctx, utils.DBTruncate)
}

func AddOnePgxPool(ctx context.Context, count int) time.Duration {
	db := NewDatabasePgxPool(ctx)
	defer CloseDatabasePgxPool(ctx, db)

	start := time.Now()
	for i := 0; i < count; i++ {
		row := models.NewPerson("PgxPool AddOne #" + strconv.Itoa(i))
		db.Exec(ctx, utils.DBInsert, row.ID, row.Name, row.CreatedAt)
	}

	return time.Since(start)
}

func AddManyPgxPool(ctx context.Context, count int) time.Duration {
	db := NewDatabasePgxPool(ctx)
	defer CloseDatabasePgxPool(ctx, db)

	rows := make([]*models.PersonModel, 0, count)
	for i := 0; i < count; i++ {
		row := models.NewPerson("PgxPool1 AddMany #" + strconv.Itoa(i))
		rows = append(rows, row)
	}

	start := time.Now()
	tx, _ := db.Begin(ctx)
	defer tx.Rollback(ctx)

	for i := 0; i < count; i++ {
		tx.Exec(ctx, utils.DBInsert, rows[i].ID, rows[i].Name, rows[i].CreatedAt)
	}

	tx.Commit(ctx)

	return time.Since(start) // time.Now().Sub(start).Nanoseconds()
}

func AddManyPgxPool2(ctx context.Context, count int) time.Duration {
	db := NewDatabasePgxPool(ctx)
	defer CloseDatabasePgxPool(ctx, db)

	start := time.Now()
	batch := &pgx.Batch{}

	for i := 0; i < count; i++ {
		row := models.NewPerson("PgxPool2 AddMany #" + strconv.Itoa(i))
		batch.Queue(utils.DBInsert, row.ID, row.Name, row.CreatedAt)
	}

	results := db.SendBatch(ctx, batch)
	defer results.Close() // batchResults

	// for i := 0; i < count; i++ {
	// 	results.Exec()
	// 	//result, err := results.Exec()
	// 	//slog.Info("Executed batch insert", slog.Any("result", result.RowsAffected()), slog.Any("error", err))
	// 	//  if err != nil {
	// 	//   var pgErr *pgconn.PgError
	// 	//   if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
	// 	//       log.Printf("user %s already exists", user.Name)
	// 	//       continue
	// 	//   }
	// 	//   return fmt.Errorf("unable to insert row: %w", err)
	// 	// }
	// }

	return time.Since(start)
}

func AddManyPgxPool3(ctx context.Context, count int) time.Duration {
	db := NewDatabasePgxPool(ctx)
	defer CloseDatabasePgxPool(ctx, db)

	start := time.Now()

	query := ""
	args := []interface{}{}
	num := 1

	for i := 0; i < count; i++ {
		row := models.NewPerson("PgxPool3 AddMany #" + strconv.Itoa(i))

		query = fmt.Sprintf("%s($%d, $%d, $%d),", query, num, num+1, num+2)
		args = append(args, row.ID, row.Name, row.CreatedAt)
		num += 3
	}

	query = `INSERT INTO "TbPerson" ("id", "name", "created_at") VALUES ` + query
	query = query[0 : len(query)-1]

	tx, _ := db.Begin(ctx)
	defer tx.Rollback(ctx)

	_, err := tx.Exec(ctx, query, args...)
	if err != nil {
		slog.Error("Failed to execute batch insert", slog.Any("error", err))
	}

	tx.Commit(ctx)

	return time.Since(start)
}

func AddManyPgxPool4(ctx context.Context, count int) time.Duration {
	db := NewDatabasePgxPool(ctx)
	defer CloseDatabasePgxPool(ctx, db)

	start := time.Now()

	rows := [][]any{}
	for i := 0; i < count; i++ {
		row := models.NewPerson("PgxPool4 AddMany #" + strconv.Itoa(i))
		rows = append(rows, []any{row.ID, row.Name, row.CreatedAt})
	}

	tx, _ := db.Begin(ctx)
	defer tx.Rollback(ctx)

	_, err := tx.CopyFrom(
		ctx,
		pgx.Identifier{`TbPerson`},
		[]string{"id", "name", "created_at"},
		pgx.CopyFromRows(rows),
	)
	if err != nil {
		slog.Error("Failed to copy rows", slog.Any("error", err))
	}

	tx.Commit(ctx)

	return time.Since(start)
}
