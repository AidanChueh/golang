package deck

import (
	"os"
	"testing"
)

func testDuplicate(t *testing.T, d Deck) {
	for i, card1 := range d {
		for j, card2 := range d {
			if j != i {
				if card1 == card2 {
					t.Errorf("Duplicate card of %v", card1)
				}
			}
		}
	}
}

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

	testDuplicate(t, d)
}

// Initially, we thought we could not test random, but think: What can I do?
// test that there are between 1-52 cards
// test that each card is valid
// leverage that new works

func TestRandom(t *testing.T) {
	completeDeck := New()
	randomDeck := Random()

	if len(randomDeck) > 52 {
		t.Errorf("Expected deck of less than or equal to 52, but got deck of greater than 52")
	}

	if len(randomDeck) < 1 {
		t.Errorf("Expected deck of greater than or equal to 1, but got deck of less than 1")
	}

	for _, card := range randomDeck {
		if !completeDeck.contains(card) {
			t.Errorf("Invalid card of %v", card)
		}
	}

	testDuplicate(t, randomDeck)
}

// func TestDeal(t *testing.T) {
// 	d1 := New()
// 	hand, rest := d1.Deal(5)
// 	d2 := append(hand, rest)

// 	if !dealer.Equals(d1, d2) {
// 		t.Errorf("Expected deck of %v, but got deck of %v", d1, d2)
// 	}

// }

func TestShuffle(t *testing.T) {
	d := Random()
	d.Print()
	d2 := d
	d2.Print()
	d.Shuffle()
	d.Print()
	d2.Print()

	if d.Equals(d2) {
		t.Errorf("Expected shuffled deck, but got same deck")
	}

}

// ----------
// Not in use
// ----------
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
