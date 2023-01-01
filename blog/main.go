package main

import "fmt"

type Mahasiswa struct {
	Nama string
	Umur int
	IPK  float64
}

func main() {
	var mahasiswa Mahasiswa
	mahasiswa.Nama = "sakura endo"
	mahasiswa.Umur = 20
	mahasiswa.IPK = 3.5
	fmt.Println(mahasiswa) // {sakura endo 20 3.5}
}
