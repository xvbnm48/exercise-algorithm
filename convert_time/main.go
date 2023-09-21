package main

import (
	"fmt"
	"strings"
)

func convertTime(time string) string {
	parts := strings.Split(time, ":")
	hour, minute, ampm := parts[0], parts[1][:2], parts[1][2:]

	// Konversi jam ke integer
	var hourInt int
	_, err := fmt.Sscanf(hour, "%d", &hourInt)
	if err != nil {
		return "Invalid time format"
	}

	// Periksa AM atau PM
	if ampm == "AM" {
		if hourInt == 12 {
			hourInt = 0
		}
	} else if ampm == "PM" {
		if hourInt != 12 {
			hourInt += 12
		}
	} else {
		return "Invalid time format"
	}

	// Format jam dan menit dalam format 24 jam
	formattedTime := fmt.Sprintf("%02d:%s", hourInt, minute)
	return formattedTime
}

func main() {
	// Contoh penggunaan
	waktuAM := "08:45 AM"
	waktuPM := convertTime(waktuAM)
	fmt.Println("Waktu dalam format PM:", waktuPM)
}
