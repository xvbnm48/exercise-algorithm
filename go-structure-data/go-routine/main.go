package main

import (
	"fmt"
	"runtime"
)

func cetak(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println(i+1, message)
	}
}
func main() {
	runtime.GOMAXPROCS(2)

	go cetak(5, "halo")
	cetak(5, "apa kabar")

	var input string
	fmt.Scanln(&input)
}
