package main

import (
	"fmt"
)

func main() {
	// String s
	s := "xabcdey"

	// String x
	x := "ab*de"

	// Inisialisasi index
	index := -1

	// Cari index pertama dari x di s
	for i := 0; i < len(s); i++ {
		if s[i] == x[0] {
			for j := i; j < i+len(x); j++ {
				if s[j] == x[j-i] || x[j-i] == '*' {
					continue
				} else {
					break
				}
			}

			if s[i:i+len(x)] == x {
				index = i
				break
			}
		}
	}

	// Cetak index
	fmt.Println(index)
}
