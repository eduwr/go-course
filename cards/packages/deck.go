package deck

import "fmt"

type Deck []string

func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
