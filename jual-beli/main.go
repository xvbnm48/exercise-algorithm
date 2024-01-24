// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// )

// // Kumpulan produk yang tersedia
// var products = []struct {
// 	name  string
// 	price int
// }{
// 	{"Produk 1", 10000},
// 	{"Produk 2", 20000},
// 	{"Produk 3", 30000},
// 	{"Produk 4", 40000},
// 	{"Produk 5", 50000},
// }

// // Fungsi untuk menampilkan daftar produk yang tersedia
// func DisplayProducts() {
// 	fmt.Println("Daftar produk yang tersedia:")
// 	for i, product := range products {
// 		fmt.Printf("%d. %s (Rp %d)\n", i+1, product.name, product.price)
// 	}
// }

// func main() {
// 	reader := bufio.NewReader(os.Stdin)

// 	// Tampilkan daftar produk yang tersedia
// 	DisplayProducts()

// 	// Minta input dari user untuk memilih produk yang akan dibeli
// 	var productIndex int
// 	fmt.Print("Pilih produk yang akan dibeli (1-5): ")
// 	fmt.Scanln(&productIndex)
// 	productIndex-- // Koreksi indeks agar sesuai dengan indeks array

// 	// Tentukan produk yang dipilih user
// 	product := products[productIndex]

// 	// Minta input dari user untuk memasukkan jumlah uang yang dibayarkan
// 	var payment int
// 	fmt.Print("Masukkan jumlah uang yang dibayarkan: ")
// 	paymentInput, _ := reader.ReadString('\n')
// 	payment, _ = strconv.Atoi(paymentInput[:len(paymentInput)-1])

// 	// Hitung kembalian
// 	change := payment - product.price

//		// Tampilkan informasi produk yang dibeli, jumlah uang yang dibayarkan, dan kembalian
//		fmt.Println("Anda membeli", product.name)
//		fmt.Println("Jumlah uang yang dibayarkan:", payment)
//		fmt.Println("Kembalian:", change)
//	}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Kumpulan produk yang tersedia
var products = []struct {
	name  string
	price int
}{
	{"Produk 1", 10000},
	{"Produk 2", 20000},
	{"Produk 3", 30000},
	{"Produk 4", 40000},
	{"Produk 5", 50000},
}

// Fungsi untuk menampilkan daftar produk yang tersedia
func DisplayProducts() {
	fmt.Println("Daftar produk yang tersedia:")
	for i, product := range products {
		fmt.Printf("%d. %s (Rp %d)\n", i+1, product.name, product.price)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Tampilkan daftar produk yang tersedia
	DisplayProducts()

	// Minta input dari user untuk memilih produk yang akan dibeli
	var productIndex int
	fmt.Print("Pilih produk yang akan dibeli (1-5): ")
	_, err := fmt.Scanln(&productIndex)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	productIndex-- // Koreksi indeks agar sesuai dengan indeks array

	// Pastikan indeks produk yang dimasukkan valid
	if productIndex < 0 || productIndex >= len(products) {
		fmt.Println("Pilihan produk tidak valid.")
		return
	}

	// Tentukan produk yang dipilih user
	product := products[productIndex]

	// Minta input dari user untuk memasukkan jumlah uang yang dibayarkan
	var payment int
	fmt.Print("Masukkan jumlah uang yang dibayarkan: ")
	paymentInput, _ := reader.ReadString('\n')
	// Membersihkan karakter newline dan carriage return
	paymentInput = strings.TrimSuffix(paymentInput, "\n")
	paymentInput = strings.TrimSuffix(paymentInput, "\r")

	payment, err = strconv.Atoi(paymentInput)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Pastikan jumlah uang yang dimasukkan valid
	if payment < product.price {
		fmt.Println("Jumlah uang yang dimasukkan tidak cukup untuk membeli produk ini.")
		return
	}

	// Hitung kembalian
	change := payment - product.price

	// Tampilkan informasi produk yang dibeli, jumlah uang yang dibayarkan, dan kembalian
	fmt.Println("Anda membeli", product.name)
	fmt.Println("Jumlah uang yang dibayarkan:", payment)
	fmt.Println("Kembalian:", change)
}
