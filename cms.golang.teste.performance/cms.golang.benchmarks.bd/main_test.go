package main

// go test -bench . -benchmem

import (
	"context"
	"testing"

	"github.com/chrismarsilva/cms.golang.benchmarks.bd/stores"
)

func BenchmarkPgxPoolGetAll(b *testing.B) {
	ctx := context.Background()
	dbPgxPool := stores.NewDatabasePgxPool(ctx)
	//defer dbPgxPool.Close()

	b.ReportAllocs() // relatar alocações
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		stores.GetAllPgxPool(ctx, dbPgxPool)
		//defer rows.Close()
	}
}
