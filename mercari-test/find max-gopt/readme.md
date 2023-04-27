Pada solusi di atas, fungsi MaximumMEX akan menerima array integer arr dan integer x sebagai input. Pertama-tama, array arr diurutkan dalam urutan menaik menggunakan fungsi sort.Ints.

Selanjutnya, kita mencari nilai MEX dengan melakukan iterasi pada array yang telah diurutkan. Nilai MEX awal diatur sebagai 1, dan kemudian nilai ini ditingkatkan secara bertahap ketika nilai array kurang dari atau sama dengan MEX. Jika nilai array lebih besar dari MEX, iterasi dihentikan.

Terakhir, kita periksa apakah nilai MEX dapat ditingkatkan dengan menambahkan atau mengurangi integer x. Untuk melakukannya, kita melakukan iterasi pada integer dari 1 hingga x dan mengecek apakah nilai yang dihasilkan dapat digunakan untuk menambahkan MEX. Jika ya, nilai MEX ditingkatkan.

Akhirnya, nilai MEX yang ditemukan dikembalikan.


Untuk input arr = [3,4,5] dan x = 1, langkah-langkah untuk mencari nilai MEX adalah sebagai berikut:

Array diurutkan menjadi [3, 4, 5].
MEX awal diatur sebagai 1.
Nilai array pertama adalah 3, yang kurang dari MEX (1), sehingga nilai MEX ditingkatkan menjadi 2.
Nilai array kedua adalah 4, yang sama dengan MEX (2), sehingga nilai MEX ditingkatkan menjadi 3.
Nilai array ketiga adalah 5, yang lebih besar dari MEX (3), sehingga iterasi dihentikan dan nilai MEX saat ini adalah 3.
Karena x = 1, kita memeriksa apakah MEX dapat ditingkatkan dengan menambahkan atau mengurangi 1. Karena 4 - 1 = 3, kita dapat menggunakan x untuk menambahkan nilai MEX. Oleh karena itu, nilai MEX ditingkatkan menjadi 4.
Karena nilai MEX yang ditemukan adalah 4, maka fungsi MaximumMEX akan mengembalikan nilai 3 (4-1). Sebagai hasilnya, hasil yang dicetak adalah 3