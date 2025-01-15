package main

import (
	"Golang-URL-shrtnr/shortener"
	"testing"
)

// TestHello tests the Url_shortener function to ensure it correctly processes
// a given URL and does not return an error. If an error occurs, the test fails.

func TestHello(t *testing.T) {
	url := "https://en.wikipedia.org/wiki/URL_shortener#Techniques"
	_, err := shortener.Url_shortener(url)
	if err != nil {
		t.Fatal(err)
	}
}
