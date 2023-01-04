package main

import "fmt"

func main() {
	fmt.Println("sakura endo")

	// 1. 1から100までの数字を出力する
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}

	var a string = "aruno nakanishi"
	for s := 1; s <= 5; s++ {
		fmt.Println("i love", a)
	}
}
