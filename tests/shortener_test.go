package main

import (
	I "Golang-URL-shrtnr/internal"
	S "Golang-URL-shrtnr/shortener"
	"testing"
)

// TestHello tests the Url_shortener function to ensure it correctly processes
// a given URL and does not return an error. If an error occurs, the test fails.

func TestHello(t *testing.T) {
	url := "https://en.wikipedia.org/wiki/URL_shortener#Techniques"
	id, base62_id, _, err := S.Url_shortener(url)
	if err != nil {
		t.Fatal(err)
	}
	transformed_id, err := I.FromBase62(base62_id)
	if err != nil {
		t.Fatal(err)
	}
	if id != transformed_id {
		t.Fatalf("expected %d, got %d", id, transformed_id)
	}
}
