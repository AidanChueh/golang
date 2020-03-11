package dealer

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/AidanChueh/golang/packages/cards"
)

// Don't write the code and assume it works
// Test it to make sure it works

func TestCombineAndSplit(t *testing.T) {
	c1 := randomDeck()
	c2 := cards.NewDeck()
	d := NewDealer(c1, c2)
	fmt.Println("initial:", c1)

	s1, s2 := d.Split(c1)
	fmt.Println("split 1:", s1)
	fmt.Println("split 2:", s2)
	f := d.Combine(s1, s2)
	fmt.Println("final", f)

	if !(equals(c1, f)) {
		t.Errorf("Ha fool you failed!!!")
	}
}

func equals(d1 cards.Deck, d2 cards.Deck) bool {
	if len(d1) != len(d2) {
		return false
	}

	for i, v := range d1 {
		if v != d2[i] {
			return false
		}
	}

	return true
}

func randomDeck() cards.Deck {
	cards := cards.Deck{}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	numCards := r.Intn(10)
	fmt.Println("random deck:", numCards)
	for i := 0; i <= numCards; i++ {
		cards = append(cards, "Ace of Spades")
	}
	return cards
}

// NewDealer should have two decks
// func TestNewDealer(t *testing.T) {
// 	c1 := cards.NewDeck()
// 	c2 := cards.NewDeck()
// 	d := NewDealer(c1, c2)

// 	if len(d.deck1) != 16 {
// 		t.Errorf("Expected deck length of 16 for deck1, but got %v", len(d.deck1))
// 	}

// 	if len(d.deck2) != 16 {
// 		t.Errorf("Expected deck length of 16 for deck2, but got %v", len(d.deck2))
// 	}
// }

// Combine should combine two decks into one
// func TestCombine(t *testing.T){
// 	c1 := cards.NewDeck()
// 	c2 := cards.NewDeck()
// 	d := NewDealer(c1, c2)
// 	c = d.Combine()

// }
