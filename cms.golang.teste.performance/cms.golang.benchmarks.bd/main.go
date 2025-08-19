package main

// go mod init github.com/chrismarsilva/cms.golang.benchmarks.bd
// go get -u github.com/google/uuid
// go get -u github.com/brianvoe/gofakeit/v7
// go get -u github.com/jackc/pgx/v5
// go get -u github.com/jackc/pgx/v5/pgxpool
// go get -u github.com/jackc/pgx/v5/stdlib
// go get -u github.com/jackc/pgtype
// go get -u github.com/timtoronto634/pgx-slog
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/postgres
// go get -u github.com/jmoiron/sqlx
// go get -u github.com/lib/pq
// go mod tidy
// go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
// sqlc generate
// sqlc push
// go run main.go

import (
	"context"
	"log/slog"
	"sort"
	"sync"
	"time"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/stores"
)

func main() {
	slog.Info("Running benchmarks...")

	ctx := context.Background()

	//models.FetchAll()
	//stores.FetchAllPgxPool(ctx)
	//stores.FetchAllGorm(ctx)
	//stores.FetchAllSqlX(ctx)
	//stores.FetchAllSqlC(ctx)

	TesteClear(ctx)
	//TesteAddOne(ctx)
	TesteAddMany(ctx)

	slog.Info("Finished benchmarks")
}

type Banco struct {
	Nome  string
	Tempo time.Duration
}

func PrintAll(nome string, count int, bancos []*Banco) {
	sort.Slice(bancos, func(i, j int) bool {
		return bancos[i].Tempo < bancos[j].Tempo
	})

	for _, b := range bancos {
		slog.Info(nome+": Resultado",
			slog.String("nome", b.Nome),
			slog.Int("count", 1000),
			slog.Duration("duration", b.Tempo))
	}
}

func TesteClear(ctx context.Context) {
	slog.Info("TesteClear: Getting started")

	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		stores.ClearPgxPool(ctx)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		stores.ClearGorm(ctx)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		stores.ClearSqlX(ctx)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		stores.ClearSqlC(ctx)
	}(&wg)

	wg.Wait()

	slog.Info("TesteClear: Finished")
}

func TesteAddOne(ctx context.Context) {
	slog.Info("TesteAddOne: Getting started")

	var wg sync.WaitGroup
	count := 1_000
	bancos := make([]*Banco, 0)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddOnePgxPool(ctx, count)
		bancos = append(bancos, &Banco{Nome: "PgxP-Add-One", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddOneGorm(ctx, count)
		bancos = append(bancos, &Banco{Nome: "Gorm-Add-One", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddOneSqlX(ctx, count)
		bancos = append(bancos, &Banco{Nome: "SqlX-Add-One", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddOneSqlC(ctx, count)
		bancos = append(bancos, &Banco{Nome: "SqlC-Add-One", Tempo: tempo})
	}(&wg)

	wg.Wait()

	PrintAll("TesteAddOne", count, bancos)

	slog.Info("TesteAddOne: Finished")
}

func TesteAddMany(ctx context.Context) {
	slog.Info("TesteAddMany: Getting started")

	// var tcases = []struct {
	// 	RecordsToCreate int
	// }{
	// 	{RecordsToCreate: 100},
	// 	{RecordsToCreate: 1_000},
	// 	{RecordsToCreate: 10_000},
	// 	{RecordsToCreate: 100_000},
	// 	{RecordsToCreate: 300_000},
	// 	{RecordsToCreate: 500_000},
	// 	{RecordsToCreate: 1_000_000},
	// }

	var wg sync.WaitGroup
	count := 1_000
	bancos := make([]*Banco, 0)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManyPgxPool(ctx, count)
		bancos = append(bancos, &Banco{Nome: "PgxP-Add-Many-1", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManyPgxPool2(ctx, count)
		bancos = append(bancos, &Banco{Nome: "PgxP-Add-Many-2", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManyPgxPool3(ctx, count)
		bancos = append(bancos, &Banco{Nome: "PgxP-Add-Many-3", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManyPgxPool4(ctx, count)
		bancos = append(bancos, &Banco{Nome: "PgxP-Add-Many-4", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManyGorm(ctx, count)
		bancos = append(bancos, &Banco{Nome: "Gorm-Add-Many-1", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManyGorm2(ctx, count)
		bancos = append(bancos, &Banco{Nome: "Gorm-Add-Many-2", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManyGorm3(ctx, count)
		bancos = append(bancos, &Banco{Nome: "Gorm-Add-Many-3", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManySqlX(ctx, count)
		bancos = append(bancos, &Banco{Nome: "SqlX-Add-Many-1", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManySqlX2(ctx, count)
		bancos = append(bancos, &Banco{Nome: "SqlX-Add-Many-2", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManySqlC(ctx, count)
		bancos = append(bancos, &Banco{Nome: "SqlC-Add-Many-1", Tempo: tempo})
	}(&wg)

	wg.Wait()

	PrintAll("TesteAddMany", count, bancos)

	slog.Info("TesteAddMany: Finished")
}
