package main

import (
	"Golang-URL-shrtnr/shortener"
	"testing"
)

func TestHello(t *testing.T) {
	url := "https://en.wikipedia.org/wiki/URL_shortener#Techniques"
	_, err := shortener.Url_shortener(url)
	if err != nil {
		t.Fatal(err)
	}
}
