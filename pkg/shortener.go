package pkg

import (
	"encoding/binary"
	// needed to print
	"github.com/google/uuid" // needed to gen uuid
)

// Url_shortener generates a shortened URL for a given input URL.
// It creates a UUID, converts part of it to a uint64, and then
// encodes that uint64 into a base62 string. The base62 string
// is appended to a base URL to form the shortened URL.
// Returns the uint64 ID, the base62 encoded string, the shortened
// URL, and an error if any occurs.

func Url_shortener(url string) (uint64, string, string, error) {

	uuid := uuid.New()

	id := binary.BigEndian.Uint64(uuid[:8])
	// uint64_id_sh := binary.BigEndian.Uint64(id[8:])

	id_62 := uint64ToBase62(id)
	// sh_url := uint64ToBase62(uint64_id_sh)

	short_url := "www.gourl.com/" + id_62 // + sh_url

	return id, id_62, short_url, nil
}
