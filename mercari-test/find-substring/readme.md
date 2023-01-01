Given s and x , determine the zero-based index of the first occurrence of x in s

String s consists of lowercase letters in the range ascii(a-z)

String x consists of lower case letters and may also contain a single wildcard character ‘*’ that represents any one character.

Example:
S=”xabcdey”
X=”ab*de”

The first match is at index 1

# penjelasan

Untuk menyelesaikan soal tersebut, kita dapat menggunakan pendekatan sebagai berikut:

1. Inisialisasi variabel index dengan -1.
2. Lakukan perulangan dari 0 hingga panjang s. Pada setiap perulangan, lakukan:
a. Jika elemen ke-i di s sama dengan elemen pertama di x, lakukan perulangan dari i hingga i + panjang x. Pada setiap perulangan, lakukan:
i. Jika elemen ke-j di s sama dengan elemen ke-j di x atau elemen ke-j di x adalah '', lanjutkan ke langkah selanjutnya.
ii. Jika elemen ke-j di s tidak sama dengan elemen ke-j di x atau elemen ke-j di x bukan '', hentikan perulangan dan lanjutkan ke langkah selanjutnya di perulangan utama.
b. Jika semua elemen di s dari i hingga i + panjang x sesuai dengan elemen di x, maka simpan i sebagai nilai dari variabel index, lalu hentikan perulangan.
3. Cetak nilai dari variabel index sebagai hasil dari soal tersebut.


Pada kode di atas, pertama-tama kita mendefinisikan string s dan x. Setelah itu, kita inisialisasi variabel index dengan -1 sebagai nilai default jika tidak ditemukan kecocokan antara x dan s. Kemudian, kita melakukan perulangan dari 0 hingga panjang s. Pada setiap perulangan, kita cek apakah elemen ke-i di s sama dengan elemen pertama di x. Jika iya, kita melakukan perulangan dari i hingga i + panjang x. Pada setiap perulangan, kita cek apakah elemen ke-j di s sama dengan elemen ke-j di x atau apakah elemen ke-j di x adalah '*'. Jika iya, maka lanjutkan ke langkah selanjutnya. Jika tidak, hentikan perulangan dan lanjutkan ke langkah selanjutnya di perulangan utama. Setelah semua elemen di s dari i hingga i + panjang x diterima, cek apakah elemen di s tersebut sama dengan x. Jika iya, simpan i sebagai nilai dari variabel index, lalu hentikan perulangan. Setelah semua elemen di s diterima, cetak nilai dari variabel index sebagai hasil dari soal tersebut.

Jika dijalankan, kode di atas akan menghasilkan output seperti ini:

Copy code
1
Sebagai catatan, kode di atas hanya salah satu contoh implementasi untuk menyelesaikan soal tersebut. Anda juga dapat menggunakan pendekatan yang berbeda sesuai dengan kebutuhan dan preferensi Anda.