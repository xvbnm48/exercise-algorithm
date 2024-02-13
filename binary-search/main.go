package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

func binarySearch(list []int, item int) any {
	// low is 0
	low := 0
	// high is a length of list mines 1
	high := len(list) - 1

	// while low is less than or equal to high
	for low <= high {
		// mid is a sum of low and high
		mid := low + high
		// guess is a value of list index mid
		guess := list[mid]

		// if guess is equal to item
		if guess == item {
			return mid
		}
		// if guess is greater than item
		if guess > item {
			// high is mid - 1
			// this is a new search area
			high = mid - 1
		} else {
			// low is mid + 1
			low = mid + 1
		}
	}
	return nil
}

func main() {
	list := []int{1, 3, 5, 7, 9}
	sort.Ints(list)
	result := binarySearch(list, 3)
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	logger.Println("Aplikasi telah dimulai.")
	logger.Println("Ini adalah pesan log info.")
	logger.Println("Aplikasi telah selesai.")
	if result == nil {
		fmt.Println("none")
	} else {
		fmt.Println(result)
	}
}
