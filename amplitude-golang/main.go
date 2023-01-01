package main

import (
	"fmt"
	"sort"
)

func main() {
	// Buat array A
	A := []int{-10, 10, 3, 4, 10, 5}

	// Urutkan array A
	sort.Ints(A)

	// Cari amplitudo minimum
	amplitudoMin := A[len(A)-1] - A[0]
	if amplitudoMin > A[len(A)-1]-A[3] {
		amplitudoMin = A[len(A)-1] - A[3]
	}
	if amplitudoMin > A[len(A)-1]-A[4] {
		amplitudoMin = A[len(A)-1] - A[4]
	}
	if amplitudoMin > A[len(A)-1]-A[5] {
		amplitudoMin = A[len(A)-1] - A[5]
	}

	// Cetak amplitudo minimum
	fmt.Println(amplitudoMin)
}
