package deck

import "fmt"

type Deck []string

func NewDeck() Deck {
	cards := Deck{}

	cardSuites := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suite := range cardSuites {
		for _, cardValue := range cardValues {
			cards = append(cards, cardValue+" of "+suite)
		}
	}

	return cards

}

func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
