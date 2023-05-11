package main

import "fmt"

func numOfStage(angka int) int {
	langkah := 0

	for angka > 0 {
		if angka%2 == 0 {
			angka /= 2
		} else {
			angka -= 1
		}
		langkah++
	}
	return langkah
}

func main() {
	angka := 90
	fmt.Printf("dari angka %v, ini memeiliki %v langkah pengurangan untuk sampai menjadi 0\n", angka, numOfStage(angka))
}
