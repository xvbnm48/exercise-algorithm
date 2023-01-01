// package main

// import (
//     "fmt"
// )

// func birthdayCakeCandles(candles []int) int {
//     // Find the maximum height of the candles
//     maxHeight := 0
//     for i := 0; i < len(candles); i++ {
//         if candles[i] > maxHeight {
//             maxHeight = candles[i]
//         }
//     }

//     // Count the number of candles with the maximum height
//     count := 0
//     for i := 0; i < len(candles); i++ {
//         if candles[i] == maxHeight {
//             count++
//         }
//     }

//     return count
// }

// func main() {
//     // Prompt the user for the number of candles
//     fmt.Print("Enter the number of candles: ")

//     // Read the number of candles from the user
//     var n int
//     fmt.Scanf("%d", &n)

//     // Create an array to store the candle heights
//     candles := make([]int, n)

//     // Prompt the user for the candle heights
//     fmt.Println("Enter the candle heights:")

//     // Read the candle heights from the user
//     for i := 0; i < n; i++ {
//         fmt.Scanf("%d", &candles[i])
//     }

//     // Print the number of tallest candles
//     fmt.Printf("The number of tallest candles is: %d\n", birthdayCakeCandles(candles))
// }

package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	start := 0
	end := 0

	for i := 0; i < len(s); i++ {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		len := max(len1, len2)

		if len > end-start {
			start = i - (len-1)/2
			end = i + len/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}

	return right - left - 1
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}
