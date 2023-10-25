package main

import "fmt"

func binarySearch(list []int, item int) any {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := low + high
		guess := list[mid]

		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return nil
}

func main() {
	list := []int{1, 3, 5, 7, 9}

	result := binarySearch(list, 3)
	if result == nil {
		fmt.Println("none")
	} else {
		fmt.Println(result)
	}
}
