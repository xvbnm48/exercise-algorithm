package main

import (
	"fmt"
	"math"
)

func main() {
	// Buat array arr
	arr := []int{-2, 3, 1, -5}

	// Inisialisasi x dengan maksimal dari arr
	x := math.MaxInt32
	for _, a := range arr {
		if a > x {
			x = a
		}
	}

	// Hitung running sum
	for i := 0; i < len(arr); i++ {
		sum := x + arr[i]
		if sum < 1 {
			x = 1 - sum
		}
	}

	// Cetak x
	fmt.Println(x)
}
