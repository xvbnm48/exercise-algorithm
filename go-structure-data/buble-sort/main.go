package main

import "fmt"

func Bubble(arr []int) {
	for i := 0; i < len(arr); i++ {
		// this j loop is for comparing the elements
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				// if arr index j is greater than arr index j+1, swap them
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	Bubble(arr)
	fmt.Println(arr)
}
