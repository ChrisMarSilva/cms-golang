package main

// go mod init github.com/chrismarsilva/cms.golang.benchmarks.bd
// go get -u github.com/google/uuid
// go get -u github.com/brianvoe/gofakeit/v7
// go get -u github.com/jackc/pgx/v5
// go get -u github.com/jackc/pgx/v5/pgxpool
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

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/models"
	"github.com/chrismarsilva/cms.golang.benchmarks.bd/stores"
)

func main() {
	slog.Info("Running benchmarks...")

	ctx := context.Background()
	models.FetchAll()
	stores.FetchAllPgxPool(ctx)
	//stores.FetchAllGorm(ctx)
	//stores.FetchAllSqlX(ctx)
	//stores.FetchAllSqlC(ctx)

	slog.Info("Finished benchmarks")
}
