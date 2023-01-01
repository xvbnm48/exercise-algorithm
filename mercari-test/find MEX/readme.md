Penjelasan 

Untuk menyelesaikan soal tersebut, kita dapat menggunakan pendekatan sebagai berikut:

Urutkan array N.
Inisialisasi variabel MEX dengan 0.
Lakukan perulangan dari 0 hingga panjang N. Pada setiap perulangan, lakukan:
a. Jika MEX sama dengan elemen ke-i di array N, tambahkan MEX dengan X dan lanjutkan ke langkah selanjutnya.
b. Jika MEX tidak sama dengan elemen ke-i di array N, maka MEX adalah elemen ke-i di array N. Kemudian, hentikan perulangan.
Cetak nilai dari MEX sebagai hasil dari soal tersebut.

Contoh implementasi dengan bahasa pemrograman Go adalah sebagai berikut:

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    // Buat array N
    N := []int{3, 4, 5}

    // Nilai X
    X := 1

    // Urutkan array N
    sort.Ints(N)

    // Inisialisasi MEX
    MEX := 0

    // Cari MEX
    for i := 0; i < len(N); i++ {
        if MEX == N[i] {
            MEX += X
        } else if MEX != N[i] {
            MEX = N[i]
            break
        }
    }

    // Cetak MEX
    fmt.Println(MEX)
}
```
Pada kode di atas, pertama-tama kita menyortir array N menggunakan fungsi sort.Ints(). Setelah itu, kita inisialisasi variabel MEX dengan 0. Kemudian, kita melakukan perulangan untuk mencari MEX dari array N dengan menggun



