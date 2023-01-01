package main

import "fmt"

func hasFlightTime(films []int, flightTime int) bool {
	// Memeriksa apakah input list kosong atau tidak
	if len(films) == 0 {
		return false
	}

	// Melakukan iterasi melalui setiap film dalam daftar
	for i, film1 := range films {
		// Melakukan iterasi melalui setiap film lainnya dalam daftar
		for j, film2 := range films {
			// Memeriksa apakah film tidak sama dan jumlah waktu yang dibutuhkan untuk menonton film sama dengan waktu penerbangan
			if i != j && film1+film2 == flightTime {
				return true
			}
		}
	}

	return false
}

func main() {
	films := []int{60, 90, 120}
	flightTime := 180

	if hasFlightTime(films, flightTime) {
		fmt.Println("Ada dua film yang dapat ditonton selama penerbangan.")
	} else {
		fmt.Println("Tidak ada dua film yang dapat ditonton selama penerbangan.")
	}

}
