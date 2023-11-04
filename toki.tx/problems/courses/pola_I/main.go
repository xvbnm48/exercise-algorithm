package main

import (
	"fmt"
)

func main() {
	var n, b int

	fmt.Scan(&n, &b)
	for i := 1; i <= n; i++ {
		if i%b == 0 {
			fmt.Print("*")
		} else {
			fmt.Print(i)
		}

		if i != n {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
