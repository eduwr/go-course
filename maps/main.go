package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	printMap(colors)
}

// map is a reference type
func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Printf("%s: %s\n", key, value)
	}
}
