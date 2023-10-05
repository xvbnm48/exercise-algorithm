package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	fmt.Println(maxCandies)
	for _, permen := range candies {
		fmt.Println("permen", permen)
		if permen > maxCandies {
			maxCandies = permen
			fmt.Println("max candies:", maxCandies)
		}
	}

	result := make([]bool, len(candies))

	for i, permen := range candies {
		fmt.Println("isi max candies:", maxCandies)
		fmt.Println("isi permen + extra candies:", permen+extraCandies)
		result[i] = permen+extraCandies >= maxCandies
	}
	return result
}

func main() {
	input := []int{
		2,
		3,
		5,
		1,
		3,
	}
	hasil := kidsWithCandies(input, 3)
	fmt.Println(hasil)
}
