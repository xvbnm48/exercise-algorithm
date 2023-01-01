Start with a given array of integers and an arbitrary initial value x. Calculate the running sum of x plus each array element, from left to right. The running sum must never get below 1. 
Determine the minimum value of x. 

Example :  arr = [-2,3,1,-5]

If x = 4, the following results are obtained: 

Running sum         arr[i] 
4                     -2
2                      3
5                      1
6                      -5
1

The final value is 1, and the running sum has never dropped below 1. 
Hence, The minimum starting value for x is 4


# penjelasan

Untuk menyelesaikan soal tersebut, kita dapat menggunakan pendekatan sebagai berikut:

1. Inisialisasi variabel x dengan maksimal dari seluruh elemen di array arr.
2. Lakukan perulangan dari 0 hingga panjang arr. Pada setiap perulangan, lakukan:
a. Hitung running sum dengan menambahkan x dan elemen ke-i di arr.
b. Jika running sum kurang dari 1, ubah nilai x dengan selisih antara 1 dan running sum, lalu lanjutkan ke langkah selanjutnya.
c. Jika running sum tidak kurang dari 1, lanjutkan ke langkah selanjutnya.
3. Cetak nilai dari x sebagai hasil dari soal tersebut.

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    // Buat array arr
    arr := []int{-2, 3, 1, -5}

    // Inisialisasi x dengan maksimal dari arr
    x := math.MaxInt32
    for _, a := range arr {
        if a > x {
            x = a
        }
    }

    // Hitung running sum
    for i := 0; i < len(arr); i++ {
        sum := x + arr[i]
        if sum < 1 {
            x = 1 - sum
        }
    }

    // Cetak x
    fmt.Println(x)
}
```