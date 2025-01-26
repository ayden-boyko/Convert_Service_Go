package main

import (
	"encoding/binary"
	"testing"

	P "github.com/ayden-boyko/Convert_Service_Go/pkg"
	"github.com/google/uuid"
)

// TestHello tests the Url_shortener function to ensure it correctly processes
// a given URL and does not return an error. If an error occurs, the test fails.

func TestConvert(t *testing.T) {
	uuid := uuid.New()

	id := binary.BigEndian.Uint64(uuid[:8])
	// uint64_id_sh := binary.BigEndian.Uint64(id[8:])

	id_62, err := P.Uint64ToBase62(id)
	if err != nil {
		t.Fatal(err)
	}

	transformed_id, err := P.Base62ToUint64(id_62)
	if err != nil {
		t.Fatal(err)
	}
	if id != transformed_id {
		t.Fatalf("expected %d, got %d", id, transformed_id)
	}
}
