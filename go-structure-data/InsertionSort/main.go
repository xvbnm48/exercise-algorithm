package main

import "fmt"

func InsertionSort(arr []int) {
	var i, j int
	for i = 1; i < len(arr); i++ {
		for j = 0; j < i; j++ {
			if arr[j] > arr[i] {
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	InsertionSort(arr)
	fmt.Println(arr)
}
