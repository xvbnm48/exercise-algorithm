package main

import "strings"

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	s = strings.ReplaceAll(s, " ", "")

	// membalikkan string
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	reversed := string(runes)
	return s == reversed
}

func main() {
	string1 := "katak"
	// string2 := "kata"

	if isPalindrome(string1) {
		println(string1, "adalah palindrome")
	} else {
		println(string1, "bukan palindrome")
	}
}
