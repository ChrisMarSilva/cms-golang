package repositories_test

import (
	"fmt"
	"testing"

	"github.com/chrismarsilva/rinha-backend-2024/internals/databases"
	"github.com/chrismarsilva/rinha-backend-2024/internals/models"
	"github.com/chrismarsilva/rinha-backend-2024/internals/repositories"
	"github.com/jmoiron/sqlx"
)

// go test
// go test -v
// go test -bench=.
// go test -bench=Add
// go test -run GetClientDbPadrao -v
// go test -run=GetClientDbPadrao -bench . -benchmem

func GetClientDbPadrao() (*sqlx.DB, *sqlx.DB) {
	driverDbWriter := databases.DatabasePostgres{}
	driverDbWriter.StartDbWriter()
	writer := driverDbWriter.GetDatabaseWriter()

	driverDbReader := databases.DatabasePostgres{}
	driverDbReader.StartDbReader()
	reader := driverDbReader.GetDatabaseReader()

	return writer, reader
}

func TestClientRepoGet(t *testing.T) {
	writer, reader := GetClientDbPadrao()
	repo := repositories.NewClientRepository(writer, reader)
	defer writer.Close()
	defer reader.Close()

	var entity models.Cliente
	if err := repo.Get(&entity, 1); err != nil {
		t.Fatalf("error %t", err)
	}

	fmt.Println(entity)
}

func TestClientRepoUpdateDebito(t *testing.T) {
	writer, reader := GetClientDbPadrao()
	repo := repositories.NewClientRepository(writer, reader)

	tx := writer.MustBegin()

	if err := repo.UpdSaldo(tx, 1, 100, "d"); err != nil {
		tx.Rollback()
		t.Fatalf("error %t", err)
	}

	tx.Commit()
}

func TestClientRepoUpdateCredito(t *testing.T) {
	writer, reader := GetClientDbPadrao()
	repo := repositories.NewClientRepository(writer, reader)

	tx := writer.MustBegin()

	if err := repo.UpdSaldo(tx, 1, 100, "c"); err != nil {
		tx.Rollback()
		t.Fatalf("error %t", err)
	}
	tx.Commit()
}

func BenchmarkClientRepoGet(b *testing.B) {
	writer, reader := GetClientDbPadrao()
	repo := repositories.NewClientRepository(writer, reader)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var entity models.Cliente
		if err := repo.Get(&entity, 1); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkClientRepoUpdate(b *testing.B) {
	writer, reader := GetClientDbPadrao()
	repo := repositories.NewClientRepository(writer, reader)

	b.ResetTimer()

	tx := writer.MustBegin()

	for i := 0; i < b.N; i++ {
		if err := repo.UpdSaldo(tx, 1, 100, "c"); err != nil {
			tx.Rollback()
			b.Fatalf("error %t", err)
		}
	}

	tx.Commit()
}

func BenchmarkClientRepoGetAndUpdate(b *testing.B) {
	writer, reader := GetClientDbPadrao()
	repo := repositories.NewClientRepository(writer, reader)

	b.ResetTimer()

	tx := writer.MustBegin()

	for i := 0; i < b.N; i++ {
		var entity models.Cliente
		if err := repo.Get(&entity, 1); err != nil {
			tx.Rollback()
			b.Fatalf("error %t", err)
		}
		if err := repo.UpdSaldo(tx, 1, 100, "c"); err != nil {
			tx.Rollback()
			b.Fatalf("error %t", err)
		}
	}

	tx.Commit()
}

func BenchmarkClientRepoGet10Mil(b *testing.B) {
	writer, reader := GetClientDbPadrao()
	repo := repositories.NewClientRepository(writer, reader)

	b.ResetTimer()
	for i := 0; i < 10_001; i++ {
		var entity models.Cliente
		if err := repo.Get(&entity, 1); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkClientRepoUpdate10Mil(b *testing.B) {
	writer, reader := GetClientDbPadrao()
	repo := repositories.NewClientRepository(writer, reader)

	b.ResetTimer()

	tx := writer.MustBegin()

	for i := 0; i < 10_001; i++ {
		if err := repo.UpdSaldo(tx, 1, 100, "c"); err != nil {
			tx.Rollback()
			b.Fatalf("error %t", err)
		}
	}

	tx.Commit()
}

func BenchmarkClientRepoGetAndUpdate10Mil(b *testing.B) {
	writer, reader := GetClientDbPadrao()
	repo := repositories.NewClientRepository(writer, reader)

	b.ResetTimer()

	tx := writer.MustBegin()

	for i := 0; i < 10_001; i++ {
		var entity models.Cliente

		b.StopTimer()
		err := repo.Get(&entity, 1)
		if err != nil {
			tx.Rollback()
			b.Fatalf("error %t", err)
		}
		b.StartTimer()

		b.StopTimer()
		err = repo.UpdSaldo(tx, 1, 100, "c")
		if err != nil {
			tx.Rollback()
			b.Fatalf("error %t", err)
		}
		b.StartTimer()
	}

	tx.Commit()
}
