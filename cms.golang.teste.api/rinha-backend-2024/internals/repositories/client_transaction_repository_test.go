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

func GetClientTransactionDbPadrao() *sqlx.DB {
	viper.AddConfigPath("./")
	viper.SetConfigFile("../../cmd/api-server/.env")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}

	driverDB := databases.DatabasePostgres{}
	driverDB.StartDB()
	return driverDB.GetDatabase()
}

func TestClientTransactionRepoGet(t *testing.T) {
	db := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(db)

	// for _, entity := range entities {
	// 	fmt.Println(entity)
	// }

	var entities []models.ClienteTransacao // entities := []models.ClienteTransacao{}
	if err := repo.GetAll(&entities, 1); err != nil {
		t.Fatalf("error %t", err)
	}
}

func TestClientTransactionRepoCreateDebito(t *testing.T) {
	db := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(db)

	if err := repo.Add(1, 100, "d", "debito"); err != nil {
		t.Fatalf("error %t", err)
	}
}

func TestClientTransactionRepoCreateCredito(t *testing.T) {
	db := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(db)

	if err := repo.Add(1, 100, "c", "credito"); err != nil {
		t.Fatalf("error %t", err)
	}
}

func BenchmarkClientTransactionRepoGet(b *testing.B) {
	db := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(db)

	for i := 0; i < b.N; i++ {
		var entities []models.ClienteTransacao
		if err := repo.GetAll(&entities, 1); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkClientTransactionRepoCreate(b *testing.B) {
	db := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(db)

	for i := 0; i < b.N; i++ {
		if err := repo.Add(1, 100, "c", "credito"); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkClientTransactionRepoGet10Mil(b *testing.B) {
	db := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(db)

	for i := 0; i < 10_001; i++ {
		var entities []models.ClienteTransacao
		if err := repo.GetAll(&entities, 1); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkClientTransactionRepoCreate10Mil(b *testing.B) {
	db := GetClientTransactionDbPadrao()
	repo := repositories.NewClientTransactionRepository(db)

	for i := 0; i < 10_001; i++ {
		if err := repo.Add(1, 100, "c", "credito"); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}
