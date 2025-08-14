package tests

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/bytedance/sonic"
	jsoniter "github.com/json-iterator/go"
)

// go test -bench . -benchmem
// go test -bench=. -benchmem

type Person struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Active bool   `json:"active"`
}

var (
	sampleData1    []Person
	sampleData10   []Person
	sampleData100  []Person
	sampleData1000 []Person

	jsonStd     = json.Marshal
	jsonStdUn   = json.Unmarshal
	jsonIter    = jsoniter.Marshal
	jsonIterUn  = jsoniter.Unmarshal
	jsonSonic   = sonic.Marshal
	jsonSonicUn = sonic.Unmarshal
)

func init() {
	sampleData1 = generateData(1)
	sampleData10 = generateData(10)
	sampleData100 = generateData(100)
	sampleData1000 = generateData(1000)
}

func generateData(n int) []Person {
	data := make([]Person, n)
	for i := 0; i < n; i++ {
		data[i] = Person{ID: i, Name: "Person " + strconv.Itoa(i), Age: 30, Active: true}
	}

	return data
}

func benchmarkMarshal(b *testing.B, marshal func(v interface{}) ([]byte, error), data interface{}) {
	b.ReportAllocs() // relatar alocações
	for i := 0; i < b.N; i++ {
		_, _ = marshal(data)
	}
}

func benchmarkUnmarshal(b *testing.B, unmarshal func(data []byte, v interface{}) error, data interface{}) {
	raw, _ := json.Marshal(data) // usar padrão para gerar entrada estável
	b.ReportAllocs()             // relatar alocações
	b.ResetTimer()               // reiniciar o timer após a preparação
	for i := 0; i < b.N; i++ {
		var v interface{}
		_ = unmarshal(raw, &v)
	}
}

// BenchmarkStd_Marshal_1-8         5618650               212.3 ns/op            64 B/op          1 allocs/op
// BenchmarkIter_Marshal_1-8        7330096               156.4 ns/op            64 B/op          1 allocs/op
// BenchmarkSonic_Marshal_1-8       8683080               134.6 ns/op            83 B/op          2 allocs/op
// BenchmarkStd_Unmarshal_1-8       1000000              1135 ns/op             616 B/op         16 allocs/op
// BenchmarkIter_Unmarshal_1-8      1372824               844.9 ns/op           528 B/op         17 allocs/op
// BenchmarkSonic_Unmarshal_1-8     1867459               642.1 ns/op           764 B/op          9 allocs/op

// BenchmarkStd_Marshal_10-8                 869325              1272 ns/op             512 B/op          1 allocs/op
// BenchmarkIter_Marshal_10-8               1000000              1151 ns/op             512 B/op          1 allocs/op
// BenchmarkSonic_Marshal_10-8              2308328               502.4 ns/op           545 B/op          2 allocs/op
// BenchmarkStd_Unmarshal_10-8               135774              8190 ns/op            4696 B/op        119 allocs/op
// BenchmarkIter_Unmarshal_10-8              153519              7947 ns/op            5330 B/op        156 allocs/op
// BenchmarkSonic_Unmarshal_10-8             281396              4309 ns/op            4544 B/op         54 allocs/op

// BenchmarkStd_Marshal_100-8                 97617             11934 ns/op            5378 B/op          1 allocs/op
// BenchmarkIter_Marshal_100-8               109899             10155 ns/op            5379 B/op          1 allocs/op
// BenchmarkSonic_Marshal_100-8              299432              3773 ns/op            5526 B/op          2 allocs/op
// BenchmarkStd_Unmarshal_100-8               15552             78605 ns/op           45384 B/op       1202 allocs/op
// BenchmarkIter_Unmarshal_100-8              17467             67137 ns/op           52510 B/op       1509 allocs/op
// BenchmarkSonic_Unmarshal_100-8             31269             38426 ns/op           46635 B/op        507 allocs/op

// BenchmarkStd_Marshal_1000-8                 9591            123955 ns/op           57393 B/op          1 allocs/op
// BenchmarkIter_Marshal_1000-8               12165             97723 ns/op           57365 B/op          1 allocs/op
// BenchmarkSonic_Marshal_1000-8              32797             36339 ns/op           58627 B/op          2 allocs/op
// BenchmarkStd_Unmarshal_1000-8               1436            780794 ns/op          443307 B/op      12005 allocs/op
// BenchmarkIter_Unmarshal_1000-8              1626            674849 ns/op          515458 B/op      15012 allocs/op
// BenchmarkSonic_Unmarshal_1000-8             2894            379473 ns/op          452181 B/op       5009 allocs/op

// --------- Marshal 1 ---------
func BenchmarkStd_Marshal_1(b *testing.B)   { benchmarkMarshal(b, jsonStd, sampleData1) }
func BenchmarkIter_Marshal_1(b *testing.B)  { benchmarkMarshal(b, jsonIter, sampleData1) }
func BenchmarkSonic_Marshal_1(b *testing.B) { benchmarkMarshal(b, jsonSonic, sampleData1) }

// --------- Unmarshal 1 ---------
func BenchmarkStd_Unmarshal_1(b *testing.B)   { benchmarkUnmarshal(b, jsonStdUn, sampleData1) }
func BenchmarkIter_Unmarshal_1(b *testing.B)  { benchmarkUnmarshal(b, jsonIterUn, sampleData1) }
func BenchmarkSonic_Unmarshal_1(b *testing.B) { benchmarkUnmarshal(b, jsonSonicUn, sampleData1) }

// --------- Marshal 10 ---------
func BenchmarkStd_Marshal_10(b *testing.B)   { benchmarkMarshal(b, jsonStd, sampleData10) }
func BenchmarkIter_Marshal_10(b *testing.B)  { benchmarkMarshal(b, jsonIter, sampleData10) }
func BenchmarkSonic_Marshal_10(b *testing.B) { benchmarkMarshal(b, jsonSonic, sampleData10) }

// --------- Unmarshal 10 ---------
func BenchmarkStd_Unmarshal_10(b *testing.B)   { benchmarkUnmarshal(b, jsonStdUn, sampleData10) }
func BenchmarkIter_Unmarshal_10(b *testing.B)  { benchmarkUnmarshal(b, jsonIterUn, sampleData10) }
func BenchmarkSonic_Unmarshal_10(b *testing.B) { benchmarkUnmarshal(b, jsonSonicUn, sampleData10) }

// --------- Marshal 100 ---------
func BenchmarkStd_Marshal_100(b *testing.B)   { benchmarkMarshal(b, jsonStd, sampleData100) }
func BenchmarkIter_Marshal_100(b *testing.B)  { benchmarkMarshal(b, jsonIter, sampleData100) }
func BenchmarkSonic_Marshal_100(b *testing.B) { benchmarkMarshal(b, jsonSonic, sampleData100) }

// --------- Unmarshal 100 ---------
func BenchmarkStd_Unmarshal_100(b *testing.B)   { benchmarkUnmarshal(b, jsonStdUn, sampleData100) }
func BenchmarkIter_Unmarshal_100(b *testing.B)  { benchmarkUnmarshal(b, jsonIterUn, sampleData100) }
func BenchmarkSonic_Unmarshal_100(b *testing.B) { benchmarkUnmarshal(b, jsonSonicUn, sampleData100) }

// --------- Marshal 1000 ---------
func BenchmarkStd_Marshal_1000(b *testing.B)   { benchmarkMarshal(b, jsonStd, sampleData1000) }
func BenchmarkIter_Marshal_1000(b *testing.B)  { benchmarkMarshal(b, jsonIter, sampleData1000) }
func BenchmarkSonic_Marshal_1000(b *testing.B) { benchmarkMarshal(b, jsonSonic, sampleData1000) }

// --------- Unmarshal 1000 ---------
func BenchmarkStd_Unmarshal_1000(b *testing.B)   { benchmarkUnmarshal(b, jsonStdUn, sampleData1000) }
func BenchmarkIter_Unmarshal_1000(b *testing.B)  { benchmarkUnmarshal(b, jsonIterUn, sampleData1000) }
func BenchmarkSonic_Unmarshal_1000(b *testing.B) { benchmarkUnmarshal(b, jsonSonicUn, sampleData1000) }
