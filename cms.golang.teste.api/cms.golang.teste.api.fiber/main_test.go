package main_test

// import (
// 	"fmt"
// 	"testing"
// )

// // go test
// // go test -v
// // go test -bench=.
// // go test -bench=Add
// go test -run TestUserrepoGetAll -v
// go test -run=XXX -bench . -benchmem

// func IntMin(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// func TestIntMinTableDriven(t *testing.T) {
// 	var tests = []struct {
// 		a, b int
// 		want int
// 	}{
// 		{0, 1, 0},
// 		{1, 0, 0},
// 		{2, -2, -2},
// 		{0, -1, -1},
// 		{-1, 0, -1},
// 	}
// 	for _, tt := range tests {
// 		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
// 		t.Run(testname, func(t *testing.T) {
// 			ans := IntMin(tt.a, tt.b)
// 			if ans != tt.want {
// 				t.Errorf("got %d, want %d", ans, tt.want)
// 			}
// 		})
// 	}
// }

// func BenchmarkIntMin(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		IntMin(1, 2)
// 	}
// }

// func TestAdd(t *testing.T) {
// 	got := Add(4, 6)
// 	want := 10
// 	if got != want {
// 		t.Errorf("got %q, wanted %q", got, want)
// 	}
// }

// // arg1 means argument 1 and arg2 means argument 2, and the expected stands for the 'result we expect'
// type addTest struct {
// 	arg1, arg2, expected int
// }

// var addTests = []addTest{
// 	addTest{2, 3, 5},
// 	addTest{4, 8, 12},
// 	addTest{6, 9, 15},
// 	addTest{3, 10, 13},
// }

// func TestAdd2(t *testing.T) {
// 	for _, test := range addTests {
// 		if output := Add(test.arg1, test.arg2); output != test.expected {
// 			t.Errorf("Output %q not equal to expected %q", output, test.expected)
// 		}
// 	}
// }

// func BenchmarkAdd(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Add(4, 6)
// 	}
// }

// func ExampleAdd() {
// 	fmt.Println(Add(4, 6))
// 	// Output: 10
// }
