package main

import (
	"fmt"

	"github.com/AidanChueh/golang/packages/dealer"
	"github.com/AidanChueh/golang/packages/deck"
)

func main() {
	c := deck.New()
	c.Print()
	c.Shuffle()
	fmt.Println(c.Deal(5))
	c1 := deck.New()
	c2 := deck.New()
	d := dealer.NewDealer(c1, c2)
	d.Print()
	fmt.Println(d.Combine())
	fmt.Println(d.Split(d))
}
