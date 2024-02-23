package main

import "fmt"

func binarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := low + (high-low)/2
		guess := list[mid]
		if guess == item {
			return mid
		} else if guess < item {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	list := []int{1, 2, 3, 4, 5}
	fmt.Println(binarySearch(list, 5))
}
