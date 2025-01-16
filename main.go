package main

// cause path is github, not some local path
import (
	shortener "Golang-URL-shrtnr/pkg"
)

func main() {
	url := "https://en.wikipedia.org/wiki/URL_shortener#Techniques"
	println(shortener.Url_shortener("www.gourl.com", url))
}
