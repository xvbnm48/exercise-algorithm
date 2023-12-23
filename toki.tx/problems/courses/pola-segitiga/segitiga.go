package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j <= n-i {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}

	// for i := 0; i < sisi; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

}
