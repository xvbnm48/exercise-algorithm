package main

import "fmt"

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != j {
				nums[i], nums[j] = nums[j], nums[i]
			}
			j++
		}
	}
	return j
}

func main() {
	// Test case 1
	nums := []int{2, 2, 3,3,3, 4}
	result := removeElement(nums, 3)
	// Expected result: 3
	fmt.Println(result)
}
