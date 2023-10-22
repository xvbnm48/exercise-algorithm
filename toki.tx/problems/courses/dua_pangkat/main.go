package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		b, err := strconv.Atoi(input)
		if err != nil {
			return
		}
		if b <= 0 {
			fmt.Println("bukan")
		} else if (b & (b - 1)) == 0 {
			fmt.Println("ya")
		} else {
			fmt.Println("bukan")
		}
	}
}
