package tests

import (
	"context"
	"testing"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/config_sqlc/tutorial"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/models"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/stores"
)

const (
	SLQC_PREFIX = "SqlC AddMany"
	SLQC_COUNT  = 1_000 // 1 / 10 / 100 / 1_000 / 10_000 / 100_000 / 1_000_000
)

func init() {
	// ctx := context.Background()
	// db := stores.NewDatabaseSqlC(ctx)
	// defer stores.CloseDatabaseSqlC(ctx, db)
	// queries := tutorial.New(db)

	//queries.Clear(ctx)
	stores.ClearSqlC(context.Background())

	// rows := models.GenerateFakePersons(SLQC_PREFIX, SLQC_COUNT)
	// for _, row := range rows {
	//  queries.Create(ctx, tutorial.CreateParams{ID: row.ID, Name: row.Name, CreatedAt: row.CreatedAt})
	// }
}

func BenchmarkSqlCAddMany1(b *testing.B) {
	ctx := context.Background()
	db := stores.NewDatabaseSqlC(ctx)
	defer stores.CloseDatabaseSqlC(ctx, db)
	queries := tutorial.New(db)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rows := models.GenerateFakePersons(SLQC_PREFIX+"1", SLQC_COUNT)

		for _, row := range rows {
			queries.Create(ctx, tutorial.CreateParams{ID: row.ID, Name: row.Name, CreatedAt: row.CreatedAt})
		}
	}
}
