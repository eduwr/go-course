package main

import (
	"fmt"
)

func main() {
	numbers := NewNumbers(11)

	for _, n := range numbers {
		if isEven(n) {
			fmt.Printf("%d is even", n)
		} else {
			fmt.Printf("%d is odd", n)
		}
		fmt.Println("")
	}

}

func NewNumbers(size int) []int {
	var numbersToTest []int

	for i := 0; i < size; i++ {
		numbersToTest = append(numbersToTest, i)
	}
	return numbersToTest
}

func isEven(n int) bool {
	return n%2 == 0
}
