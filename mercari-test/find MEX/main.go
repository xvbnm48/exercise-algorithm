package main

import (
	"fmt"
	"sort"
)

func main() {
	// Buat array N
	N := []int{3, 4, 5}

	// Nilai X
	X := 1

	// Urutkan array N
	sort.Ints(N)

	// Inisialisasi MEX
	MEX := 0

	// Cari MEX
	for i := 0; i < len(N); i++ {
		if MEX == N[i] {
			MEX += X
		} else if MEX != N[i] {
			MEX = N[i]
			break
		}
	}

	// Cetak MEX
	fmt.Println(MEX)
}
