package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEndpoint(t *testing.T) {
	req, err := http.NewRequest("GET", "/fizzbuzz", nil)
	if err != nil {
		t.Fatalf("Error encounterd on GET request to /fizzbuzz")
	}

	w := httptest.NewRecorder()
	FizzBuzzEndpoint(w, req)

	if w.Code != 200 {
		t.Fatalf("Expected 200 status code but got %d", w.Code)
	}
	expectedBody := "["
	for i := 1; i <= 100; i++ {
		expectedBody = fmt.Sprintf("%s%s,", expectedBody, fizzBuzz(i))
	}
	expectedBody = strings.TrimSuffix(expectedBody, ",")
	expectedBody = fmt.Sprintf("%s%s", expectedBody, "]")
	if w.Body.String() != expectedBody {
		t.Fatalf("Expected (%s) response body but got (%s)", expectedBody, w.Body.String())
	}

}

func TestFizzBuzz(t *testing.T) {
	const buzz = "\"Buzz\""
	const fizz = "\"Fizz\""
	const fizzbuzz = "\"FizzBuzz\""

	var fizzBuzzTests = []struct {
		i        int
		expected string
	}{
		{1, "1"},
		{2, "2"},
		{3, fizz},
		{4, "4"},
		{5, buzz},
		{6, fizz},
		{7, "7"},
		{8, "8"},
		{9, fizz},
		{10, buzz},
		{11, "11"},
		{15, fizzbuzz},
		{30, fizzbuzz},
		{33, fizz},
		{35, buzz},
		{50, buzz},
	}
	for _, tt := range fizzBuzzTests {
		actual := fizzBuzz(tt.i)
		if actual != tt.expected {
			t.Errorf("fizzBuzz(%d): expected %d, actual %d", tt.i, tt.expected, actual)
		}
	}

}
