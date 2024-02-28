package repositories_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/chrismarsilva/rinha-backend-2024/internals/databases"
	"github.com/chrismarsilva/rinha-backend-2024/internals/models"
	"github.com/chrismarsilva/rinha-backend-2024/internals/repositories"
	"github.com/chrismarsilva/rinha-backend-2024/internals/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

// go test
// go test -v
// go test -bench=.
// go test -bench=Add
// go test -run GetClientDbPadrao -v
// go test -run=GetClientDbPadrao -bench . -benchmem
// go test -run="BenchmarkClientRepoGet BenchmarkClientRepoGetWithouPrepare BenchmarkClientRepoGetWithPgx" -bench . -benchmem

//----------------------------------------------------------------

func GetClientDbPadrao() *pgxpool.Pool {
	cfg := utils.NewConfig()

	// driverDb := databases.DatabasePostgres{}
	// driverDb.StartDbConn(cfg)
	// db := driverDb.GetDatabaseConn()

	// driverDbWriter := databases.DatabasePostgres{}
	// driverDbWriter.StartDbWriter()
	// writer := driverDbWriter.GetDatabaseWriter()

	// driverDbReader := databases.DatabasePostgres{}
	// driverDbReader.StartDbReader()
	// reader := driverDbReader.GetDatabaseReader()

	driverDbPgx := databases.DatabasePostgres{}
	driverDbPgx.StartDbConnPgx(cfg)
	dbPgx := driverDbPgx.GetDatabaseConnPgx()

	return dbPgx
}

//----------------------------------------------------------------

func TestClientRepoGet(t *testing.T) {
	dbPgx := GetClientDbPadrao()
	repo := repositories.NewClientRepository(dbPgx)
	defer dbPgx.Close()

	var entity models.Cliente
	if err := repo.Get(&entity, 1); err != nil {
		t.Fatalf("error %t", err)
	}

	fmt.Println(entity)
}

// func TestClientRepoGetWithouPrepare(t *testing.T) {
// 	db, dbPgx := GetClientDbPadrao()
// 	repo := repositories.NewClientRepository(db, dbPgx)
// 	defer db.Close()
// 	defer dbPgx.Close()
//
// 	var entity models.Cliente
// 	if err := repo.GetWithouPrepare(&entity, 1); err != nil {
// 		t.Fatalf("error %t", err)
// 	}
//
// 	fmt.Println(entity)
// }

// func TestClientRepoGetWithPgx(t *testing.T) {
// 	db, dbPgx := GetClientDbPadrao()
// 	repo := repositories.NewClientRepository(db, dbPgx)
// 	defer db.Close()
// 	defer dbPgx.Close()
//
// 	var entity models.Cliente
// 	if err := repo.GetWithPgx(&entity, 1); err != nil {
// 		t.Fatalf("error %t", err)
// 	}
//
// 	fmt.Println(entity)
// }

//----------------------------------------------------------------

func TestClientRepoUpdateDebito(t *testing.T) {
	dbPgx := GetClientDbPadrao()
	repo := repositories.NewClientRepository(dbPgx)
	defer dbPgx.Close()

	tx, err := dbPgx.Begin(context.Background())
	if err != nil {
		t.Fatalf("error %t", err)
	}

	if err := repo.UpdSaldo(tx, 1, 100, "d"); err != nil {
		tx.Rollback(context.Background())
		t.Fatalf("error %t", err)
	}

	tx.Commit(context.Background())
}

func TestClientRepoUpdateCredito(t *testing.T) {
	dbPgx := GetClientDbPadrao()
	repo := repositories.NewClientRepository(dbPgx)
	defer dbPgx.Close()

	tx, err := dbPgx.Begin(context.Background())
	if err != nil {
		t.Fatalf("error %t", err)
	}

	if err := repo.UpdSaldo(tx, 1, 100, "c"); err != nil {
		tx.Rollback(context.Background())
		t.Fatalf("error %t", err)
	}
	tx.Commit(context.Background())
}

// func TestClientRepoUpdateDebitoWithouPrepare(t *testing.T) {
// 	db, dbPgx := GetClientDbPadrao()
// 	repo := repositories.NewClientRepository(db, dbPgx)
// 	defer db.Close()
// 	defer dbPgx.Close()
//
// 	tx := db.MustBegin()
//
// 	if err := repo.UpdSaldoWithouPrepare(tx, 1, 100, "d"); err != nil {
// 		tx.Rollback(context.Background())
// 		t.Fatalf("error %t", err)
// 	}
//
// 	tx.Commit(context.Background())
// }

// func TestClientRepoUpdateDebitoWithPgx(t *testing.T) {
// 	db, dbPgx := GetClientDbPadrao()
// 	repo := repositories.NewClientRepository(db, dbPgx)
// 	defer db.Close()
// 	defer dbPgx.Close()
//
// 	tx, err := dbPgx.Begin()
// 	if err != nil {
// 		t.Fatalf("error %t", err)
// 	}
//
// 	if err := repo.UpdSaldoWithPgx(tx, 1, 100, "d"); err != nil {
// 		tx.Rollback(context.Background())
// 		t.Fatalf("error %t", err)
// 	}
//
// 	tx.Commit(context.Background())
// }

//----------------------------------------------------------------

func BenchmarkClientRepoGet(b *testing.B) {
	dbPgx := GetClientDbPadrao()
	repo := repositories.NewClientRepository(dbPgx)
	defer dbPgx.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var entity models.Cliente
		if err := repo.Get(&entity, 1); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

// func BenchmarkClientRepoGetWithouPrepare(b *testing.B) {
// 	db, dbPgx := GetClientDbPadrao()
// 	repo := repositories.NewClientRepository(db, dbPgx)
// 	defer db.Close()
// 	defer dbPgx.Close()
//
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		var entity models.Cliente
// 		if err := repo.GetWithouPrepare(&entity, 1); err != nil {
// 			b.Fatalf("error %t", err)
// 		}
// 	}
// }

// func BenchmarkClientRepoGetWithPgx(b *testing.B) {
// 	db, dbPgx := GetClientDbPadrao()
// 	repo := repositories.NewClientRepository(db, dbPgx)
// 	defer db.Close()
// 	defer dbPgx.Close()
//
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		var entity models.Cliente
// 		if err := repo.GetWithPgx(&entity, 1); err != nil {
// 			b.Fatalf("error %t", err)
// 		}
// 	}
// }

//----------------------------------------------------------------

func BenchmarkClientRepoUpdate(b *testing.B) {
	dbPgx := GetClientDbPadrao()
	repo := repositories.NewClientRepository(dbPgx)
	defer dbPgx.Close()

	b.ResetTimer()

	tx, err := dbPgx.Begin(context.Background())
	if err != nil {
		b.Fatalf("error %t", err)
	}

	for i := 0; i < b.N; i++ {
		if err := repo.UpdSaldo(tx, 1, 100, "c"); err != nil {
			tx.Rollback(context.Background())
			b.Fatalf("error %t", err)
		}
	}

	tx.Commit(context.Background())
}

// func BenchmarkClientRepoUpdateWithouPrepare(b *testing.B) {
// 	db, dbPgx := GetClientDbPadrao()
// 	repo := repositories.NewClientRepository(db, dbPgx)
// 	defer db.Close()
// 	defer dbPgx.Close()
//
// 	b.ResetTimer()
//
// 	tx := db.MustBegin()
//
// 	for i := 0; i < b.N; i++ {
// 		if err := repo.UpdSaldoWithouPrepare(tx, 1, 100, "c"); err != nil {
// 			tx.Rollback(context.Background())
// 			b.Fatalf("error %t", err)
// 		}
// 	}
//
// 	tx.Commit(context.Background())
// }

// func BenchmarkClientRepoUpdateWithPgx(b *testing.B) {
// 	db, dbPgx := GetClientDbPadrao()
// 	repo := repositories.NewClientRepository(db, dbPgx)
// 	defer db.Close()
// 	defer dbPgx.Close()
//
// 	b.ResetTimer()
//
// 	tx, err := db.Begin()
// 	if err != nil {
// 		b.Fatalf("error %t", err)
// 	}
//
// 	// defer func() {
// 	//     if err != nil {
// 	//         tx.Rollback(context.Background())
// 	//     } else {
// 	//         tx.Commit(context.Background())
// 	//     }
// 	// }()
//
// 	for i := 0; i < b.N; i++ {
// 		if err := repo.UpdSaldoWithPgx(tx, 1, 100, "c"); err != nil {
// 			tx.Rollback(context.Background())
// 			b.Fatalf("error %t", err)
// 		}
// 	}
//
// 	tx.Commit(context.Background())
// }

//----------------------------------------------------------------

func BenchmarkClientRepoGetAndUpdate(b *testing.B) {
	dbPgx := GetClientDbPadrao()
	repo := repositories.NewClientRepository(dbPgx)
	defer dbPgx.Close()

	b.ResetTimer()

	tx, err := dbPgx.Begin(context.Background())
	if err != nil {
		b.Fatalf("error %t", err)
	}

	for i := 0; i < b.N; i++ {
		var entity models.Cliente
		if err := repo.Get(&entity, 1); err != nil {
			tx.Rollback(context.Background())
			b.Fatalf("error %t", err)
		}
		if err := repo.UpdSaldo(tx, 1, 100, "c"); err != nil {
			tx.Rollback(context.Background())
			b.Fatalf("error %t", err)
		}
	}

	tx.Commit(context.Background())
}

//----------------------------------------------------------------

func BenchmarkClientRepoGet10Mil(b *testing.B) {
	dbPgx := GetClientDbPadrao()
	repo := repositories.NewClientRepository(dbPgx)
	defer dbPgx.Close()

	b.ResetTimer()
	for i := 0; i < 10_001; i++ {
		var entity models.Cliente
		if err := repo.Get(&entity, 1); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkClientRepoUpdate10Mil(b *testing.B) {
	dbPgx := GetClientDbPadrao()
	repo := repositories.NewClientRepository(dbPgx)
	defer dbPgx.Close()

	b.ResetTimer()

	tx, err := dbPgx.Begin(context.Background())
	if err != nil {
		b.Fatalf("error %t", err)
	}

	for i := 0; i < 10_001; i++ {
		if err := repo.UpdSaldo(tx, 1, 100, "c"); err != nil {
			tx.Rollback(context.Background())
			b.Fatalf("error %t", err)
		}
	}

	tx.Commit(context.Background())
}

func BenchmarkClientRepoGetAndUpdate10Mil(b *testing.B) {
	dbPgx := GetClientDbPadrao()
	repo := repositories.NewClientRepository(dbPgx)
	defer dbPgx.Close()

	b.ResetTimer()
	
	tx, err := dbPgx.Begin(context.Background())
	if err != nil {
		b.Fatalf("error %t", err)
	}

	for i := 0; i < 10_001; i++ {
		var entity models.Cliente

		b.StopTimer()
		err := repo.Get(&entity, 1)
		if err != nil {
			tx.Rollback(context.Background())
			b.Fatalf("error %t", err)
		}
		b.StartTimer()

		b.StopTimer()
		err = repo.UpdSaldo(tx, 1, 100, "c")
		if err != nil {
			tx.Rollback(context.Background())
			b.Fatalf("error %t", err)
		}
		b.StartTimer()
	}

	tx.Commit(context.Background())
}

//----------------------------------------------------------------
