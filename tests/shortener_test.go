package main

import (
	"shortener"
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello World"
	got := shortener.Hello()
	if want != got {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}