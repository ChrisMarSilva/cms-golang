package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/models"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/stores"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/utils"
	"github.com/jackc/pgx/v5"
)

const (
	PGX_PREFIX = "PgxPool AddMany"
	PGX_COUNT  = 1_000 // 1 / 10 / 100 / 1_000 / 10_000 / 100_000 / 1_000_000
)

func init() {
	stores.ClearPgxPool(context.Background())

	// rows := models.GenerateFakePersons(PGX_PREFIX, PGX_COUNT)
	// for _, row := range rows {
	// 	db.Exec(ctx, utils.DBInsert, row.ID, row.Name, row.CreatedAt)
	// }
}

func BenchmarkPgxPoolAddMany1(b *testing.B) {
	ctx := context.Background()
	db := stores.NewDatabasePgxPool(ctx)
	defer stores.CloseDatabasePgxPool(ctx, db)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rows := models.GenerateFakePersons(PGX_PREFIX+"1", PGX_COUNT)

		tx, err := db.Begin(ctx)
		if err != nil {
			b.Fatalf("failed to begin tx: %v", err)
		}

		//ok := false
		for _, row := range rows {
			_, err := tx.Exec(ctx, utils.DBInsert, row.ID, row.Name, row.CreatedAt)
			if err != nil {
				b.Error("failed to insert row:", err)
				_ = tx.Rollback(ctx)
				//ok = false
				break
			}
			//ok = true
		}

		//if ok {
		err = tx.Commit(ctx)
		if err != nil {
			// b.Error("failed to commit tx:", err)
			_ = tx.Rollback(ctx)
		}
		//}
	}
}

func BenchmarkPgxPoolAddMany2(b *testing.B) {
	ctx := context.Background()
	db := stores.NewDatabasePgxPool(ctx)
	defer stores.CloseDatabasePgxPool(ctx, db)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rows := models.GenerateFakePersons(PGX_PREFIX+"2", PGX_COUNT)

		batch := &pgx.Batch{}
		for _, row := range rows {
			batch.Queue(utils.DBInsert, row.ID, row.Name, row.CreatedAt)
		}

		results := db.SendBatch(ctx, batch)

		for i := 0; i < PGX_COUNT; i++ {
			_, err := results.Exec()
			if err != nil {
				results.Close()
				b.Fatalf("failed to execute batch: %v", err)
			}
		}

		results.Close()
	}
}

func BenchmarkPgxPoolAddMany3(b *testing.B) {
	ctx := context.Background()
	db := stores.NewDatabasePgxPool(ctx)
	defer stores.CloseDatabasePgxPool(ctx, db)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rows := models.GenerateFakePersons(PGX_PREFIX+"3", PGX_COUNT)

		query := ""
		args := []interface{}{}
		num := 1

		for _, row := range rows {
			query = fmt.Sprintf("%s($%d, $%d, $%d),", query, num, num+1, num+2)
			args = append(args, row.ID, row.Name, row.CreatedAt)
			num += 3
		}

		query = `INSERT INTO "TbPerson" ("id", "name", "created_at") VALUES ` + query
		query = query[0 : len(query)-1]

		tx, err := db.Begin(ctx)
		if err != nil {
			b.Fatalf("failed to begin tx: %v", err)
		}

		_, err = tx.Exec(ctx, query, args...)
		if err != nil {
			b.Error("failed to insert row:", err)
			_ = tx.Rollback(ctx)
			break
		}

		err = tx.Commit(ctx)
		if err != nil {
			// b.Error("failed to commit tx:", err)
			_ = tx.Rollback(ctx)
		}
	}
}

func BenchmarkPgxPoolAddMany4(b *testing.B) {
	ctx := context.Background()
	db := stores.NewDatabasePgxPool(ctx)
	defer stores.CloseDatabasePgxPool(ctx, db)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rows := models.GenerateFakePersons(PGX_PREFIX+"4", PGX_COUNT)

		persons := [][]any{}
		for _, row := range rows {
			persons = append(persons, []any{row.ID, row.Name, row.CreatedAt})
		}

		tx, err := db.Begin(ctx)
		if err != nil {
			b.Fatalf("failed to begin tx: %v", err)
		}

		tableName := pgx.Identifier{`TbPerson`}
		columnNames := []string{"id", "name", "created_at"}
		rowSrc := pgx.CopyFromRows(persons)

		_, err = tx.CopyFrom(ctx, tableName, columnNames, rowSrc)
		if err != nil {
			b.Error("failed to copy rows:", err)
			_ = tx.Rollback(ctx)
			break
		}

		err = tx.Commit(ctx)
		if err != nil {
			// b.Error("failed to commit tx:", err)
			_ = tx.Rollback(ctx)
		}
	}
}
