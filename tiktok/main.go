package main

import "fmt"

func longestOnes(A []int, K int) int {
	lBound, rBound := 0, 0
	for i := 0; i < len(A); i++ {
		if A[i] == 0 {
			K--
		}
		if K < 0 {
			break
		}
		rBound++
	}
	maxWidth := rBound
	for i := rBound; i < len(A); i++ {
		if A[i] == 1 {
			curWidth := i - lBound + 1
			if curWidth > maxWidth {
				maxWidth = curWidth
			}
		} else {
			for A[lBound] != 0 {
				lBound++
			}
			lBound++
		}
	}
	return maxWidth
}
func main() {
	A := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	K := 2
	fmt.Println(longestOnes(A, K))

}
