package pkg

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"
)

const (
	base         uint64 = 62
	characterSet        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// Uint64ToBase62 takes a uint64 and returns a string representing it in base 62.
// The base 62 string is used to shorten URLs.
func Uint64ToBase62(uint64_id uint64) (string, error) {
	// check that the input is a uint64
	if reflect.TypeOf(uint64_id).Kind() != reflect.Uint64 {
		return "", fmt.Errorf("cannot convert %d to base62", uint64_id)
	}
	// handle the edge case of 0
	if uint64_id == 0 {
		return "0", nil
	}
	// loop until uint64_id is 0
	result := ""
	for uint64_id > 0 {
		// get the remainder of uint64_id divided by 62
		r := uint64_id % 62
		// add the corresponding character of the characterSet to the result
		result = string(characterSet[r]) + result
		// divide uint64_id by 62
		uint64_id /= 62
	}
	// return the result
	return result, nil
}

// FromBase62 takes a base62 string and returns the uint64 it represents.
// It decodes the string to retrieve the key that was used to generate the
// shortened URL.
func FromBase62(encoded string) (uint64, error) {
	if reflect.TypeOf(encoded).Kind() != reflect.String {
		return 0, fmt.Errorf("cannot convert %s to uint64", encoded)
	}
	var val uint64
	for index, char := range encoded {
		pow := len(encoded) - (index + 1)
		pos := strings.IndexRune(characterSet, char)
		if pos == -1 {
			return 0, errors.New("invalid character: " + string(char))
		}

		val += uint64(pos) * uint64(math.Pow(float64(base), float64(pow)))
	}

	return val, nil
}
