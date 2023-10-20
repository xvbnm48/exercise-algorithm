package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		kata := scanner.Text()
		fmt.Println(kata)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
