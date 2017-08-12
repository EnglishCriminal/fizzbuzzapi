package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const buzz = "\"Buzz\""
const fizz = "\"Fizz\""
const fizzbuzz = "\"FizzBuzz\""

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

func TestFizzBuzzEndpointRange(t *testing.T) {
	var endpointRangeTests = []struct {
		description        string
		queryString        string
		expectedStatusCode int
		expectedBody       string
	}{
		{"start and end parameters", "?start=1&end=3", 200, fmt.Sprintf("[1,2,%s]", fizz)},
		{"solo start parameter - int", "?start=1", 200, "[1]"},
		{"solo start parameter - fizz", "?start=3", 200, fmt.Sprintf("[%s]", fizz)},
		{"solo start parameter - buzz", "?start=5", 200, fmt.Sprintf("[%s]", buzz)},
		{"solo start parameter - fizzbuzz", "?start=15", 200, fmt.Sprintf("[%s]", fizzbuzz)},
		{"solo start parameter - invalid type - string", "?start=three", 422, errParamInvalidType},
		{"solo start parameter - invalid type - int64", "?start=999999999999999999999", 422, errParamInvalidType},
		{"solo end parameter", "?end=2", 422, errStartParamRequired},
		{"start and end parameters - invalid type - start", "?start=one&end=3", 422, errParamInvalidType},
		{"start and end parameters - invalid type - end", "?start=1&end=three", 422, errParamInvalidType},
		{"start and end parameters - start greater than end", "?start=9&end=1", 422, errStartParamGreaterThanEndParam},
		{"solo start paramater - multiples", "?start=1&start=2", 422, errParamMultiple},
		{"start and end paramaters - multiple start", "?start=1&start=2&end=100", 422, errParamMultiple},
		{"start and end paramaters - multiple end", "?start=1&end=100&end=2000", 422, errParamMultiple},
	}
	for _, tt := range endpointRangeTests {
		url := fmt.Sprintf("/fizzbuzz%s", tt.queryString)

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatalf("%s - error encounterd on GET request to %s", tt.description, url)
		}

		w := httptest.NewRecorder()
		FizzBuzzEndpoint(w, req)

		if w.Code != tt.expectedStatusCode {
			t.Fatalf("%s - expected (%d) status code but got (%d)", tt.description, tt.expectedStatusCode, w.Code)
		}
		if w.Body.String() != tt.expectedBody {
			t.Fatalf("%s - expected (%s) response body but got (%s)", tt.description, tt.expectedBody, w.Body.String())
		}
	}
}

func TestParseIntKey(t *testing.T) {
	var parseIntKeyTests = []struct {
		description       string
		valuesJson        string
		key               string
		expectedValue     int
		expectedErrString string
	}{
		{"solo start - invalid parameter", `{"start":["1"]}`, "start", 1, ""},
		{"start and end parameters", `{"start":["5"], "end":["3"]}`, "end", 3, ""},
		{"solo start - invalid parameter", `{"start": ["invalid"]}`, "start", 0, errParamInvalidType},
		{"start and end parameters - invalid start parameter", `{"start":["invalid"], "end":["3"]}`, "start", 0, errParamInvalidType},
		{"start and end parameters - invalid end parameter", `{"start":["5"], "end":["invalid"]}`, "end", 0, errParamInvalidType},
	}

	for _, tt := range parseIntKeyTests {
		var values map[string][]string
		err := json.Unmarshal([]byte(tt.valuesJson), &values)
		if err != nil {
			t.Errorf("%s - unable to unmarshal values JSON : %s", tt.description, tt.valuesJson)
		}

		actualValue, actualError := parseIntKey(values, tt.key)
		if actualValue != tt.expectedValue {
			t.Errorf("%s - parseIntKey(%v)[%s]: expected value %d, actual value %d", tt.description, tt.valuesJson, tt.key, tt.expectedValue, actualValue)
		}
		if actualError != nil {
			if actualError.Error() != tt.expectedErrString {
				t.Errorf("%s - parseIntKey(%v)[%s]: expected error %s, actual error %s", tt.description, tt.valuesJson, tt.key, tt.expectedErrString, actualError)
			}
		} else if tt.expectedErrString != "" {
			t.Errorf("%s - parseIntKey(%v)[%s]: expected error %s, got nothing", tt.description, tt.valuesJson, tt.key, tt.expectedErrString)
		}

	}
}

func TestParseRange(t *testing.T) {
	var parseRangeTests = []struct {
		description       string
		queryString       string
		expectedMin       int
		expectedMax       int
		expectedErrString string
	}{
		{"start and end parameters", "?start=1&end=3", 1, 3, ""},
		{"solo start parameter - int", "?start=1", 1, 1, ""},
		{"solo start parameter - invalid type - string", "?start=three", 1, 100, errParamInvalidType},
		{"solo end parameter", "?end=2", 1, 100, errStartParamRequired},
		{"start and end parameters - invalid type - start", "?start=one&end=3", 1, 100, errParamInvalidType},
		{"start and end parameters - invalid type - end", "?start=1&end=three", 1, 100, errParamInvalidType},
		{"start and end parameters - start greater than end", "?start=9&end=1", 9, 1, errStartParamGreaterThanEndParam},
	}

	for _, tt := range parseRangeTests {
		url := fmt.Sprintf("/fizzbuzz%s", tt.queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			t.Fatalf("%s - error encounterd on GET request to %s",tt.description, url)
		}
		actualMin, actualMax, actualErr := parseRange(req)
		if actualMin != tt.expectedMin {
			t.Errorf("%s - parseRange() with query string (%s) : expected rangeMin %d, actual rangeMin %d", tt.description, tt.queryString, tt.expectedMin, actualMin)
		}
		if actualMax != tt.expectedMax {
			t.Errorf("%s - parseRange() with query string (%s) : expected rangeMax %d, actual rangeMax %d", tt.description, tt.queryString, tt.expectedMax, actualMax)
		}
		if actualErr != nil {
			if actualErr.Error() != tt.expectedErrString {
				t.Errorf("%s - parseRange() with query string (%s) : expected error %s, actual error %s", tt.description, tt.queryString, tt.expectedErrString, actualErr)
			}
		} else if tt.expectedErrString != "" {
			t.Errorf("%s - parseRange() with query string (%s) : expected error %s, got nothing", tt.description, tt.queryString, tt.expectedErrString)
		}
	}
}

func TestFizzBuzz(t *testing.T) {

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
