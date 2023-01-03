package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Mendapatkan informasi user agent
	userAgent := r.UserAgent()

	// Mendapatkan informasi alamat IP
	ipAddress := r.RemoteAddr

	// Tampilkan informasi user agent dan alamat IP
	fmt.Fprintf(w, "User agent: %s\n", userAgent)
	fmt.Fprintf(w, "Alamat IP: %s\n", ipAddress)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
