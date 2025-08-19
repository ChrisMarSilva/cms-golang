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

// go test -bench=.
// go test -bench . -benchmem
// go test -bench=. -benchmem
// go test -bench=. -benchmem ./tests
// go test -bench=. -benchmem ./tests/addMany
// go test -bench=. -benchmem ./tests/addMany/pgx_test.go
// go test -bench=. -benchmem ./tests/addMany/gorm_test.go
// go test -bench=. -benchmem ./gorm_test.go

// go test -bench=BenchmarkGormAddMany1 -benchtime=30s
// go test -bench=BenchmarkGormAddMany1 -benchtime=30s -count=5
// go test -bench=BenchmarkGormAddMany1 -benchtime=30s -count=5 | tee mc_square.txt
// go test -bench=BenchmarkGormAddMany1 -cpuprofile cpu.prof

// go test -bench=. -benchmem ./gorm_test.go
// go test -bench=. -benchmem=100x ./gorm_test.go   // Definindo uma contagem de iterações explícita (100 vezes):
// go test -bench=. -benchmem=10s ./gorm_test.go    // Especifique o tempo (10 segundos):

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
	"log/slog"
	"sort"
	"sync"
	"time"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/stores"
)

const (
	QTDE = 1000
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

	// PgxP-Add-Many-2 count=1000 duration=40.8323ms
	// PgxP-Add-Many-4 count=1000 duration=327.2482ms
	// SqlX-Add-Many-2 count=1000 duration=359.5803ms
	// Gorm-Add-Many-2 count=1000 duration=453.1124ms
	// Gorm-Add-Many-3 count=1000 duration=479.3799ms
	// PgxP-Add-Many-3 count=1000 duration=513.7703ms
	// SqlX-Add-Many-1 count=1000 duration=2.4044577s
	// PgxP-Add-Many-1 count=1000 duration=2.443251s
	// Gorm-Add-Many-1 count=1000 duration=2.4829096s
	// SqlC-Add-Many-1 count=1000 duration=2.6225748s

	// BenchmarkGormAddMany2               44          23322598 ns/op         1825594 B/op      30918 allocs/op
	// BenchmarkPgxPoolAddMany4            40          29630278 ns/op         1166991 B/op      27972 allocs/op
	// BenchmarkGormAddMany3               32          32297978 ns/op        12957082 B/op      37796 allocs/op
	// BenchmarkPgxPoolAddMany3            30          45883727 ns/op        12384616 B/op      31738 allocs/op
	// BenchmarkPgxPoolAddMany2            24          42130671 ns/op         1680157 B/op      33944 allocs/op
	// BenchmarkSqlXAddMany2               22          82011186 ns/op        12563313 B/op      31751 allocs/op
	// BenchmarkGormAddMany1                2         784265250 ns/op         4943772 B/op      65018 allocs/op
	// BenchmarkPgxPoolAddMany1             2         805484200 ns/op          893408 B/op      29926 allocs/op
	// BenchmarkSqlXAddMany1                2         785752300 ns/op         1069760 B/op      32931 allocs/op
	// BenchmarkSqlCAddMany1                1        2332101100 ns/op         1505768 B/op      43959 allocs/op

	var wg sync.WaitGroup
	bancos := make([]*Banco, 0)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManyPgxPool(ctx, QTDE)
		bancos = append(bancos, &Banco{Nome: "PgxP-Add-Many-1", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManyPgxPool2(ctx, QTDE)
		bancos = append(bancos, &Banco{Nome: "PgxP-Add-Many-2", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManyPgxPool3(ctx, QTDE)
		bancos = append(bancos, &Banco{Nome: "PgxP-Add-Many-3", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManyPgxPool4(ctx, QTDE)
		bancos = append(bancos, &Banco{Nome: "PgxP-Add-Many-4", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManyGorm(ctx, QTDE)
		bancos = append(bancos, &Banco{Nome: "Gorm-Add-Many-1", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManyGorm2(ctx, QTDE)
		bancos = append(bancos, &Banco{Nome: "Gorm-Add-Many-2", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManyGorm3(ctx, QTDE)
		bancos = append(bancos, &Banco{Nome: "Gorm-Add-Many-3", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManySqlX(ctx, QTDE)
		bancos = append(bancos, &Banco{Nome: "SqlX-Add-Many-1", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManySqlX2(ctx, QTDE)
		bancos = append(bancos, &Banco{Nome: "SqlX-Add-Many-2", Tempo: tempo})
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		tempo := stores.AddManySqlC(ctx, QTDE)
		bancos = append(bancos, &Banco{Nome: "SqlC-Add-Many-1", Tempo: tempo})
	}(&wg)

	wg.Wait()

	PrintAll("TesteAddMany", QTDE, bancos)

	slog.Info("TesteAddMany: Finished")
}
