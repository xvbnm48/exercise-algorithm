package main

import "fmt"

func binarySearch(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := low + high
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
func binarysearchString(list []string, item string) int {
	low := 0
	high := len(list) - 1
	for low <= high {
		mid := low + high
		guess := list[mid]
		fmt.Println("guess: ", guess)
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
	//list := []int{1, 3, 5, 7, 9}
	nameList := []string{"sakura", "sasuke", "naruto", "hinata", "sai"}
	resultStr := binarysearchString(nameList, "sakura")
	if resultStr == -1 {
		println("none")
	} else {
		println(resultStr)
	}
	//result := binarySearch(list, 7)
	//if result == -1 {
	//	println("none")
	//} else {
	//	println(result)
	//}
}
