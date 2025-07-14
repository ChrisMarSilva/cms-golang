package main

// go test -bench=.
// go test -bench=. -cpuprofile=cpu.out -memprofile=mem.out
// go test -bench=. -benchmem -cpuprofile=cpu.prof -memprofile=mem.prof

// go tool pprof cpu.prof
// go tool pprof -web cpu.prof
// top 20
// web

import (
	"context"
	"fmt"
	"runtime"
	"testing"

	_ "net/http/pprof"

	"github.com/google/uuid"
)

// 200655944100 ns/op   = 3,344265735 seconds
// 180370038700 ns/op   = 3,006167311 seconds
//     0.003549 ns/op   = 0,003549 milliseconds

func BenchmarkProcessJobs(b *testing.B) {
	sizes := []int{10_000} // 1, 10, 100, 500, 1_000, 2_000, 5_000  // 1_000, 10_000, 100_000 // 100_000, 500_000, 1_000_000

	for _, size := range sizes {
		b.Run(fmt.Sprintf("%dJobs", size), func(b *testing.B) {
			b.ReportAllocs()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			jobs := make(chan Job, maxQueueSize)
			for i := 0; i < runtime.NumCPU(); i++ {
				go processWorkers(ctx, i, jobs)
			}
			go resetUseDefault()

			b.ResetTimer()
			for i := 0; i < size; i++ {
				jobs <- Job{ID: uuid.New()}
			}
			close(jobs)
			cancel()
		})
	}
}

// func BenchmarkMillionJobs(b *testing.B) {
// 	sizes := []int{100_000, 500_000, 1_000_000}

// 	for _, size := range sizes {
// 		b.Run(fmt.Sprintf("%dJobs", size), func(b *testing.B) {
// 			b.ReportAllocs()
// 			b.ResetTimer()
// 			runBenchmark(size)
// 		})
// 	}
// }

// func runBenchmark(n int) {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	jobs := make(chan Job, maxQueueSize)
// 	go resetUseDefault()
// 	go processWorkers(ctx, 1, jobs)

// 	for i := 0; i < n; i++ {
// 		jobs <- Job{ID: uuid.New()}
// 	}
// 	close(jobs)
// }
