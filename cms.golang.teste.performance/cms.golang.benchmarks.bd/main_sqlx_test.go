package main

import (
	"context"
	"testing"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/models"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/stores"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/utils"
)

func init() {
	stores.ClearSqlX(context.Background())
	// for _, row := range models.GenerateFakePersons(100 {
	// 	db.MustExec(utils.DBInsert, row.ID, row.Name, row.CreatedAt)
	// }
}

func BenchmarkInsertSqlX(b *testing.B) {
	ctx := context.Background()
	db := stores.NewDatabaseSqlX(ctx)
	defer stores.CloseDatabaseSqlX(ctx, db)

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		//b.StopTimer()
		row := models.NewPerson("SqlX - User")
		//b.StartTimer()

		db.MustExec(utils.DBInsert, row.ID, row.Name, row.CreatedAt)
	}
}

// func BenchmarkPgxPoolInsertMulti(b *testing.B) { }
// func BenchmarkPgxPoolUpdate(b *testing.B) { }
// func BenchmarkPgxPoolGetOne(b *testing.B) { }
// func BenchmarkPgxPoolGetAll(b *testing.B) { }
