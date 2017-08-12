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

## Tests

There is one unit and one integration test, to run the tests issue the `go test` command from the repository directory.

## Acceptance Criteria
```
When I perform a GET request to the /fizzbuzz endpoint
Then an array is returned with the length of 100
When the array is read as 1 indexed
Then indexes that are multiples of three are replaced with “Fizz” instead of the number
And indexes that are multiples of five are replaced with “Buzz”
And indexes that are multiples of both three and five are replace with “FizzBuzz"
And all other array values are their respective values
```
