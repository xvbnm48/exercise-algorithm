// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )

// // Kumpulan hadiah yang tersedia pada sistem lucky spin
// var prizes = []string{
// 	"Hadiah 1",
// 	"Hadiah 2",
// 	"Hadiah 3",
// 	"Hadiah 4",
// 	"Hadiah 5",
// }

// // Fungsi untuk menentukan hadiah yang akan diberikan kepada user
// func DeterminePrize() string {
// 	// Tampilkan daftar hadiah yang tersedia
// 	fmt.Println("Daftar hadiah yang tersedia:")
// 	for i, prize := range prizes {
// 		fmt.Printf("%d. %s\n", i+1, prize)
// 	}

// 	// Minta input dari admin untuk memilih hadiah yang akan diberikan kepada user
// 	var prizeIndex int
// 	fmt.Print("Pilih hadiah yang akan diberikan (1-5): ")
// 	fmt.Scanln(&prizeIndex)
// 	prizeIndex-- // Koreksi indeks agar sesuai dengan indeks array

// 	// Kembalikan hadiah yang dipilih admin
// 	return prizes[prizeIndex]
// }

// func main() {
// 	// Seed generator random number dengan waktu sekarang
// 	rand.Seed(time.Now().Unix())

// 	// Tentukan hadiah yang akan diberikan kepada user
// 	prize := DeterminePrize()

// 	// Tampilkan hadiah yang didapatkan
// 	fmt.Println("Anda mendapatkan hadiah:", prize)
// }

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Kumpulan hadiah yang tersedia pada sistem lucky spin
var prizes = []string{
	"Hadiah 1",
	"Hadiah 2",
	"Hadiah 3",
	"Hadiah 4",
	"Hadiah 5",
}

// Fungsi untuk menentukan hadiah yang akan diberikan kepada user
func DeterminePrize() string {
	// Tampilkan daftar hadiah yang tersedia
	fmt.Println("Daftar hadiah yang tersedia:")
	for i, prize := range prizes {
		fmt.Printf("%d. %s\n", i+1, prize)
	}

	// Minta input dari admin untuk memilih hadiah yang akan diberikan kepada user
	var prizeIndex int
	fmt.Print("Pilih hadiah yang akan diberikan (1-5): ")
	fmt.Scanln(&prizeIndex)
	prizeIndex-- // Koreksi indeks agar sesuai dengan indeks array

	// Kembalikan hadiah yang dipilih admin
	return prizes[prizeIndex]
}

func main() {
	// Seed generator random number dengan waktu sekarang
	rand.Seed(time.Now().Unix())

	// Minta input dari user untuk melakukan deposit
	var deposit int
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan jumlah deposit: ")
	depositInput, _ := reader.ReadString('\n')
	deposit, _ = strconv.Atoi(depositInput[:len(depositInput)-1])

	// Tentukan hadiah yang akan diberikan kepada user
	prize := DeterminePrize()

	// Tampilkan informasi deposit dan hadiah yang didapatkan
	fmt.Printf("Anda telah melakukan deposit sebesar %d\n", deposit)
	fmt.Println("Anda mendapatkan hadiah:", prize)
}
