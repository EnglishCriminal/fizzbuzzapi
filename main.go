package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const errParamInvalidType string = "each query parameter must be a 32 bit integer"
const errParamMultiple string = "each parameter may only be specified once in the query string"
const errParseQueryString string = "unable to parse query string"
const errStartParamRequired string = "end query parameter must be accompanied by start query parameter"
const errStartParamGreaterThanEndParam string = "end query parameter must be greater than start query parameter"

func FizzBuzzEndpoint(w http.ResponseWriter, req *http.Request) {
	separator := ","

	rangeMin, rangeMax, err := parseRange(req)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte("["))
	for i := rangeMin; i <= rangeMax; i++ {
		if i == rangeMax {
			separator = ""
		}
		w.Write([]byte(fmt.Sprintf("%s%s", fizzBuzz(i), separator)))
	}
	w.Write([]byte("]"))
}

func fizzBuzz(i int) (output string) {
	if i%3 == 0 {
		output = "Fizz"
	}
	if i%5 == 0 {
		output = fmt.Sprintf("%sBuzz", output)
	}
	if output == "" {
		output = strconv.Itoa(i)
	} else {
		output = fmt.Sprintf("\"%s\"", output)
	}
	return
}

func parseIntKey(queryParams map[string][]string, key string) (value int, err error) {
	var keys []string
	var ok bool

	keys, ok = queryParams[key]
	if ok {
		if len(keys) > 0 {
			if len(keys) > 1 {
				err = errors.New(errParamMultiple)
				return
			}

			value, err = strconv.Atoi(keys[0])
			if err != nil {
				err = errors.New(errParamInvalidType)
				return
			}
		}
	}
	return
}

func parseRange(req *http.Request) (rangeMin int, rangeMax int, err error) {
	var end int
	var start int

	rangeMin = 1
	rangeMax = 100

	queryParams := req.URL.Query()

	start, err = parseIntKey(queryParams, "start")
	if err != nil {
		return
	}
	end, err = parseIntKey(queryParams, "end")
	if err != nil {
		return
	}

	if start > 0 {
		rangeMin = start
	} else if end > 0 {
		err = errors.New(errStartParamRequired)
		return
	}

	if end > 0 {
		rangeMax = end
		if start > end {
			err = errors.New(errStartParamGreaterThanEndParam)
			return
		}
	} else if start > 0 {
		rangeMax = start
	}

	return
}

func main() {
	http.HandleFunc("/fizzbuzz", FizzBuzzEndpoint)
	log.Fatal(http.ListenAndServe(":80", nil))
}
