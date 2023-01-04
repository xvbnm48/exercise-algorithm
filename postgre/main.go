package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // driver PostgreSQL
)

func main() {
	// Buka koneksi ke database
	db, err := sql.Open("postgres", `postgres://postgres:fariz@localhost/latihan?sslmode=disable`)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	// Persiapkan statement untuk menginsert data
	// stmt, err := db.Prepare("INSERT INTO  anime(title, genre, episodes, rating) VALUES($1, $2, $3, $4)")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer stmt.Close()

	// update data
	// stmt, err := db.Prepare("UPDATE anime SET title = $1, genre = $2, episodes = $3, rating = $4 WHERE anime_id = $5")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer stmt.Close()

	// get data
	stmt, err := db.Prepare("SELECT * FROM anime WHERE anime_id = $1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	// Eksekusi statement dengan mengisi placeholder dengan data yang akan diinsert
	// res, err := stmt.Exec("bloom into you / yagate kimi ni naru", "Yuri", 22, 9, 2)
	// cetak data dari get
	var anime_id int
	var title string
	var genre string
	var episodes int
	var rating float64
	err = stmt.QueryRow(2).Scan(&anime_id, &title, &genre, &episodes, &rating)
	// res, err := stmt.Exec(2)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Cek jumlah baris yang terpengaruh oleh statement
	// rowCount, err := res.RowsAffected()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("Number of rows affected:", rowCount)
	fmt.Println(anime_id, title, genre, episodes, rating)

}
