package main

import "fmt"

func main() {
	cards := []string{
		newCard(),
		"Ace of diamonds",
	}
	cards = append(cards, "Six of Spades")

	for i, card := range cards {
		fmt.Println(i, card)
	}

	count := len(cards)
	for i := 0; i < count; i++ {
		println(i, cards[i])
	}

}

func newCard() string {
	return "Five of Diamonds"
}
