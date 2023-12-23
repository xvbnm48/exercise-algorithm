package main

import "fmt"

func main() {
	var input int
	fmt.Scanln(&input)

	for i := 1; i <= input; i++ {
		if i%10 == 0 {
			continue
		}
		if i == 42 {
			fmt.Println("ERROR")
			break
		}
		fmt.Println(i)
	}
}
