package main

import "fmt"

func main() {
	// example for string manipulation
	var negara = "Japan"
	fmt.Println(negara)
	//_ = fmt.Sprintf("negara index 2: %c, ", negara[2])
	fmt.Println("negara index 2: ", string(negara[2]))
	// kalau 2:4 itu index 2 sampai 3
	fmt.Println("negara index 2:4, ", negara[2:4])
	// kalau 2: itu index 2 sampai akhir
	fmt.Println("negara index 2:, ", negara[2:])
	// kalau :4 itu index awal sampai 3
	fmt.Println("negara index :4, ", negara[:4])
	// kalau : itu index awal sampai akhir
	fmt.Println("negara index :4, ", negara[:])
}
