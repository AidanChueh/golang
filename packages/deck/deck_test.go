package deck

import (
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	d := New()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := New()
	deck.saveToFile("_decktesting")

	loadedDeck := NewFromFile("_decktesting")

	if len(loadedDeck) != 160 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("-decktesting")
}
