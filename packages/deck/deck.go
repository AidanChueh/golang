package deck

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Deck represents a deck of cards
type Deck []string

// Declaring the variables
var cardSuits []string
var cardValues []string

// Initializing the variables
func init() {
	cardSuits = []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues = []string{"Ace", "Two", "Three", "Four", "Five"}
}

// Implement a function to return size of deck

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
// avoid hard coding, such as writing "Ace of Spades"
// if it is an early exit condition, do it as soon as possible
func Random() Deck {
	cards := Deck{}

	if len(cardSuits) == 0 || len(cardValues) == 0 {
		return cards
	}

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	numCards := r.Intn(10)

	for i := 0; i <= numCards; i++ {
		cards = append(cards, cardName(cardSuits[0], cardValues[0]))
	}
	return cards
}

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

func cardName(suit string, value string) string {
	return value + " of " + suit
}

// Print prints the deck
func (d Deck) Print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deal deals the deck
func (d Deck) Deal(handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]
}

func (d Deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d Deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Shuffle shuffles the deck
func (d Deck) Shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d Deck) contains(card string) bool {
	for _, c := range d {
		if c == card {
			return true
		}
	}

	return false
}
