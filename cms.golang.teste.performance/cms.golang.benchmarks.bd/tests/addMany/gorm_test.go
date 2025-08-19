package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/models"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/stores"
)

const (
	GORM_PREFIX = "Gorm AddMany"
	GORM_COUNT  = 1_000 // 1 / 10 / 100 / 1_000 / 10_000 / 100_000 / 1_000_000
)

func init() {
	stores.ClearGorm(context.Background())

	// rows := models.GenerateFakePersons(GORM_PREFIX, GORM_COUNT)
	// for _, row := range rows {
	// 	db.Create(&row)
	// }
}

func BenchmarkGormAddMany1(b *testing.B) {
	ctx := context.Background()
	db := stores.NewDatabaseGorm(ctx)
	defer stores.CloseDatabaseGorm(ctx, db)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rows := models.GenerateFakePersons(GORM_PREFIX+"1", GORM_COUNT)

		//tx := db.Begin()
		//tx := db.Session(&gorm.Session{PrepareStmt: true, CreateBatchSize: GORM_COUNT})
		tx := db.Begin()
		if tx.Error != nil {
			b.Fatalf("failed to begin tx: %v", tx.Error)
		}

		for _, row := range rows {
			err := tx.Create(&row).Error
			if err != nil {
				b.Error("failed to insert row:", err)
				_ = tx.Rollback()
				break
			}
		}

		err := tx.Commit().Error
		if err != nil {
			// b.Error("failed to commit tx:", err)
			_ = tx.Rollback()
		}
	}
}

func BenchmarkGormAddMany2(b *testing.B) {
	ctx := context.Background()
	db := stores.NewDatabaseGorm(ctx)
	defer stores.CloseDatabaseGorm(ctx, db)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rows := models.GenerateFakePersons(GORM_PREFIX+"2", GORM_COUNT)

		tx := db.Begin()
		if tx.Error != nil {
			b.Fatalf("failed to begin tx: %v", tx.Error)
		}

		err := tx.Create(&rows).Error
		if err != nil {
			b.Error("failed to insert rows:", err)
			_ = tx.Rollback()
			break
		}

		err = tx.Commit().Error
		if err != nil {
			// b.Error("failed to commit tx:", err)
			_ = tx.Rollback()
		}
	}
}

func BenchmarkGormAddMany3(b *testing.B) {
	ctx := context.Background()
	db := stores.NewDatabaseGorm(ctx)
	defer stores.CloseDatabaseGorm(ctx, db)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		rows := models.GenerateFakePersons(GORM_PREFIX+"3", GORM_COUNT)

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

		tx := db.Begin()
		if tx.Error != nil {
			b.Fatalf("failed to begin tx: %v", tx.Error)
		}

		err := tx.Exec(query, args...).Error
		if err != nil {
			b.Fatalf("failed to exec query: %v", err)
		}

		err = tx.Commit().Error
		if err != nil {
			// b.Error("failed to commit tx:", err)
			_ = tx.Rollback()
		}
	}
}
