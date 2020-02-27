package main

import "fmt"

func main() {
	fibonacci(100)
	fibonaccir(100, 0, 1)
}

func fibonacci(max int) {
	var n1 int = 0
	var n2 int = 1
	for i := n1 + n2; i < max; {
		fmt.Println(i)
		i = n1 + n2
		n1 = n2
		n2 = i
	}
}

func fibonaccir(max, n1, n2 int) {
	// readability
	fmt.Println(n2)

	current := n1 + n2
	if current > max {
		return
	}

	n1 = n2
	n2 = current

	fibonaccir(max, n1, n2)
}

// How did the logic work when I printed n2?
// How else could I have implemented fibonacci recursively
