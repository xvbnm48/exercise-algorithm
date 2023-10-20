package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	total := 0
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		kata := scanner.Text()
		bilangan, err := strconv.Atoi(kata)
		if err != nil {
			fmt.Println(err)
			return
		}
		total += bilangan
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	fmt.Println(total)
}
