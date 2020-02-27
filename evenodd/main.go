package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Println(num, " is even")
		} else {
			fmt.Println(num, " is odd")
		}
	}
}

// Project Approach
// I first create a slice of int from 0-10 and used a for loop to iterate through the slice and print out each value
// Then I created the conditionals of even and odd and printed out the corresponding statement
// I ran into the challenge of printing an int and a string and I realized you had to print both values separately and not together
