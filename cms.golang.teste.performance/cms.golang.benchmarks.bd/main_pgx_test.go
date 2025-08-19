package main

// go test -bench . -benchmem

// func setup() { }
// func cleanup() { }

// b.ReportAllocs()
// b.ResetTimer()
// b.StopTimer()
// 	if err != nil { b.Fatal(err) }
// 	b.StartTimer()
// for b.Loop() {
// for i := 0; i < b.N; i++ {

import (
	"context"
	"testing"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/models"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/stores"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/utils"
)

func init() {
	stores.ClearPgxPool(context.Background())
	// for _, row := range models.GenerateFakePersons(100) {
	// 	db.Exec(ctx, utils.DBInsert, row.ID, row.Name, row.CreatedAt)
	// }
}

func BenchmarkInsertPgxPool(b *testing.B) {
	ctx := context.Background()
	db := stores.NewDatabasePgxPool(ctx)
	defer stores.CloseDatabasePgxPool(ctx, db)

	b.ReportAllocs()
	b.ResetTimer()
	for b.Loop() {
		//b.StopTimer()
		row := models.NewPerson("PgxPool - User")
		//b.StartTimer()

		db.Exec(ctx, utils.DBInsert, row.ID, row.Name, row.CreatedAt)
	}
}

// func BenchmarkPgxPoolInsertMulti(b *testing.B) {
// 	ctx := context.Background()
// 	db := stores.NewDatabasePgxPool(ctx)
// 	defer db.Close()

// 	db.Exec(ctx, utils.DBDeleteAll)

// 	rows := models.GenerateFakePersons(100)

// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		for _, row := range rows {
// 			_, err := db.Exec(ctx, utils.DBInsert, row.ID, row.Name, row.CreatedAt)

// 			b.StopTimer()
// 			if err != nil {
// 				b.Fatal(err)
// 			}
// 			b.StartTimer()
// 		}
// 	}
// }

// func BenchmarkPgxPoolUpdate(b *testing.B) {
// 	ctx := context.Background()
// 	db := stores.NewDatabasePgxPool(ctx)
// 	defer db.Close()

// 	db.Exec(ctx, utils.DBDeleteAll)

// 	row := models.GenerateFakePerson()
// 	_, err := db.Exec(ctx, utils.DBInsert, row.ID, row.Name, row.CreatedAt)
// 	if err != nil {
// 		b.Fatal(err)
// 	}

// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		_, err := db.Exec(ctx, utils.DBUpdate, row.ID, row.Name, row.CreatedAt)

// 		b.StopTimer()
// 		if err != nil {
// 			b.Fatal(err)
// 		}
// 		b.StartTimer()
// 	}
// }

// func BenchmarkPgxPoolGetOne(b *testing.B) {
// 	ctx := context.Background()
// 	db := stores.NewDatabasePgxPool(ctx)
// 	defer db.Close()

// 	db.Exec(ctx, utils.DBDeleteAll)

// 	row := models.GenerateFakePerson()
// 	_, err := db.Exec(ctx, utils.DBInsert, row.ID, row.Name, row.CreatedAt)
// 	if err != nil {
// 		b.Fatal(err)
// 	}

// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		var model models.PersonModel
// 		err := db.QueryRow(ctx, utils.DBSelectOne, 1).Scan(&model.ID, &model.Name, &model.CreatedAt)

// 		b.StopTimer()
// 		if err != nil {
// 			b.Fatal(err)
// 		}
// 		b.StartTimer()
// 	}
// }

// func BenchmarkPgxPoolGetAll(b *testing.B) {
// 	ctx := context.Background()
// 	db := stores.NewDatabasePgxPool(ctx)
// 	defer db.Close()

// 	db.Exec(ctx, utils.DBDeleteAll)

// 	rows := models.GenerateFakePersons(100)
// 	for _, row := range rows {
// 		_, err := db.Exec(ctx, utils.DBInsert, row.ID, row.Name, row.CreatedAt)
// 		if err != nil {
// 			b.Fatal(err)
// 		}
// 	}

// 	b.ReportAllocs()
// 	b.ResetTimer()

// 	for i := 0; i < b.N; i++ {
// 		model := make([]models.PersonModel, len(rows))
// 		rows, err := db.Query(ctx, utils.DBSelectAll)
// 		if err != nil {
// 			b.Fatal(err)
// 		}

// 		for j := 0; rows.Next() && j < len(model); j++ {
// 			err = rows.Scan(&model[j].ID, &model[j].Name, &model[j].CreatedAt)
// 			if err != nil {
// 				b.Fatal(err)
// 			}
// 			rows.Close()
// 		}
// 	}
// }
