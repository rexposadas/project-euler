package main

import "fmt"

// http://projecteuler.net/problem=1
// If we list all the natural numbers below 10 that are multiples of 3 or 5,
// we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.
func main() {
	var total int
	for i := 0; i < 1000; i++ {

		if i%3 == 0 {
			total += i
			continue
		}

		if i%5 == 0 {
			total += i
		}

	}

	fmt.Printf("The sum of all the multiples of 3 or 5 below 1000 was : %d\n", total)
}
