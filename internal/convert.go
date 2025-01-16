package internal

import (
	"errors"
	"math"
	"strings"
)

const (
	base         uint64 = 62
	characterSet        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// uint64ToBase62 takes a uint64 and converts it to a base62 number.
// This is a convenience function for generating a short string from a
// uint64, which is useful for generating short URLs.
func Uint64ToBase62(uint64_id uint64) string {
	if uint64_id == 0 {
		return "0"
	}
	result := ""
	for uint64_id > 0 {
		r := uint64_id % 62
		result = string(characterSet[r]) + result
		uint64_id /= 62
	}
	return result
}

func FromBase62(encoded string) (uint64, error) {
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
