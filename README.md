## Synopsis

This is a code exercise to write a FizzBuzz API in Golang. This code was written with performance in mind, which is why no encoding is used, output is streamed, and `fmt.Sprintf()` is used in lieu of string concatenation.

## Requirements

- Create a FizzBuzz microservice which returns the FizzBuzz output as an array.
- This microservice should listen for HTTP on :80, no authentication is necessary.
- It should be a git repository with a Dockerfile and it should be able to be built as a single step docker build.

## Running in Docker

The repository contains a Dockerfile.

To build the docker image type the following command from the repository directory:

`docker build -t fizzbuzzapi .`

To run type the following command, this get a docker machine running on your local machine and will bind from port 80 of the docker machine to port 8000 on the host machine.

`docker run --publish 8000:80 --name fizzbuzzapi --rm fizzbuzzapi`


## Usage

After the service is running you can load the service by using the following URL

`http://localhost:8000/fizzbuzzapi`

By default, the service returns a range of 1..100. You can modify this range by supplying a `start` and `end` query parameter. The following URL will return the range of 1..10.  The value must be a 32 bit integer.

`http://localhost:8000/fizzbuzzapi?start=1&end=10`

To return a single value, one can either use the same integer for `start` and `end`, or supply the `start` value by itself.

`http://localhost:8000/fizzbuzzapi?start=15`

## Tests

There are three unit and two integration tests, to run the tests issue the `go test` command from the repository directory.

## Acceptance Criteria
```
When I perform a GET request to the /fizzbuzz endpoint
Then an array is returned with the length of 100
When the array is read as 1 indexed
Then indexes that are multiples of three are replaced with “Fizz” instead of the value of the index
And indexes that are multiples of five are replaced with “Buzz”
And indexes that are multiples of both three and five are replaced with “FizzBuzz"
And all other array values are the value of the index

When I perform a GET request to the /fizzbuzz endpoint with start and end query parameters
Then an array is returned with the length of the range between start and end
When the array is read as 1 indexed
Then indexes that are multiples of three are replaced with “Fizz” instead of the value of the index
And indexes that are multiples of five are replaced with “Buzz”
And indexes that are multiples of both three and five are replaced with “FizzBuzz"
And all other array values are the value of the index

When I perform a GET request to the /fizzbuzz endpoint with a start query parameter
Then an array is returned with the length of of one
And if the value is a multiple of three, it is replaced with “Fizz” instead of the value of the index
And if the value is a multiple of five, it is replaced with “Buzz”
And if the value is a multiple of both three and five, it is replaced with “FizzBuzz"
And if the value is not a multiple of either three or five, the value of the index is returned

When I perform a GET request to the /fizzbuzz endpoint with a query parameter that is not a 32 bit integer
Then an error is returned describing the type requirements

When I perform a GET request to the /fizzbuzz endpoint with a start query parameter that is greater than the end query parameter
Then an error is returned describing the value requirements

When I perform a GET request to the /fizzbuzz endpoint with a start query parameter that is greater than the end query parameter
Then an error is returned describing the requirement of the start query parameter

When I perform a GET request to the /fizzbuzz endpoint with multiple values for a single query parameter
Then an error is returned describing the requirement of a single value per key
```
