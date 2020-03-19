package deck

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// ---------
// Variables
// ---------

// Deck represents a deck of cards
type Deck []string

// Declaring variables
var cardSuits []string
var cardValues []string

// Initializing variables
func init() {
	cardSuits = []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
}

// --------------------
// Non Receiver Methods
// --------------------

// New creates and returns a new deck of cards
func New() Deck {
	cards := Deck{}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, cardName(suit, value))
		}
	}

	return cards
}

// Random creates and returns a random deck
// For now: Random deck is deck of random size
func Random() Deck {
	cards := Deck{}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	totalCards := len(cardSuits) * len(cardValues)
	numCards := r.Intn(totalCards) + 1

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			if cards.sizeOfDeck() < numCards {
				cards = append(cards, cardName(suit, value))
			}
		}
	}

	return cards
}

// Cardname returns the name of the card
func cardName(suit string, value string) string {
	return value + " of " + suit
}

// ----------------
// Receiver Methods
// ----------------

// Print prints the deck
func (d Deck) Print() {
	fmt.Println("--------")
	for _, card := range d {
		fmt.Println(card)
	}
	fmt.Println("--------")
}

// Deal returns the hand and the remaining deck
func (d Deck) Deal(handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

// Shuffle shuffles the deck
func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(d.sizeOfDeck() - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

// Equals determines if the decks are equal
func (d Deck) Equals(d2 Deck) bool {
	if len(d) != len(d2) {
		return false
	}

	for i, v := range d {
		if v != d2[i] {
			return false
		}
	}

	return true
}

// SizeOfDeck returns the size of the deck
func (d Deck) sizeOfDeck() int {
	return len(d)
}

// Contains checks if the deck contains the card
func (d Deck) contains(card string) bool {
	for _, c := range d {
		if c == card {
			return true
		}
	}

	return false
}

// ----------
// Not in use
// ----------

func newFromFile(filename string) Deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to New()
		// Option #2 - log the error and entirely quite the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return Deck(s)
}

func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d Deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}
