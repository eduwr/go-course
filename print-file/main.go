package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args

	if len(args) > 2 {
		log.Fatal("Too much arguments received! expected: 1, found: " + fmt.Sprint((len(args) - 1)))
	}

	data, err := os.ReadFile(args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
