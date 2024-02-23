package main

import "fmt"

// find a smallest number in a array
func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func selectionSort(arr []int) []int {
	newArr := make([]int, 0, len(arr))
	fmt.Println("panjang array:", len(arr))
	for len(arr) > 0 {
		smallestIndex := findSmallest(arr)
		fmt.Println("small index : ", smallestIndex)
		newArr = append(newArr, arr[smallestIndex])
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
		fmt.Println("arr", arr)
	}
	return newArr
}

func main() {
	arr := []int{5, 3, 6, 2, 10}
	fmt.Println(selectionSort(arr))
}
