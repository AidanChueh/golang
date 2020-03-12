package dealer

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/AidanChueh/golang/packages/deck"
)

// Dealer represents a dealer of decks
type Dealer struct {
	deck1 deck.Deck
	deck2 deck.Deck
}

// NewDealer creates a new dealer
func NewDealer(decka deck.Deck, deckb deck.Deck) Dealer {
	decks := Dealer{deck1: decka, deck2: deckb}
	return decks
}

// Print prints the decks
func (d Dealer) Print() {
	fmt.Println("------ Deck 1:")
	d.deck1.Print()
	fmt.Println("------ Deck 2:")
	d.deck2.Print()
}

// Combine combines the decks
// Append function for arrays/slices: http://mussatto.github.io/golang/append/arrays/slices/2016/11/09/golang-append-two-arrays.html
func (d Dealer) Combine() deck.Deck {
	return append(d.deck1, d.deck2...)
}

// Split splits the deck
func (d Dealer) Split(deck deck.Deck) (deck.Deck, deck.Deck) {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	splitPosition := r.Intn(len(deck) - 1)
	fmt.Println(splitPosition)
	return deck[:splitPosition], deck[splitPosition:]
}
