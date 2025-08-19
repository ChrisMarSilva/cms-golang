package main

import (
	"context"
	"testing"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/config_sqlc/tutorial"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/models"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/stores"
)

func init() {
	ctx := context.Background()
	db := stores.NewDatabaseSqlC(ctx)
	defer stores.CloseDatabaseSqlC(ctx, db)
	queries := tutorial.New(db)

	queries.Clear(ctx)
	// for _, row := range models.GenerateFakePersons(100 {
	// 	queries.Create(ctx, tutorial.CreateParams{ID: row.ID, Name: row.Name, CreatedAt: row.CreatedAt})
	// }
}

func BenchmarkInsertSqlC(b *testing.B) {
	ctx := context.Background()
	db := stores.NewDatabaseSqlC(ctx)
	defer stores.CloseDatabaseSqlC(ctx, db)
	queries := tutorial.New(db)

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		//b.StopTimer()
		row := models.NewPerson("SqlC - User")
		arg := tutorial.CreateParams{ID: row.ID, Name: row.Name, CreatedAt: row.CreatedAt}
		//b.StartTimer()

		queries.Create(ctx, arg)
	}
}

// func BenchmarkPgxPoolInsertMulti(b *testing.B) { }
// func BenchmarkPgxPoolUpdate(b *testing.B) { }
// func BenchmarkPgxPoolGetOne(b *testing.B) { }
// func BenchmarkPgxPoolGetAll(b *testing.B) { }
