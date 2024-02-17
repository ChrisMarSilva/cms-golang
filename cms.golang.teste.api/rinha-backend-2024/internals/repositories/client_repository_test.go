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
// go test -run GetClientDbPadrao -v
// go test -run=GetClientDbPadrao -bench . -benchmem

func GetClientDbPadrao() (*sqlx.DB, *sqlx.DB) {
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

	if err := repo.UpdSaldo(1, 100, "d"); err != nil {
		t.Fatalf("error %t", err)
	}
}

func TestClientRepoUpdateCredito(t *testing.T) {
	writer, reader := GetClientDbPadrao()
	repo := repositories.NewClientRepository(writer, reader)

	if err := repo.UpdSaldo(1, 100, "c"); err != nil {
		t.Fatalf("error %t", err)
	}
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

	for i := 0; i < b.N; i++ {
		if err := repo.UpdSaldo(1, 100, "c"); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkClientRepoGetAndUpdate(b *testing.B) {
	writer, reader := GetClientDbPadrao()
	repo := repositories.NewClientRepository(writer, reader)

	for i := 0; i < b.N; i++ {
		var entity models.Cliente
		if err := repo.Get(&entity, 1); err != nil {
			b.Fatalf("error %t", err)
		}
		if err := repo.UpdSaldo(1, 100, "c"); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkClientRepoGet10Mil(b *testing.B) {
	writer, reader := GetClientDbPadrao()
	repo := repositories.NewClientRepository(writer, reader)

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

	for i := 0; i < 10_001; i++ {
		if err := repo.UpdSaldo(1, 100, "c"); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}

func BenchmarkClientRepoGetAndUpdate10Mil(b *testing.B) {
	writer, reader := GetClientDbPadrao()
	repo := repositories.NewClientRepository(writer, reader)

	for i := 0; i < 10_001; i++ {
		var entity models.Cliente
		if err := repo.Get(&entity, 1); err != nil {
			b.Fatalf("error %t", err)
		}
		if err := repo.UpdSaldo(1, 100, "c"); err != nil {
			b.Fatalf("error %t", err)
		}
	}
}
