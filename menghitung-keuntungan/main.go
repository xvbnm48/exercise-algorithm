package main

import "fmt"

func getMaxProfit(stockPricesYesterday []int) int {
	minPrice := stockPricesYesterday[0]
	maxProfit := 0

	for _, currentPrice := range stockPricesYesterday {
		// Cek apakah ada keuntungan yang bisa didapat dari harga saham saat ini
		potentialProfit := currentPrice - minPrice
		if potentialProfit > maxProfit {
			maxProfit = potentialProfit
		}

		// Update harga terendah
		if currentPrice < minPrice {
			minPrice = currentPrice
		}
	}

	return maxProfit
}

func main() {
	stockPricesYesterday := []int{10, 10, 10, 10, 10, 10, 10, 2, 2, 2, 1, 0, 20}
	fmt.Println(getMaxProfit(stockPricesYesterday)) // Output: 500
}
