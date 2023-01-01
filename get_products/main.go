package main

import "fmt"

func get_products(numbers []int) []int {
	// Memeriksa apakah input list kosong
	if len(numbers) == 0 {
		return []int{}
	}

	// Menyiapkan variabel untuk menyimpan hasil produk
	result := make([]int, len(numbers))

	// Melakukan iterasi melalui setiap bilangan dalam daftar
	for i, number := range numbers {
		// Menyimpan produk dari semua bilangan kecuali bilangan pada indeks saat ini
		product := 1
		for j, num := range numbers {
			if i != j {
				product *= num
			}
		}
		result[i] = product
	}

	return result
}

func main() {
	// Input list
	numbers := []int{1, 7, 3, 4}

	// Memanggil fungsi get_products
	result := get_products(numbers)

	// Menampilkan hasil
	fmt.Println(result) // Output: [84 12 28 21]
}
