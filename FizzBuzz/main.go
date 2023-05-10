package main

import (
	"fmt"
	"strconv"
)

func fizbuzz(n int) {
	for i := 1; i < n; i++ {
		divBy3 := i%3 == 0
		divBy5 := i%5 == 0
		currStr := ""

		if divBy3 {
			currStr += "Fizz"
		} else if divBy5 {
			currStr += "Buzz"
		} else {
			currStr = strconv.Itoa(i)
		}
		fmt.Println(currStr)
	}
}

func main() {
	fizbuzz(100)
}
