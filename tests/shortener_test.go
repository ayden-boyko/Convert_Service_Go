package main

import (
	P "Golang-URL-shrtnr/pkg"
	"testing"
)

// TestHello tests the Url_shortener function to ensure it correctly processes
// a given URL and does not return an error. If an error occurs, the test fails.

func TestHello(t *testing.T) {
	url := "https://en.wikipedia.org/wiki/URL_shortener#Techniques"
	id, base62_id, _, err := P.Url_shortener(url)
	if err != nil {
		t.Fatal(err)
	}
	transformed_id, err := P.FromBase62(base62_id)
	if err != nil {
		t.Fatal(err)
	}
	if id != transformed_id {
		t.Fatalf("expected %d, got %d", id, transformed_id)
	}
}
