package main

import "fmt"

// search nemo
func main() {
	animals := []string{
		"cow",
		"chicken",
		"nemo",
		"snake",
	}

	for i := 0; i < len(animals); i++ {
		if animals[i] == "nemo" {
			fmt.Println("finding", animals[i])
		}
	}
}
