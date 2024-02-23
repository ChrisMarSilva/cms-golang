package repositories_test

import (
	"testing"

	"github.com/chrismarsilva/rinha-backend-2024/internals/databases"
	"github.com/chrismarsilva/rinha-backend-2024/internals/models"
	"github.com/chrismarsilva/rinha-backend-2024/internals/repositories"
	"github.com/chrismarsilva/rinha-backend-2024/internals/utils"
	"github.com/jmoiron/sqlx"
)

// go test
// go test -v
// go test -bench=.
// go test -bench=Add
// go test -run TestRepoGet -v
// go test -run=XXX -bench . -benchmem

func GetClientTransactionDbPadrao() *sqlx.DB {
	cfg := utils.NewConfig()

	driverDb := databases.DatabasePostgres{}
	driverDb.StartDbConn(cfg)
	db := driverDb.GetDatabaseConn()

	// driverDbWriter := databases.DatabasePostgres{}
	// driverDbWriter.StartDbWriter()
	// writer := driverDbWriter.GetDatabaseWriter()

	// driverDbReader := databases.DatabasePostgres{}
	// driverDbReader.StartDbReader()
	// reader := driverDbReader.GetDatabaseReader()

	return db
}

func TestClientTransactionRepoGet(t *testing.T) {
	db := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(db)
	defer db.Close()

	// for _, entity := range entities {
	// 	fmt.Println(entity)
	// }

	entities := map[int]models.ClienteTransacao{}
	if err := repo.GetAll(&entities, 1); err != nil {
		t.Fatalf("error %t", err)
	}
}

func TestClientTransactionRepoCreateDebito(t *testing.T) {
	db := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(db)
	defer db.Close()

	tx := db.MustBegin()

	if err := repo.Add(tx, 1, 100, "d", "debito"); err != nil {
		tx.Rollback()
		t.Fatalf("error %t", err)
	}

	tx.Commit()
}

func TestClientTransactionRepoCreateCredito(t *testing.T) {
	db := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(db)
	defer db.Close()

	tx := db.MustBegin()

	if err := repo.Add(tx, 1, 100, "c", "credito"); err != nil {
		tx.Rollback()
		t.Fatalf("error %t", err)
	}

	tx.Commit()
}

func BenchmarkClientTransactionRepoGet(b *testing.B) {
	db := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(db)
	defer db.Close()

	for i := 0; i < b.N; i++ {
		entities := map[int]models.ClienteTransacao{}
		if err := repo.GetAll(&entities, 1); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkClientTransactionRepoCreate(b *testing.B) {
	db := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(db)
	defer db.Close()

	tx := db.MustBegin()

	for i := 0; i < b.N; i++ {
		if err := repo.Add(tx, 1, 100, "c", "credito"); err != nil {
			tx.Rollback()
			b.Fatalf("error %t", err)
		}
	}

	tx.Commit()
}

func BenchmarkClientTransactionRepoGet10Mil(b *testing.B) {
	db := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(db)
	defer db.Close()

	for i := 0; i < 10_001; i++ {
		entities := map[int]models.ClienteTransacao{}
		if err := repo.GetAll(&entities, 1); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkClientTransactionRepoCreate10Mil(b *testing.B) {
	db := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(db)
	defer db.Close()

	tx := db.MustBegin()

	for i := 0; i < 10_001; i++ {
		if err := repo.Add(tx, 1, 100, "c", "credito"); err != nil {
			tx.Rollback()
			b.Fatalf("error %t", err)
		}
	}

	tx.Commit()
}
