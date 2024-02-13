package main

import "fmt"

func isValid(s string) bool {
	var stack []rune // Membuat tumpukan kosong menggunakan tipe rune

	for _, char := range s {
		if char == '(' || char == '{' || char == '[' {
			// Jika karakter buka kurung, masukkan ke dalam tumpukan
			stack = append(stack, char)
		} else {
			// Jika karakter tutup kurung, periksa apakah tumpukan kosong
			if len(stack) == 0 {
				return false
			}

			// Ambil elemen terakhir dari tumpukan
			top := stack[len(stack)-1]
			fmt.Println(" top element :", top)

			// Periksa apakah karakter tutup kurung sesuai dengan yang diharapkan
			if (char == ')' && top == '(') || (char == '}' && top == '{') || (char == ']' && top == '[') {
				// Pop elemen terakhir dari tumpukan
				stack = stack[:len(stack)-1]
			} else {
				return false // Tidak sesuai, string tidak valid
			}
		}
	}

	// String dianggap valid jika tumpukan kosong pada akhirnya
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("{}[][]"))
}
