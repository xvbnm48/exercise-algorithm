package main

import (
	"fmt"
	"sort"
)

type KthLargest struct {
	// The slice of numbers in the stream
	nums []int

	// The value of K
	k int
}

// The constructor accepts an integer array of initial numbers and an integer K
func Constructor(nums []int, k int) KthLargest {
	// Initialize the KthLargest instance with the given numbers and K
	this := KthLargest{nums: nums, k: k}

	// Sort the numbers in the slice in non-ascending order
	sort.Sort(sort.Reverse(sort.IntSlice(this.nums)))

	// Return the initialized instance
	return this
}

// The add function adds a new number to the stream and returns the Kth largest number
func (this *KthLargest) Add(num int) int {
	// Append the new number to the slice
	this.nums = append(this.nums, num)

	// Sort the numbers in the slice in non-ascending order
	sort.Sort(sort.Reverse(sort.IntSlice(this.nums)))

	// Return the Kth largest number
	return this.nums[this.k-1]
}

func main() {
	// Initialize the KthLargest instance with the given numbers and K
	kthLargest := Constructor([]int{4, 1, 3, 12, 7, 14}, 3)

	// Add a new number to the stream and print the Kth largest number
	fmt.Println(kthLargest.Add(3))  // Output: 4
	fmt.Println(kthLargest.Add(5))  // Output: 5
	fmt.Println(kthLargest.Add(10)) // Output: 5
	fmt.Println(kthLargest.Add(9))  // Output: 8
	fmt.Println(kthLargest.Add(4))  // Output: 8
}
