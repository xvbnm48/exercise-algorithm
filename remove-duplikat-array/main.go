package main

import (
	"fmt"
	"math"
)

func removeDuplicates(nums []int) int {
	last := math.MinInt32

	// Use two pointers to keep track of the array
	i, j := 0, 0
	for j < len(nums) {
		// If the current element is not a duplicate, add it to the result array
		if nums[j] != last {
			last = nums[j]
			nums[i] = nums[j]
			i++
		}
		j++
	}

	// Return the number of unique elements in the array
	return i
}

func main() {
	nums := []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}
	fmt.Println(removeDuplicates(nums)) // Output: 2
}
