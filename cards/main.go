package main

import (
	"fmt"

	"github.com/AidanChueh/golang/packages/dealer"
	"github.com/AidanChueh/golang/packages/deck"
)

func main() {
	c := deck.NewDeck()
	c.Print()
	c.Shuffle()
	fmt.Println(c.Deal(5))
	c1 := deck.NewDeck()
	c2 := deck.NewDeck()
	d := dealer.NewDealer(c1, c2)
	d.Print()
	fmt.Println(d.Combine())
	fmt.Println(d.Split(d))
}
