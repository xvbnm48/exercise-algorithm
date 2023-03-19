package main

import (
	"fmt"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTicket int = 50
	var remainingTickets uint = 50
	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTicket, remainingTickets, conferenceName)
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v are still available. \n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend our conference")

	var firtsName string
	var lastName string
	var email string
	var userTickets uint
	fmt.Println("enter your first name:")
	fmt.Scan(&firtsName)
	fmt.Println("enter your last name:")
	fmt.Scan(&lastName)
	fmt.Println("enter your email:")
	fmt.Scan(&email)

	fmt.Println("enter the number of tickets you want to book")
	fmt.Scan(&userTickets)
	//userName = "John"
	fmt.Println("Hello", firtsName, lastName, "!")

	remainingTickets = remainingTickets - userTickets
	fmt.Printf("thanks you %v %v for booking %v tickets. Your tickets will be sent to %v\n", firtsName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
