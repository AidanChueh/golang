package dealer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/AidanChueh/golang/packages/cards"
)

// Dealer represents a dealer of decks
type Dealer struct {
	deck1 cards.Deck
	deck2 cards.Deck
}

// NewDealer creates a new dealer
func NewDealer(decka cards.Deck, deckb cards.Deck) Dealer {
	decks := Dealer{deck1: decka, deck2: deckb}
	return decks
}

// Print prints the decks
func (d Dealer) Print() {
	d.deck1.Print()
	d.deck2.Print()
}

// Combine combines the decks
// Append function for arrays/slices: http://mussatto.github.io/golang/append/arrays/slices/2016/11/09/golang-append-two-arrays.html
func (d Dealer) Combine(deck1 cards.Deck, deck2 cards.Deck) cards.Deck {
	return append(deck1, deck2...)
}

// Split splits the deck
func (d Dealer) Split(deck cards.Deck) (cards.Deck, cards.Deck) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	splitPosition := r.Intn(len(deck) - 1)
	fmt.Println(splitPosition)
	return deck[:splitPosition], deck[splitPosition:]
}
