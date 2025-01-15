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
