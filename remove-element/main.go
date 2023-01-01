package main

import "fmt"

func removeElement(nums []int, val int) int {
	// Use two pointers to keep track of the array
	i, j := 0, 0
	for j < len(nums) {
		// If the current element is not a duplicate, add it to the result array
		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}
		j++
	}

	// Return the number of unique elements in the array
	return i
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	val := 3
	fmt.Println(removeElement(nums, val)) // Output: 2
}
