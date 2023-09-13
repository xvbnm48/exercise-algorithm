They say that the grass is always greener on the other side. This means that other people always seem to be doing better than you. Since we are programmers, we can quantify the greenness of all the "other sides" by writing some code!

There are N gardens in a row. The gardens are numbered 1 to n and the i -th garden has a[i] greenness of , with higher numbers representing greener gardens. Let us define the "greenness on the other side" for the owner of the i-th garden as the greenest garden besides the i-th garden.
Create a program to enumerate all of the "greenness on the other side" for every garden in a row.
Given the greenness of  N  gardens, please calculate the "greenness on the other side" for each garden.


Dalam soal ini, kita diminta untuk menghitung "greenness on the other side" untuk setiap taman dalam sebuah barisan taman, di mana setiap taman memiliki tingkat kehijauan yang berbeda.


```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
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

```

Dalam solusi ini, kita menghitung dua array, yaitu leftMax dan rightMax, yang masing-masing akan menyimpan nilai maksimum di sebelah kiri dan kanan setiap taman. Kemudian, kita menggunakan kedua array ini untuk menghitung "greenness on the other side" untuk setiap taman dan mencetak hasilnya. Solusi ini akan bekerja dengan benar sesuai dengan deskripsi lengkap soal.

```go
func main() {
	lines := getStdin()
	n, _ := strconv.Atoi(lines[0])
	greenness := make([]int, n)

```

Di dalam fungsi main(), kita pertama-tama membaca input dari standar masukan (stdin) menggunakan fungsi getStdin(). Input pertama (baris kedua kode) adalah jumlah taman (N), jadi kita mengkonversi nilainya ke dalam tipe data integer (dengan strconv.Atoi) dan membuat slice greenness untuk menyimpan tingkat kehijauan dari setiap taman.


```go
	greennessStr := strings.Split(lines[1], " ")
	for i := 0; i < n; i++ {
		greenness[i], _ = strconv.Atoi(greennessStr[i])
	}

```

Selanjutnya, kita mengambil baris kedua input (yang berisi tingkat kehijauan masing-masing taman) dan memisahkannya menjadi potongan-potongan berdasarkan spasi. Kemudian, kita mengonversi potongan-potongan tersebut menjadi nilai integer dan menyimpannya di dalam slice greenness.

```go
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	leftMax[0] = greenness[0]
	rightMax[n-1] = greenness[n-1]

```

Kemudian, kita membuat dua slice lagi, yaitu leftMax dan rightMax. Ini digunakan untuk menyimpan nilai maksimum di sebelah kiri dan kanan setiap taman. Pada awalnya, kita menginisialisasi leftMax dengan nilai tingkat kehijauan taman pertama dan rightMax dengan nilai taman terakhir.


```go
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], greenness[i])
	}

	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], greenness[i])
	}

```

Kemudian, kita melakukan iterasi melalui taman-taman, baik dari kiri ke kanan (pada leftMax) maupun dari kanan ke kiri (pada rightMax). Pada setiap iterasi, kita memperbarui nilai maksimum di sebelah kiri dan kanan taman saat ini dengan membandingkannya dengan nilai maksimum yang telah dihitung sebelumnya.

```go
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

```

Akhirnya, kita membuat slice result yang akan menyimpan "greenness on the other side" untuk setiap taman. Di sini, kita memeriksa apakah taman saat ini berada di ujung barisan taman atau tidak, dan kemudian menghitung "greenness on the other side" sesuai dengan deskripsi yang diberikan dalam soal.


```go
	for _, v := range result {
        fmt.Println(v)
        }
}

```

Terakhir, kita mencetak hasilnya ke layar menggunakan perulangan for dan fmt.Println(). Ini akan mencetak "greenness on the other side" untuk setiap taman sesuai dengan urutan yang benar.