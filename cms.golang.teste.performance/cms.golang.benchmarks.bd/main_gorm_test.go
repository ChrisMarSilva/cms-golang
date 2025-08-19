package main

import (
	"context"
	"testing"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/models"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/stores"
)

func init() {
	stores.ClearGorm(context.Background())
	// for _, row := range models.GenerateFakePersons(100) {
	// 	db.Create(&row)
	// }
}

func BenchmarkInsertGorm(b *testing.B) {
	ctx := context.Background()
	db := stores.NewDatabaseGorm(ctx)
	defer stores.CloseDatabaseGorm(ctx, db)

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		//b.StopTimer()
		row := models.NewPerson("Gorm - User")
		//b.StartTimer()

		db.Create(&row)
	}
}

// func BenchmarkPgxPoolInsertMulti(b *testing.B) { }
// func BenchmarkPgxPoolUpdate(b *testing.B) { }
// func BenchmarkPgxPoolGetOne(b *testing.B) { }
// func BenchmarkPgxPoolGetAll(b *testing.B) { }
