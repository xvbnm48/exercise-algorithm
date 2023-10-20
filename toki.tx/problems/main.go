package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&A[i])
	}

	B := make([]int, K)
	for i := 0; i < K; i++ {
		fmt.Scan(&B[i])
	}

	patokTerjauh := make([]int, K)

	for i := 0; i < K; i++ {
		posisi := 0
		jarakDitempuh := 0

		for j := 0; j < N; j++ {
			jarakDitempuh += A[j]
			if jarakDitempuh >= B[i] {
				posisi = j + 1
				break
			}
		}

		patokTerjauh[i] = posisi
	}

	for i := 0; i < K; i++ {
		fmt.Println(patokTerjauh[i])
	}
}
