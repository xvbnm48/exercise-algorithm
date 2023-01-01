// package main

// import "fmt"

// const MOD = 1000000007

// func ways(pizza []string, k int) int {
// 	// Hitung jumlah 'A' pada tiap baris dan kolom
// 	rowCounts := make([]int, len(pizza))
// 	colCounts := make([]int, len(pizza[0]))
// 	for i := 0; i < len(pizza); i++ {
// 		for j := 0; j < len(pizza[0]); j++ {
// 			if pizza[i][j] == 'A' {
// 				rowCounts[i]++
// 				colCounts[j]++
// 			}
// 		}
// 	}

// 	// Buat tabel untuk menyimpan jumlah cara memotong pizza
// 	// dengan k-1 potongan sesuai dengan batas baris dan kolom
// 	dp := make([][][]int, len(pizza))
// 	for i := 0; i < len(pizza); i++ {
// 		dp[i] = make([][]int, len(pizza[0]))
// 		for j := 0; j < len(pizza[0]); j++ {
// 			dp[i][j] = make([]int, k)
// 		}
// 	}

// 	// Tentukan jumlah cara memotong pizza dengan k-1 potongan
// 	// untuk tiap batas baris dan kolom
// 	for i := 0; i < len(pizza); i++ {
// 		for j := 0; j < len(pizza[0]); j++ {
// 			for x := 0; x < k; x++ {
// 				// Jika ada setidaknya satu 'A' pada bagian atas
// 				// dan bawah dari batas baris, maka jumlah cara
// 				// memotong pizza bertambah
// 				if i > 0 && rowCounts[i-1] > 0 && rowCounts[len(pizza)-1]-rowCounts[i-1] > 0 {
// 					dp[i][j][x] += dp[i-1][j][x]
// 					if dp[i][j][x] >= MOD {
// 						dp[i][j][x] -= MOD
// 					}
// 				}

// 				// Jika ada setidaknya satu 'A' pada bagian kiri
// 				// dan kanan dari batas kolom, maka jumlah cara
// 				// memotong pizza bertambah
// 				if j > 0 && colCounts[j-1] > 0 && colCounts[len(pizza[0])-1]-colCounts[j-1] > 0 {
// 					dp[i][j][x] += dp[i][j-1][x]
// 					if dp[i][j][x] >= MOD {
// 						dp[i][j][x] -= MOD
// 					}
// 				}
// 			}
// 		}

// 		// Kembalikan jumlah cara memotong pizza dengan k-1 potongan
// 		// pada batas baris dan kolom terakhir
// 		return dp[len(pizza)-1][len(pizza[0])-1][k-1]
// 	}
// 	// Kembalikan jumlah cara memotong pizza dengan k-1 potongan
// 	// pada batas baris dan kolom terakhir
// 	return dp[len(pizza)-1][len(pizza[0])-1][k-1]
// }

// func main() {
// 	pizza := []string{
// 		"A.A",
// 		"AAA",
// 		"A.A",
// 	}

// 	k := 3

// 	fmt.Println(ways(pizza, k))
// }

package main

import "fmt"

func ways(pizza []string, k int) int {
	s := make([][]int, 51)
	for i := 0; i < len(s); i++ {
		s[i] = make([]int, 51)
	}
	count := func(r1, c1, r2, c2 int) int {
		var sum int
		for i := r1; i <= r2; i++ {
			sum += s[i][c2+1] - s[i][c1]
		}
		return sum
	}
	f := make([][][]int, 51)
	for i := 0; i < len(f); i++ {
		f[i] = make([][]int, 51)
		for j := 0; j < len(f[i]); j++ {
			f[i][j] = make([]int, 11)
		}
	}

	mod := int(1e9 + 7)
	rows, cols := len(pizza), len(pizza[0])
	for i := 0; i < rows; i++ {
		s[i][0] = 0
		for j := 0; j < cols; j++ {
			var tmp int
			if pizza[i][j] == 'A' {
				tmp = 1
			}
			s[i][j+1] = s[i][j] + tmp
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if count(i, j, rows-1, cols-1) > 0 {
				f[i][j][0] = 1
			}
		}
	}

	for p := 1; p < k; p++ {
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				f[i][j][p] = 0
				for q := j; q < cols; q++ {
					if count(i, j, rows-1, q) > 0 {
						f[i][j][p] = (f[i][j][p] + f[i][q+1][p-1]) % mod
					}
				}
				for q := i; q < rows; q++ {
					if count(i, j, q, cols-1) > 0 {
						f[i][j][p] = (f[i][j][p] + f[q+1][j][p-1]) % mod
					}
				}
			}
		}
	}
	return f[0][0][k-1]
}

func main() {
	pizza := []string{
		"A.A",
		"AAA",
		"A.A",
	}

	k := 3

	fmt.Println(ways(pizza, k))
}
