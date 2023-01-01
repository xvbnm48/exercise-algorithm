package main

import "fmt"

func romanToInt(s string) int {
	m := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	prev := 0

	// Iterate over the string in reverse order
	for i := len(s) - 1; i >= 0; i-- {
		// Get the integer value of the current Roman numeral
		current := m[s[i]]

		// If the current value is greater than or equal to the previous value,
		// we add it to the result
		if current >= prev {
			result += current
		} else {
			// Otherwise, we subtract it from the result
			result -= current
		}

		// Save the current value for the next iteration
		prev = current
	}

	return result
}

func main() {
	s := "XIV"
	fmt.Println(romanToInt(s)) // Output: 1994``
}
