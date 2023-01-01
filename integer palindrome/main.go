package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	// Negative numbers are not palindromes
	if x < 0 {
		return false
	}

	// Convert the integer to a string
	s := strconv.Itoa(x)

	// Compare the first and last characters, second and second-to-last characters, and so on
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}

	return true
}

func main() {
	x := 12321
	fmt.Println(isPalindrome(x)) // Output: true

	x = 1234
	fmt.Println(isPalindrome(x)) // Output: false
}
