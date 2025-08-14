package tests

import (
	"testing"

	gofrs "github.com/gofrs/uuid/v5"
	google "github.com/google/uuid"
	shortuuid "github.com/lithammer/shortuuid/v4"
	satori "github.com/satori/go.uuid"
	ksuid "github.com/segmentio/ksuid"
)

// go test -bench . -benchmem
// go test -bench=. -benchmem

// Satori UUID: 390fa672-1f4e-4993-a679-5120d3c760df
// Gofrs  UUID: ff45accd-a31e-4476-b530-8c4c20e9b936
// Google UUID: bb5d81e0-e03b-40e9-9f82-3f128426722c
// KSUID  UUID: 31HjFSi0RbbR3zXF8NUV5yq7JcL
// Short  UUID: 6YiTFdJtfdZYasfiTmctzn

// BenchmarkSatoriUUID_New-8        9.109.189               119.8 ns/op             0 B/op          0 allocs/op
// BenchmarkGofrsUUID_New-8         8.520.874               140.7 ns/op            16 B/op          1 allocs/op
// BenchmarkGoogleUUID_New-8        8.277.855               140.1 ns/op            16 B/op          1 allocs/op

// BenchmarkKSUID_New-8             7612888               147.7 ns/op             0 B/op          0 allocs/op
// BenchmarkShortUUID_New-8         5534548               208.9 ns/op            40 B/op          2 allocs/op

func BenchmarkGoogleUUID_New(b *testing.B) {
	b.ReportAllocs() // relatar alocações
	for i := 0; i < b.N; i++ {
		_ = google.New()
	}
}
func BenchmarkGofrsUUID_New(b *testing.B) {
	b.ReportAllocs() // relatar alocações
	for i := 0; i < b.N; i++ {
		_, _ = gofrs.NewV4()
	}
}
func BenchmarkSatoriUUID_New(b *testing.B) {
	b.ReportAllocs() // relatar alocações
	for i := 0; i < b.N; i++ {
		_ = satori.NewV4()
	}
}
func BenchmarkKSUID_New(b *testing.B) {
	b.ReportAllocs() // relatar alocações
	for i := 0; i < b.N; i++ {
		_ = ksuid.New()
	}
}
func BenchmarkShortUUID_New(b *testing.B) {
	b.ReportAllocs() // relatar alocações
	for i := 0; i < b.N; i++ {
		_ = shortuuid.New()
	}
}
