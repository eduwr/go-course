package deck

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := NewDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d\n", len(d))
	}

	firstCard := "Ace of Spades"
	lastCard := "King of Clubs"

	if d[0] != firstCard {
		t.Errorf("Expected first card to be `%s` but got `%s`\n", firstCard, d[0])
	}

	if d[51] != lastCard {
		t.Errorf("Expected last card to be `%s` but got `%s`\n", lastCard, d[51])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	tempFileName := "_decktesting"
	os.Remove(tempFileName)

	d := NewDeck()
	d.SaveToFile(tempFileName)

	loadedDeck := NewDeckFromFile(tempFileName)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d\n", len(d))
	}

	os.Remove(tempFileName)

}
