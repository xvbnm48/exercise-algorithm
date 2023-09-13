package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := getStdin()
	n, _ := strconv.Atoi(lines[0])
	greenness := make([]int, n)

	// Mengisi nilai tingkat kehijauan dari input
	greennessStr := strings.Split(lines[1], " ")
	for i := 0; i < n; i++ {
		greenness[i], _ = strconv.Atoi(greennessStr[i])
	}

	// Menghitung "greenness on the other side" untuk setiap taman
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	leftMax[0] = greenness[0]
	rightMax[n-1] = greenness[n-1]

	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], greenness[i])
	}

	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], greenness[i])
	}

	result := make([]int, n)

	for i := 0; i < n; i++ {
		if i == 0 {
			result[i] = rightMax[i+1]
		} else if i == n-1 {
			result[i] = leftMax[i-1]
		} else {
			result[i] = max(leftMax[i-1], rightMax[i+1])
		}
	}

	// Mengembalikan hasil
	for _, v := range result {
		fmt.Println(v)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getStdin() []string {
	stdin, _ := ioutil.ReadAll(os.Stdin)
	return strings.Split(strings.TrimRight(string(stdin), "\n"), "\n")
}
