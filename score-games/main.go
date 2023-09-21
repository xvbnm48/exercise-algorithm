package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Points(games []string) int {
	points := 0

	for _, value := range games {
		skor := strings.Split(value, ":")
		x, y := skor[0], skor[1]
		xSkor, _ := strconv.Atoi(x)
		ySkor, _ := strconv.Atoi(y)

		if xSkor > ySkor {
			points += 3
		} else if xSkor == ySkor {
			points++
		}
	}

	return points
}
func main() {
	games := []string{"1:0", "2:0", "3:0"}
	fmt.Println(Points(games))
}
