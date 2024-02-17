package repositories_test

import (
	"fmt"
	"testing"

	"github.com/chrismarsilva/rinha-backend-2024/internals/databases"
	"github.com/chrismarsilva/rinha-backend-2024/internals/models"
	"github.com/chrismarsilva/rinha-backend-2024/internals/repositories"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

// go test
// go test -v
// go test -bench=.
// go test -bench=Add
// go test -run TestRepoGet -v
// go test -run=XXX -bench . -benchmem

func GetClientTransactionDbPadrao() (*sqlx.DB, *sqlx.DB) {
	viper.AddConfigPath("./")
	viper.SetConfigFile("../../cmd/api-server/.env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	driverDbWriter := databases.DatabasePostgres{}
	driverDbWriter.StartDbWriter()
	writer := driverDbWriter.GetDatabaseWriter()

	driverDbReader := databases.DatabasePostgres{}
	driverDbReader.StartDbReader()
	reader := driverDbReader.GetDatabaseReader()

	return writer, reader
}

func TestClientTransactionRepoGet(t *testing.T) {
	writer, reader := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(writer, reader)
	defer writer.Close()
	defer reader.Close()

	// for _, entity := range entities {
	// 	fmt.Println(entity)
	// }

	entities := map[int]models.ClienteTransacao{}
	if err := repo.GetAll(&entities, 1); err != nil {
		t.Fatalf("error %t", err)
	}
}

func TestClientTransactionRepoCreateDebito(t *testing.T) {
	writer, reader := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(writer, reader)

	if err := repo.Add(1, 100, "d", "debito"); err != nil {
		t.Fatalf("error %t", err)
	}
}

func TestClientTransactionRepoCreateCredito(t *testing.T) {
	writer, reader := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(writer, reader)

	if err := repo.Add(1, 100, "c", "credito"); err != nil {
		t.Fatalf("error %t", err)
	}
}

func BenchmarkClientTransactionRepoGet(b *testing.B) {
	writer, reader := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(writer, reader)

	for i := 0; i < b.N; i++ {
		entities := map[int]models.ClienteTransacao{}
		if err := repo.GetAll(&entities, 1); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkClientTransactionRepoCreate(b *testing.B) {
	writer, reader := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(writer, reader)

	for i := 0; i < b.N; i++ {
		if err := repo.Add(1, 100, "c", "credito"); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkClientTransactionRepoGet10Mil(b *testing.B) {
	writer, reader := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(writer, reader)

	for i := 0; i < 10_001; i++ {
		entities := map[int]models.ClienteTransacao{}
		if err := repo.GetAll(&entities, 1); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkClientTransactionRepoCreate10Mil(b *testing.B) {
	writer, reader := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(writer, reader)

	for i := 0; i < 10_001; i++ {
		if err := repo.Add(1, 100, "c", "credito"); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}
