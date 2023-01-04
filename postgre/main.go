package main

import (
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/lib/pq" // driver PostgreSQL
)

func main() {
	// Buka koneksi ke database
	password := "fariz shahab"
	encodedPassword := url.PathEscape(password)
	connectionString := fmt.Sprintf("postgres://postgres:%s@localhost/latihan?sslmode=disable", encodedPassword)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Persiapkan statement untuk menginsert data
	stmt, err := db.Prepare("INSERT INTO  mahasiswa(nim, nama, jenis_kelamin,tanggal_lahir, alamat, fakultas, jurusan) VALUES($1, $2, $3, $4, $5, $6, $7)")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	// Eksekusi statement dengan mengisi placeholder dengan data yang akan diinsert
	res, err := stmt.Exec(1, "Fariz Shahab", "Laki-laki", "1997-01-01", "Jl. Cikutra No. 1", "Teknik", "Informatika")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Cek jumlah baris yang terpengaruh oleh statement
	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Number of rows affected:", rowCount)
}
