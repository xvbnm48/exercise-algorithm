package main

import "fmt"

func main() {
	var N, num, max, min int
	fmt.Scan(&N)
	fmt.Scan(&num)

	max = num
	min = num

	for i := 1; i < N; i++ {
		fmt.Scan(&num)
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}

	fmt.Printf("%d %d \n", max, min)
}
