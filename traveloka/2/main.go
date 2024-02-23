package main

import (
	"fmt"
)

func minimumGroups(predators []int32) int32 {
	// Fungsi untuk menghitung kedalaman spesies
	depthOfSpecies := func(i int32) int32 {
		depth := int32(0)
		for predators[i] != -1 {
			i = predators[i]
			depth++
		}
		return depth + 1
	}

	maxDepth := int32(0)
	for i, _ := range predators {
		depth := depthOfSpecies(int32(i))
		if depth > maxDepth {
			maxDepth = depth
		}
	}

	return maxDepth
}

func main() {
	predators := []int32{-1, 8, 6, 0, 7, 3, 8, 9, -1, 6}
	fmt.Println("Jumlah grup minimum yang dibutuhkan:", minimumGroups(predators))
}
