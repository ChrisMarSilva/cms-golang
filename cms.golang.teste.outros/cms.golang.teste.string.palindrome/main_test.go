package main

import (
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

func TestInvertText(t *testing.T) {

	type Case struct {
		input  string
		result string
	}

	testCases := []Case{
		{input: "ana", result: "ana"},
		{input: "mamam", result: "mamam"},
		{input: "leonardo", result: "odranoel"},
	}

	for _, testCase := range testCases {
		if invertText(testCase.input) != testCase.result {
			t.Error("FAIL", testCase)
		}
	}

}

func TestIsPalindrome(t *testing.T) {

	type Case struct {
		input  string
		result bool // expectedResult
	}

	testCases := []Case{
		{input: "ana", result: true},
		{input: "mamam", result: true},
		{input: "leonardo", result: false},
	}

	for _, testCase := range testCases {
		if isPalindrome(testCase.input) != testCase.result {
			t.Error("FAIL", testCase)
		}
	}

}

func TestIsPalindrome2(t *testing.T) {

	type Case struct {
		input  string
		result bool // expectedResult
	}

	testCases := []Case{
		{input: "ana", result: true},
		{input: "mamam", result: true},
		{input: "leonardo", result: false},
	}

	for _, testCase := range testCases {
		if isPalindrome2(testCase.input) != testCase.result {
			t.Error("FAIL", testCase)
		}
	}

}

func BenchmarkInvertText(b *testing.B) {
	input := "MAMAM" // ana // MAMAM // leonardo
	for i := 0; i < b.N; i++ {
		invertText(input)
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	input := "MAMAM" // ana // MAMAM // leonardo
	for i := 0; i < b.N; i++ {
		isPalindrome(input)
	}
}

func BenchmarkIsPalindrome2(b *testing.B) {
	input := "MAMAM" // ana // MAMAM // leonardo
	for i := 0; i < b.N; i++ {
		isPalindrome2(input)
	}
}
