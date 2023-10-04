package main

import "fmt"

func main() {
	animals := []string{
		"sapi",
		"ular",
		"bebek",
		"anjing",
	}

	for i := 0; i < len(animals); i++ {
		if animals[i] == "anjing" {
			fmt.Println("finding", animals[i], "!")
		}
	}
}
