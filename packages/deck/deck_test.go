package deck

import (
	"os"
	"testing"
)

// %v means print anything no matter the type
// coding for test stability

// writing code for test stability
// write test accordingly
func TestNew(t *testing.T) {
	d := New()

	if len(d) != len(cardSuits)*len(cardValues) {
		t.Errorf("Expected deck length of %v, but got %v", len(cardSuits)*len(cardValues), len(d))
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := cardName(suit, value)
			if !d.contains(card) {
				t.Errorf("Expected %v, but not found", card)
			}
		}
	}
}

func TestSaveToDeckAndNewFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := New()
	deck.saveToFile("_decktesting")

	loadedDeck := newFromFile("_decktesting")

	if len(loadedDeck) != 160 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("-decktesting")
}
