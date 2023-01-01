package main

func findinteger() {
	// create an array of all numbers from 1 to 100
	var numbers []int
	for i := 1; i <= 100; i++ {
		numbers = append(numbers, i)
	}

	// create a map to keep track of which numbers we have seen
	var seen = make(map[int]bool)

	// loop over the given array and mark each number as seen
	for _, num := range givenArray {
		seen[num] = true
	}

	// loop over the numbers from 1 to 100 and find the missing numbers
	for _, num := range numbers {
		if !seen[num] {
			// num is missing
		}
	}

}
