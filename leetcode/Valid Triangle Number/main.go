package main

import "fmt"

func triangleNumber(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	count := 0

	// Loop pertama untuk elemen pertama dari triplet
	for side1Index := 0; side1Index < n-2; side1Index++ {
		// Loop kedua untuk elemen kedua dari triplet
		for side2Index := side1Index + 1; side2Index < n-1; side2Index++ {
			// Loop ketiga untuk elemen ketiga dari triplet
			for side3Index := side2Index + 1; side3Index < n; side3Index++ {
				// Memeriksa apakah triplet memenuhi syarat pembentukan segitiga
				if nums[side1Index]+nums[side2Index] > nums[side3Index] &&
					nums[side1Index]+nums[side3Index] > nums[side2Index] &&
					nums[side2Index]+nums[side3Index] > nums[side1Index] {
					// Jika memenuhi, tambahkan ke jumlah count
					count++
				}
			}
		}
	}

	return count
}

func main() {
	// Test case 1
	nums := []int{2, 2, 3, 4}
	result := triangleNumber(nums)
	fmt.Println(result)
}
