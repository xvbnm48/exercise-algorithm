package main

import (
	"fmt"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTicket = 50
	var remainingTickets int = 50
	//fmt.Println("welcome to our", conferenceName, "booking application")
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Println("We have", remainingTickets, "tickets left")
	fmt.Println("Get your tickets here to attend our conference")

}
