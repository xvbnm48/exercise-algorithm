package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())

	for i := 1; i <= input; i++ {
		if i%10 == 0 {
			continue
		}
		if i == 42 {
			fmt.Println("ERROR")
		}
		fmt.Println(i)
	}

}
