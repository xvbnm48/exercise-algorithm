package main

import "fmt"

func plusMinus(arr []int) {
	var pos, neg, zero float64
	length := float64(len(arr))

	for _, num := range arr {
		if num > 0 {
			pos++
		} else if num < 0 {
			neg++
		} else {
			zero++
		}
	}

	fmt.Printf("%.6f\n", pos/length)
	fmt.Printf("%.6f\n", neg/length)
	fmt.Printf("%.6f\n", zero/length)
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	plusMinus(arr)
	// plusMinus([]int{-4, 3, -9, 0, 4, 1}) // Output: 0.333333 0.333333 0.333333

}
