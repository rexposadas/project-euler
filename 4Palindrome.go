package main

import (
	"fmt"
	"strconv"
)

// http://projecteuler.net/problem=4

// A palindromic number reads the same both ways. 
// The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

func r1(s string) string {
	arr := []byte(s)
	length := len(s) - 1
	for i := length; i > -1; i-- {
		arr[length-i] = s[i]
	}
	return string(arr)
}

func IsPalindrome(word string) bool {
	r := r1(word)
	if word == r {
		return true
	}

	return false
}

func main() {

	var answer int
	var left int
	var right int

	for i := 555; i < 999; i++ {
		for j := 555; j <= 999; j++ {
			a := i * j

			if IsPalindrome(strconv.Itoa(a)) {

				if a > answer {
					answer = a
					left = i
					right = j
				}
			}
		}
	}
	fmt.Printf("Answer: %d => %d * %d \n", answer, left, right)

}
