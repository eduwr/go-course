package main

import (
	deck "cards/packages"
)

func main() {
	cards := deck.Deck{
		newCard(),
		"Ace of diamonds",
	}
	cards = append(cards, "Six of Spades")

	cards.Print()

}

func newCard() string {
	return "Five of Diamonds"
}
