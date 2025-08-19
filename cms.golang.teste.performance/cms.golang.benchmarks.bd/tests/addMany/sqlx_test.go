package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/models"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/stores"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/utils"
)

const (
	SLQX_PREFIX = "SqlX AddMany"
	SLQX_COUNT  = 1_000 // 1 / 10 / 100 / 1_000 / 10_000 / 100_000 / 1_000_000
)

func init() {
	stores.ClearSqlX(context.Background())

	// rows := models.GenerateFakePersons(SLQX_PREFIX, SLQX_COUNT)
	// for _, row := range rows {
	//  db.MustExec(utils.DBInsert, row.ID, row.Name, row.CreatedAt)
	// }
}

func BenchmarkSqlXAddMany1(b *testing.B) {
	ctx := context.Background()
	db := stores.NewDatabaseSqlX(ctx)
	defer stores.CloseDatabaseSqlX(ctx, db)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rows := models.GenerateFakePersons(SLQX_PREFIX+"1", SLQX_COUNT)

		tx, err := db.Beginx()
		if err != nil {
			b.Fatalf("failed to begin tx: %v", err)
		}

		for _, row := range rows {
			_, err := tx.Exec(utils.DBInsert, row.ID, row.Name, row.CreatedAt)
			if err != nil {
				b.Error("failed to insert row:", err)
				_ = tx.Rollback()
				break
			}
		}

		err = tx.Commit()
		if err != nil {
			// b.Error("failed to commit tx:", err)
			_ = tx.Rollback()
		}
	}
}

func BenchmarkSqlXAddMany2(b *testing.B) {
	ctx := context.Background()
	db := stores.NewDatabaseSqlX(ctx)
	defer stores.CloseDatabaseSqlX(ctx, db)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rows := models.GenerateFakePersons(SLQX_PREFIX+"2", SLQX_COUNT)

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

		tx, err := db.Beginx()
		if err != nil {
			b.Fatalf("failed to begin tx: %v", err)
		}

		_, err = tx.Exec(query, args...)
		if err != nil {
			_ = tx.Rollback()
			b.Fatalf("failed to execute statement: %v", err)
		}

		err = tx.Commit()
		if err != nil {
			// b.Error("failed to commit tx:", err)
			_ = tx.Rollback()
		}
	}
}
