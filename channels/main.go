package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://eduardo.wronscki.dev",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://amazon.comst",
	}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
	}

	count := len(links)
	for i := 0; i < count; i++ {
		fmt.Println(<-c)
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		c <- l + ": might be down!"
		return
	}
	c <- l + ": is up!"
}
