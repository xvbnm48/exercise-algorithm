package main

import (
	"fmt"
)

func ways(total int32, k int32) int32 {
	const mod int32 = 1000000007
	dp := make([]int32, total+1)
	dp[0] = 1 // Hanya ada satu cara untuk mencapai total 0

	for i := int32(1); i <= k; i++ {
		for j := i; j <= total; j++ {
			dp[j] = (dp[j] + dp[j-i]) % mod
		}
	}

	return dp[total]
}

func main() {
	total := int32(8)
	k := int32(2)
	fmt.Println("Jumlah cara untuk mencapai total:", ways(total, k))
}
