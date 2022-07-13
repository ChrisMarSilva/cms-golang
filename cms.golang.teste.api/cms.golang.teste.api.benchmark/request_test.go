package main

import (
	"net/http"
	"strconv"
	"testing"
)

// go test
// go test -v
// go test -bench=Add
// go test -run TestInvertText2 -v
// go test -run TestIsPalindrome -v

// go test -bench=.
// go test -bench=. -count 5 -run=^#
// go test -run=XXX -bench . -benchmem

func BenchmarkRequestGin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, err := http.Get("http://localhost:" + strconv.Itoa(portGin))
		if err != nil {
			b.FailNow()
		}
		if resp.StatusCode != http.StatusOK {
			print("error")
		}
		err = resp.Body.Close()
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkRequestEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, err := http.Get("http://localhost:" + strconv.Itoa(portEcho))
		if err != nil {
			b.FailNow()
		}
		if resp.StatusCode != http.StatusOK {
			print("error")
		}
		err = resp.Body.Close()
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkRequestMux(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, err := http.Get("http://localhost:" + strconv.Itoa(portMux))
		if err != nil {
			b.FailNow()
		}
		if resp.StatusCode != http.StatusOK {
			print("error")
		}
		err = resp.Body.Close()
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkRequestChi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, err := http.Get("http://localhost:" + strconv.Itoa(portChi))
		if err != nil {
			b.FailNow()
		}
		if resp.StatusCode != http.StatusOK {
			print("error")
		}
		err = resp.Body.Close()
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkRequestFiber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, err := http.Get("http://localhost:" + strconv.Itoa(portFiber))
		if err != nil {
			b.FailNow()
		}
		if resp.StatusCode != http.StatusOK {
			print("error")
		}
		err = resp.Body.Close()
		if err != nil {
			b.FailNow()
		}
	}
}

func BenchmarkRequestFastHttpRouting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		resp, err := http.Get("http://localhost:" + strconv.Itoa(portFastHttpRouting))
		if err != nil {
			b.FailNow()
		}
		if resp.StatusCode != http.StatusOK {
			print("error")
		}
		err = resp.Body.Close()
		if err != nil {
			b.FailNow()
		}
	}
}
