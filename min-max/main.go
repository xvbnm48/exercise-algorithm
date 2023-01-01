package main

import (
	"fmt"
	"sort"
)

func minMaxSum(arr []int) (int, int) {
	// Urutkan slice
	sort.Ints(arr)

	// Hitung minimum dan maksimum jumlah
	minSum := arr[0] + arr[1] + arr[2] + arr[3]
	maxSum := arr[1] + arr[2] + arr[3] + arr[4]

	return minSum, maxSum
}

func main() {
	// Input
	minSum, maxSum := minMaxSum([]int{1, 2, 3, 4, 5})
	fmt.Println(minSum) // Output: 10
	fmt.Println(maxSum) // Output: 14

}
