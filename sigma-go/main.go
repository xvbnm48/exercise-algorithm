package main

import "fmt"

func main() {
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		panic(err)
	}
	if err != nil {
		return
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += 3 * i
	}
	fmt.Println(sum)

}
