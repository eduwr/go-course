package main

import (
	deck "cards/packages"
)

func main() {
	cards := deck.NewDeck()
	cards.Shuffle()
	cards.Print()
}
