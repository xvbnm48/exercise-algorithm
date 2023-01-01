// package main

// import (
// 	"fmt"
// )

// func main() {
// 	words := []string{"cat", "dada", "breath", "taking", "doll"}
// 	patterns := []string{"drctkla", "rdatlbohe", "cat"}

// 	for _, pattern := range patterns {
// 		for _, word := range words {
// 			if match(word, pattern) {
// 				fmt.Println("Kata pertama yang cocok dengan pola adalah:", word)
// 				break
// 			}
// 		}
// 	}
// }

// func match(word, pattern string) bool {
// if len(word) != len(pattern) {
// 	return false
// }

// for i := 0; i < len(word); i++ {
// 	if word[i] != pattern[i] {
// 		return false
// 	}
// }

// return true
// }
package main

import (
	"fmt"
)

func main() {
	words := []string{"cat", "dada", "breath", "taking", "doll"}
	patterns := []string{"drctkla", "rdatlbohe"}

	for _, pattern := range patterns {
		found := false
		for _, word := range words {
			if match(word, pattern) {
				fmt.Println("Kata pertama yang cocok dengan pola adalah:", word)
				found = true
				break
			}
		}
		if !found {
			fmt.Println("Tidak ada kata yang cocok dengan pola.")
		}
	}
}

func match(word, pattern string) bool {
	// Implementasikan logika untuk mengecek apakah kata cocok dengan pola
	// Contoh:
	// Kata "cat" cocok dengan pola "drctkla"
	// Kata "breath" cocok dengan pola "rdatlbohe"
	if len(word) != len(pattern) {
		return false
	}

	for i := 0; i < len(word); i++ {
		if word[i] != pattern[i] {
			return false
		}
	}

	return true
}
