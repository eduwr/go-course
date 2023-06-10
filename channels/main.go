package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://eduardo.wronscki.dev",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, l := range links {
		go checkLink(l, c)
	}

	for link := range c {
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(link)
	}
}

func checkLink(l string, c chan string) {
	_, err := http.Get(l)
	if err != nil {
		fmt.Println(l + ": might be down!")
		c <- l
		return
	}
	fmt.Println(l + ": is up!")
	c <- l
}
