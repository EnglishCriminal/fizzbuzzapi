package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func FizzBuzzEndpoint(w http.ResponseWriter, req *http.Request) {
	separator := ","
	w.Write([]byte("["))
	for i := 1; i <= 100; i++ {
		if i == 100 {
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
	}else{
		output = fmt.Sprintf("\"%s\"", output)
	}
	return
}

func main() {
	http.HandleFunc("/fizzbuzz", FizzBuzzEndpoint)
	log.Fatal(http.ListenAndServe(":80", nil))
}
