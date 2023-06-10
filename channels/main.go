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

	for _, l := range links {
		checkLink(l)
	}
}

func checkLink(l string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l, " might be down!")
		return
	}

	fmt.Println(l, " is up!")
}
