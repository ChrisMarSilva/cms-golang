package main

import (
	"fmt"
	"testing"
)

// go test -bench .
// go test -bench=Improved .
// go test -bench . -benchmem
// go test -bench . -benchtime=1000x
// go test -bench . -benchtime=2s
// go test -bench . -count=5

var num = 100

var table = []struct{
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 74382},
	{input: 382399},
}

func BenchmarkPrimeNumbers(b *testing.B){
	for i := 0; i < b.N; i++ {
		primeNumbers(num)
	}
}

func BenchmarkPrimeNumbers2(b *testing.B){
	for i := 0; i < b.N; i++ {
		primeNumbers2(num)
	}
}

func BenchmarkPrimeNumbersImproved(b *testing.B){
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isPrimeImproved(v.input)
			}
		})		
	}
}

func BenchmarkPrimeNumbersImproved2(b *testing.B){
	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				isPrimeImproved2(v.input)
			}
		})		
	}
}
