package tests

import (
	"strconv"
	"sync"
	"testing"

	"github.com/chrismarsilva/cms.golang.benchmarks/models"
)

// go test -bench . -benchmem
// go test -bench=. -benchmem
// go test -bench=. -benchmem ./tests/list

// BenchmarkTest_4_0-4     31478712                92.97 ns/op            0 B/op          0 allocs/op
// BenchmarkTest_2_2-4     10690956               120.0 ns/op            63 B/op          1 allocs/op
// BenchmarkTest_1_0-4      5405918               291.8 ns/op           262 B/op          1 allocs/op
// BenchmarkTest_2_1-4      4800096               272.9 ns/op           352 B/op          1 allocs/op
// BenchmarkTest_2_0-4      4282762               336.6 ns/op           265 B/op          1 allocs/op
// BenchmarkTest_3_1-4      3257796               482.5 ns/op            96 B/op          2 allocs/op
// BenchmarkTest_3_0-4      2093244               656.4 ns/op           248 B/op          2 allocs/op
// BenchmarkTest_5_0-4        10000            108847 ns/op           15950 B/op       1901 allocs/op

func BenchmarkTest_1_0(b *testing.B) {
	var rows []models.Person

	for i := 0; i < b.N; i++ {
		rows = append(rows, models.Person{ID: i, Name: "Person " + strconv.Itoa(i), Age: i, Active: true})
	}
}

func BenchmarkTest_2_0(b *testing.B) {
	rows := make([]models.Person, 0)

	for i := 0; i < b.N; i++ {
		rows = append(rows, models.Person{ID: i, Name: "Person " + strconv.Itoa(i), Age: i, Active: true})
	}
}

func BenchmarkTest_2_1(b *testing.B) {
	rows := make([]models.Person, b.N) // já aloca capacidade necessária

	for i := 0; i < b.N; i++ {
		rows = append(rows, models.Person{ID: i, Name: "Person " + strconv.Itoa(i), Age: i, Active: true})
	}
}

func BenchmarkTest_2_2(b *testing.B) {
	rows := make([]models.Person, b.N)

	for i := 0; i < b.N; i++ {
		rows[i] = models.Person{ID: i, Name: "Person " + strconv.Itoa(i), Age: i, Active: true}
	}
}

func BenchmarkTest_3_0(b *testing.B) {
	rows := make(map[int]models.Person, 0)

	for i := 0; i < b.N; i++ {
		rows[i] = models.Person{ID: i, Name: "Person " + strconv.Itoa(i), Age: i, Active: true}
	}
}

func BenchmarkTest_3_1(b *testing.B) {
	rows := make(map[int]models.Person, b.N) // já aloca espaço suficiente

	for i := 0; i < b.N; i++ {
		rows[i] = models.Person{ID: i, Name: "Person " + strconv.Itoa(i), Age: i, Active: true}
	}
}

func BenchmarkTest_4_0(b *testing.B) { // BenchmarkCopySlice
	source := make([]models.Person, b.N)
	for i := 0; i < b.N; i++ {
		source[i] = models.Person{ID: i, Name: "Person " + strconv.Itoa(i), Age: i, Active: true}
	}

	rows := make([]models.Person, b.N)
	//b.ResetTimer() // só mede a cópia
	copy(rows, source)
}

var pool = sync.Pool{
	New: func() interface{} {
		return make([]models.Person, 0, 1000)
	},
}

func BenchmarkTest_5_0(b *testing.B) { // BenchmarkSyncPool
	for i := 0; i < b.N; i++ {
		rows := pool.Get().([]models.Person)
		rows = rows[:0] // reseta
		for j := 0; j < 1000; j++ {
			rows = append(rows, models.Person{ID: j, Name: "Person " + strconv.Itoa(j), Age: j, Active: true})
		}

		pool.Put(rows)
	}
}
