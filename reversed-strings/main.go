package main

import "fmt"

//func reversedStrings(input string) string {
//	runes := []rune(input)
//
//	// reverse the slice of runes
//	for i := 0; i < len(runes)-1; i++ {
//		// j ini adalah index terakhir dari slice runes
//		for j := 0; j < len(runes)-i-1; j++ {
//			// maksud ini adalah jika runes[j] lebih besar dari runes[j+1] j+1 itu adalah index setelah j (j+1) jadi index didepan runes[j] artinya ke arah kiri dari runes[j] yang dimana j itu adalah angka paling kanan dari slice runes
//			if runes[j] > runes[j+1] {
//				// jika runes j lebih besar dari runes j+1 maka runes j akan ditukar dengan runes j+1 yang artinya j akan pindah ke sisi kanan dan i akan pindah ke sisi kiri
//				runes[j], runes[j+1] = runes[j+1], runes[j]
//			}
//		}
//	}
//	reversed := string(runes)
//	return reversed
//}

func Solution(word string) string {
	runes := []rune(word)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {
	input := "anjayyyyy"
	result := Solution(input)
	fmt.Println(result)
}
