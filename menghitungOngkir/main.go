package main

import (
	"fmt"
)

// this is a example of dynamic programming problem using study case of minimum cost shipping
func minimumCostShipping(costs [][]int, start int, end int) int {
	// Inisialisasi tabel untuk menyimpan hasil submasalah
	dp := make([][]int, len(costs))
	for i := range dp {
		dp[i] = make([]int, len(costs[i]))
	}

	// Hitung biaya minimum untuk mencapai setiap kota
	for i := end; i >= start; i-- {
		for j := end; j >= start; j-- {
			if i == end && j == end {
				dp[i][j] = costs[i][j]
			} else if i == end {
				dp[i][j] = costs[i][j] + dp[i][j+1]
			} else if j == end {
				dp[i][j] = costs[i][j] + dp[i+1][j]
			} else {
				dp[i][j] = costs[i][j] + min(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[start][start]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	costs := [][]int{
		{0, 10, 15, 20},
		{5, 0, 9, 10},
		{6, 13, 0, 12},
		{8, 8, 9, 0},
	}
	start := 0
	end := 3

	minCost := minimumCostShipping(costs, start, end)
	fmt.Printf("Biaya minimum pengiriman dari kota %d ke kota %d adalah %d", start, end, minCost)
}
