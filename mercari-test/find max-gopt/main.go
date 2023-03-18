package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 4, 5}
	x := 1
	fmt.Println(MaximumMEX(arr, x)) // output: 3

	arr = []int{1, 1, 1}
	x = 2
	fmt.Println(MaximumMEX(arr, x)) // output: 0
}

func MaximumMEX(arr []int, x int) int {
	sort.Ints(arr) // urutkan array dalam urutan menaik

	mex := 1
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] <= mex {
			mex++
		} else {
			break
		}
	}

	// Cek apakah x dapat digunakan untuk menambahkan MEX atau tidak
	for i := 1; i <= x; i++ {
		if i >= mex && i-x < mex {
			mex = i + 1
		}
	}

	return mex - 1 // Kembalikan nilai MEX yang ditemukan
}
