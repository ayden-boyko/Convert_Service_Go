package shortener

import (
	"encoding/binary"
	"fmt" // needed to print

	"github.com/google/uuid" // needed to gen uuid
)

const (
	base         uint64 = 62
	characterSet        = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// uint64ToBase62 takes a uint64 and converts it to a base62 number.
// This is a convenience function for generating a short string from a
// uint64, which is useful for generating short URLs.
func uint64ToBase62(uint64_id uint64) string {
	if uint64_id == 0 {
		return "0"
	}
	result := ""
	for uint64_id > 0 {
		r := uint64_id % 62
		result += string(characterSet[r])
		uint64_id /= 62
	}
	return result
}

// Url_shortener takes a URL and shortens it to a shorter URL. It
// uses a UUID to generate a random number, and then converts the
// number to a base62 number. The base62 number is then used to
// generate the short URL.
//
// The function returns the short URL as a string, as well as an error
// if anything goes wrong.
func Url_shortener(url string) (int, error) {

	id := uuid.New()

	uint64_id_fh := binary.BigEndian.Uint64(id[:8])
	// uint64_id_sh := binary.BigEndian.Uint64(id[8:])

	println(uint64_id_fh)

	fh_url := uint64ToBase62(uint64_id_fh)
	// sh_url := uint64ToBase62(uint64_id_sh)

	short_url := "www.gourl.com/" + fh_url // + sh_url

	return fmt.Println(short_url)
}
