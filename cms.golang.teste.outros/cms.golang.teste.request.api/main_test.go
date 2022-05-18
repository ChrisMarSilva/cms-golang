package main_test

import (
	"io"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

// go test
// go test -v
// go test -bench=Add
// go test -run TestUserrepoGetAll -v

// go test -bench=.
// go test -bench=. -count 5 -run=^#
// go test -run=XXX -bench . -benchmem

func TestRoute(t *testing.T) {

	tests := []struct {
		description  string // description of the test case
		method       string // type path to test
		route        string // route path to test
		expectedCode int    // expected HTTP status code
		body         io.Reader
	}{
		{description: "get HTTP status 200", method: "GET", route: "http://localhost:8011/", expectedCode: 200, body: nil},
	}

	for _, test := range tests {
		req := httptest.NewRequest(test.method, test.route, test.body)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		resp := w.Result()
		if test.expectedCode != resp.StatusCode {
			body, _ := ioutil.ReadAll(resp.Body)
			println("body", string(body))
		}
	} // for _, test := range tests {

}

func BenchmarkGin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "http://localhost:7003/", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		resp := w.Result()
		if 200 != resp.StatusCode {
			body, _ := ioutil.ReadAll(resp.Body)
			println(string(body))
		}
	}
}

func BenchmarkFiber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "http://localhost:7004/", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		resp := w.Result()
		if 200 != resp.StatusCode {
			body, _ := ioutil.ReadAll(resp.Body)
			println(string(body))
		}
	}
}

func BenchmarkPython(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "http://localhost:8011/", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		resp := w.Result()
		if 200 != resp.StatusCode {
			body, _ := ioutil.ReadAll(resp.Body)
			println(string(body))
		}
	}
}

func BenchmarkDotNet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest("GET", "https://localhost:44332/WeatherForecast/ok", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		resp := w.Result()
		if 200 != resp.StatusCode {
			body, _ := ioutil.ReadAll(resp.Body)
			println(string(body))
		}
	}
}
