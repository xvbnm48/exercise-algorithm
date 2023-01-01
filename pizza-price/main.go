package main

import (
	"fmt"
	"math"
)

func main() {
	// Daftar harga pizza
	pizzas := []int{800, 850, 900}

	// Daftar harga topping
	toppings := []int{100, 150}

	// Jumlah uang yang tersedia
	x := 1000

	// Inisialisasi harga pesanan minimum
	minPrice := math.MaxInt32

	// Cari harga pesanan yang paling mendekati x
	for _, pizza := range pizzas {
		for _, topping1 := range toppings {
			for _, topping2 := range toppings {
				// Hitung harga pesanan
				price := pizza + topping1 + topping2

				// Bandingkan dengan harga pesanan minimum sebelumnya
				if price <= x && x-price < x-minPrice {
					minPrice = price
				}
			}
		}
	}

	// Cetak harga pesanan minimum
	fmt.Println(minPrice)
}
