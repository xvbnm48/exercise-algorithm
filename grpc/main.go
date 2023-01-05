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

	// check if sakura endo is my girlfriend then print "yes i love her" else if aruno nakanishi is my girlfriend then print "yes i love her" else print "i love sakura endo and aruno nakanishi"

	gf1 := "sakura endo"
	gf2 := "aruno nakanishi"
	var result string

	if result == gf1 {
		fmt.Println("yes i love her")
	} else if result == gf2 {
		fmt.Println("yes i love her")
	} else {
		fmt.Println("i love sakura endo and aruno nakanishi")
	}
}
