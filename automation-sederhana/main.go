package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 5)
	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("Running task...")
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	fmt.Println("Press 'q' to quit.")
	for {
		var input string
		fmt.Scanln(&input)
		if input == "q" {
			break
		}
	}

	close(quit)
	fmt.Println("Quit.")
}

// Path: automation-sederhana\main.go
