package main

import (
	"fmt"

	"github.com/AidanChueh/golang/packages/cards"
)

func main() {
	c := cards.NewDeck()
	c.Print()
	c.Shuffle()
	fmt.Println(c.Deal(5))
}
