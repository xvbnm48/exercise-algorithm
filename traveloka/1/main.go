package main

import (
	"fmt"
)

func slowestKey(keyTimes [][]int) string {
	if len(keyTimes) == 0 {
		return ""
	}

	maxDuration := keyTimes[0][1] // Durasi awal adalah waktu tekan tombol pertama
	slowestKey := keyTimes[0][0]  // Tombol terlambat awal adalah tombol pertama

	for i := 1; i < len(keyTimes); i++ {
		duration := keyTimes[i][1] - keyTimes[i-1][1]
		if duration > maxDuration {
			maxDuration = duration
			slowestKey = keyTimes[i][0]
		}
	}

	// Mengonversi tombol terlambat dari sebuah int ke string dengan menambahkannya ke 'a'
	return string('a' + slowestKey)
}

func main() {
	keyTimes := [][]int{{0, 2}, {1, 5}, {0, 9}, {2, 15}}
	fmt.Printf("Tombol terlambat adalah: %s\n", slowestKey(keyTimes))
}
