package main

import (
	"fmt"
	"strings"
	"testing"
)

// go test -bench . -benchmem

const (
	smallString = "HunCoding"
	longString = "HunCodingStingBuilderTestPerformancHunCodingStingBuilderTestPerformanceeHunCodingStingBuilderTestPerformance"
	// longString = "HunCodingStingBuilderTestPerformancHunCodingStingBuilderTestPerformanceeHunCodingStingBuilderTestPerformanceHunCodingStingBuilderTestPerformancHunCodingStingBuilderTestPerformanceeHunCodingStingBuilderTestPerformanceHunCodingStingBuilderTestPerformancHunCodingStingBuilderTestPerformanceeHunCodingStingBuilderTestPerformanceHunCodingStingBuilderTestPerformancHunCodingStingBuilderTestPerformanceeHunCodingStingBuilderTestPerformanceHunCodingStingBuilderTestPerformancHunCodingStingBuilderTestPerformanceeHunCodingStingBuilderTestPerformanceHunCodingStingBuilderTestPerformancHunCodingStingBuilderTestPerformanceeHunCodingStingBuilderTestPerformance"
)

func generateStringArray(s string) (data []string){
	for i := 0; i < 100; i++ {
		data = append(data, s)
	}
	return
}

// func BenchmarkWithSprintfSmallerString(b *testing.B){
// 	data := generateStringArray(smallString)
// 	var s string
// 	for n := 0; n < b.N; n++ {
// 		s = fmt.Sprintf(s, data)
// 		_ = s
// 	}
// }

func BenchmarkWithSprintfLongerString(b *testing.B){
	data := generateStringArray(longString)
	var s string
	for n := 0; n < b.N; n++ {
		s = fmt.Sprintf(s, data)
		_ = s
	}
}

// func BenchmarkWithOperatorfSmallerString(b *testing.B){
// 	data := generateStringArray(smallString)

// 	f := func(s []string) (allStr string){
// 		for _, x := range s {
// 			allStr += x
// 		}
// 		return
// 	}

// 	for n := 0; n < b.N; n++ {
// 		_ = f(data)
// 	}
// }

func BenchmarkWithOperatorLongerString(b *testing.B){
	data := generateStringArray(longString)

	f := func(s []string) (allStr string){
		for _, x := range s {
			allStr += x
		}
		return
	}

	for n := 0; n < b.N; n++ {
		_ = f(data)
	}
}

// func BenchmarkWithJoinSmallerString(b *testing.B){
// 	data := generateStringArray(smallString)
// 	var s string
// 	for n := 0; n < b.N; n++ {
// 		s = strings.Join(data, "")
// 		_ = s
// 	}
// }

func BenchmarkWithJoinfLongerString(b *testing.B){
	data := generateStringArray(longString)
	var s string
	for n := 0; n < b.N; n++ {
		s = strings.Join(data, "")
		_ = s
	}
}

// func BenchmarkWithBuilderSmallerString(b *testing.B){
// 	data := generateStringArray(smallString)
// 	var sb strings.Builder
// 	var s string
// 	for n := 0; n < b.N; n++ {
// 		for _, s := range data {
// 			sb.WriteString(s)
// 		}
// 		s = sb.String()
// 		_ = s
// 	}
// }

func BenchmarkWithBuilderLongerString(b *testing.B){
	data := generateStringArray(longString)
	var sb strings.Builder
	var s string
	for n := 0; n < b.N; n++ {
		for _, s := range data {
			sb.WriteString(s)
		}
		s = sb.String()
		_ = s
	}
}
