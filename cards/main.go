package main

import (
	deck "cards/packages"
	"fmt"
)

func main() {
	cards := deck.NewDeck()

	hand, remainingCards := deck.Deal(cards, 5)

	fmt.Println("Hand")
	fmt.Println("")
	hand.Print()
	fmt.Println("")
	fmt.Println("Remaining")
	fmt.Println("")
	remainingCards.Print()

}
