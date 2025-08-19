package stores

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/models"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/utils"
	"github.com/jmoiron/sqlx"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDatabaseSqlX(ctx context.Context) *sqlx.DB {
	db, err := sqlx.Connect(utils.DBDriverPgx, utils.DBUri)
	if err != nil {
		slog.Error("Failed to connect to database", slog.Any("error", err))
		return nil
	}

	return db
}

func CloseDatabaseSqlX(ctx context.Context, db *sqlx.DB) {
	if db != nil {
		db.Close()
	}
}

func GetAllSqlX(ctx context.Context, db *sqlx.DB, models *[]models.PersonModel) error {
	return db.Select(&models, utils.DBQuery)
}

func FetchAllSqlX(ctx context.Context) {

	slog.Info("Connecting to database using SQLX...")

	db := NewDatabaseSqlX(ctx)
	defer CloseDatabaseSqlX(ctx, db)

	var persons []models.PersonModel
	err := db.Select(&persons, utils.DBQuery)
	if err != nil {
		slog.Error("Failed to query persons from database", slog.Any("error", err))
		return
	}

	for _, p := range persons {
		slog.Info("PersonModel", "person", p.Name)
	}

	slog.Info("Persons retrieved from database", slog.Int("count", len(persons)))
}

func ClearSqlX(ctx context.Context) {
	db := NewDatabaseSqlX(ctx)
	defer CloseDatabaseSqlX(ctx, db)

	//db.Exec(utils.DBDeleteSqlX)
	db.MustExec(utils.DBDeleteSqlX)
}

func AddOneSqlX(ctx context.Context, count int) time.Duration {
	db := NewDatabaseSqlX(ctx)
	defer CloseDatabaseSqlX(ctx, db)

	start := time.Now()
	for i := 0; i < count; i++ {
		row := models.NewPerson("SqlX AddOne #" + strconv.Itoa(i))
		db.MustExec(utils.DBInsert, row.ID, row.Name, row.CreatedAt)
	}

	return time.Since(start)
}

func AddManySqlX(ctx context.Context, count int) time.Duration {
	db := NewDatabaseSqlX(ctx)
	defer CloseDatabaseSqlX(ctx, db)

	start := time.Now()

	rows := make([]*models.PersonModel, 0, count)
	for i := 0; i < count; i++ {
		row := models.NewPerson("SqlX1 AddMany #" + strconv.Itoa(i))
		rows = append(rows, row)
	}

	tx := db.MustBegin() // Beginx
	defer tx.Rollback()

	for i := 0; i < count; i++ {
		tx.MustExec(utils.DBInsert, rows[i].ID, rows[i].Name, rows[i].CreatedAt)
	}

	tx.Commit()

	return time.Since(start)
}

func AddManySqlX2(ctx context.Context, count int) time.Duration {
	db := NewDatabaseSqlX(ctx)
	defer CloseDatabaseSqlX(ctx, db)

	start := time.Now()
	tx := db.MustBegin() // Beginx
	defer tx.Rollback()

	query := ""
	args := []interface{}{}
	num := 1

	for i := 0; i < count; i++ {
		row := models.NewPerson("SqlX2 AddMany #" + strconv.Itoa(i))

		query = fmt.Sprintf("%s($%d, $%d, $%d),", query, num, num+1, num+2)
		args = append(args, row.ID, row.Name, row.CreatedAt)
		num += 3
	}

	query = `INSERT INTO "TbPerson" ("id", "name", "created_at") VALUES ` + query
	query = query[0 : len(query)-1]

	stmt, _ := tx.Prepare(query)
	defer stmt.Close()

	//_, err := tx.Exec(query, args...)
	stmt.Exec(args...)
	// if err != nil {
	// 	slog.Error("Failed to execute batch insert", slog.Any("error", err))
	// }

	tx.Commit()

	return time.Since(start)
}
