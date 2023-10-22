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
		hasil, err := strconv.Atoi(input)
		if err != nil {
			return
		}

		for i := hasil; i >= 1; i-- {
			if hasil%i == 0 {
				fmt.Println(i)
			}
		}
	}
}
