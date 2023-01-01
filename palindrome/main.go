package main

import (
	"fmt"
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 || len(s) <= numRows {
		return s
	}

	rows := make([]string, numRows)
	curRow := 0
	down := false

	for i := 0; i < len(s); i++ {
		rows[curRow] += string(s[i])

		if curRow == 0 || curRow == numRows-1 {
			down = !down
		}

		if down {
			curRow++
		} else {
			curRow--
		}
	}

	return strings.Join(rows, "")
}

func main() {
	s := "PAYPALISHIRING"
	numRows := 3
	fmt.Println(convert(s, numRows)) // Output: "PAHNAPLSIIGYIR"
}
