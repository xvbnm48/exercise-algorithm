package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Bank merupakan struct untuk menyimpan data bank merchant
type Bank struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Category string `json:"category"`
}

var banks []Bank

func init() {
	banks = []Bank{
		Bank{ID: "bca", Name: "Bank Central Asia", Category: "Local Bank"},
		Bank{ID: "bni", Name: "Bank Negara Indonesia", Category: "Local Bank"},
		Bank{ID: "mandiri", Name: "Bank Mandiri", Category: "Local Bank"},
		Bank{ID: "citi", Name: "Citibank", Category: "International Bank"},
		Bank{ID: "hsbc", Name: "HSBC", Category: "International Bank"},
	}
}

// bankListHandler merupakan fungsi untuk menangani request list bank merchant
func bankListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(banks)
	if err != nil {
		return
	}
}

func LogRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	}
}

func main() {
	http.HandleFunc("/banks", LogRequest(bankListHandler))
	fmt.Println("Starting server at port 8000...")
	// make logger for server start
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		return
	}
}
