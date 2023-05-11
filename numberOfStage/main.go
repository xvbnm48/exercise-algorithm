//package main
//
//import "fmt"
//
//func numOfStage(angka int) int {
//	langkah := 0
//
//	for angka > 0 {
//		if angka%2 == 0 {
//			angka /= 2
//		} else {
//			angka -= 1
//		}
//		langkah++
//	}
//	return langkah
//
//}
//
//func main() {
//	angka := 90
//	fmt.Printf("dari angka %v, ini memeiliki %v langkah pengurangan untuk sampai menjadi 0\n",// angka, numOfStage(angka))
//}

package main

import "fmt"

func findStepNum(num int) int {
	step := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}

		step++
	}
	return step
}

func main() {
	angka := 48
	fmt.Printf("Dari angka ini %v memiliki jumlah langkah %vx untuk agar menjadi 0 ", angka, findStepNum(angka))
}
