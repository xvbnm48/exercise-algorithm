package main

import "fmt"

func containsDuplicate(nums []int) bool {
	// make seen for store value in slice of array bool
	seen := make(map[int]bool)

	// for loop for range of nums
	for _, num := range nums {
		fmt.Println("seen", seen)
		if seen[num] {
			return true
		}
		seen[num] = true
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(containsDuplicate(nums))
}
