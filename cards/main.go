package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()}
	fmt.Println(cards)
	cards = append(cards, "Six of Spades")
}

func newCard() string {
	return "Five of Diamonds"
}
