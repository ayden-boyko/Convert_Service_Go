# Golang-URL-shrtnr

This is a simple URL shortener project written in Go.
It uses a UUID to generate a short string, which is then appended
to a base URL to form the shortened URL.

## How to use

1. Clone the repository
2. Run `go run main.go`
3. Open a web browser and navigate to `http://localhost:8080/`

## How it works

1. The program generates a UUID using the `github.com/google/uuid` package.
2. It takes the first 8 bytes of the UUID and converts them to a uint64.
3. It encodes the uint64 into a base62 string using the `pkg/convert.go` package.
4. It appends the base62 string to a base URL to form the shortened URL.
5. It returns the shortened URL.

## Testing

The program is tested using the `testing` package.
The test checks that the `Url_shortener` function correctly processes
a given URL and does not return an error.

## Packages used

The following packages are used:

- `github.com/google/uuid` for generating UUIDs
- `math` for calculating the base62 value
- `strings` for manipulating strings
- `errors` for returning errors
- `encoding/binary` for converting between types
