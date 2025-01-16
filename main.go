package main

// cause path is github, not some local path
import (
	"Golang-URL-shrtnr/shortener"
)

func main() {
	url := "https://en.wikipedia.org/wiki/URL_shortener#Techniques"
	println(shortener.Url_shortener(url))
}
