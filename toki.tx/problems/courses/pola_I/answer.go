package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	for i := 1; i <= N; i++ {
		if i%K == 0 {
			//fmt.Println("i nya ", i, "k nya:", K)
			fmt.Println("hasil modulus:", 6%K)
			fmt.Print("*")
		} else {
			fmt.Print(i)
		}

		if i != N {
			fmt.Print(" ")
		}
	}

	fmt.Println()
}
