package main

import (
	"fmt"
	"strings"
)

func main() {
	// this code about booking conference tickets with a basic golang knowledge
	conferenceName := "Go Conference"
	const conferenceTicket int = 50
	var remainingTickets uint = 50
	var bookings []string

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTicket, remainingTickets, conferenceName)
	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total %v tickets and %v are still available. \n", conferenceTicket, remainingTickets)
	fmt.Println("Get your tickets here to attend our conference")

	for {
		// var firtsName is a string variable
		var firtsName string
		var lastName string
		var email string
		var userTickets uint
		fmt.Println("enter your first name:")
		_, _ = fmt.Scan(&firtsName)

		fmt.Println("enter your last name:")
		_, _ = fmt.Scan(&lastName)

		fmt.Println("enter your email:")
		_, _ = fmt.Scan(&email)

		fmt.Println("enter the number of tickets you want to book")
		_, _ = fmt.Scan(&userTickets)

		isValidName := len(firtsName) > 2 && len(lastName) > 2
		isValidEmail := strings.Contains(email, "@")
		isValidTickets := userTickets > 0 && userTickets <= remainingTickets

		if isValidTickets && isValidEmail && isValidName {
			fmt.Println("Hello", firtsName, lastName, "!")

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firtsName+" "+lastName)

			fmt.Printf("the whole slice is %v\n", bookings)
			fmt.Printf("the first element is %v\n", bookings[0])
			fmt.Printf("slice type is %T\n", bookings)
			fmt.Printf("slice length is %v\n", len(bookings))

			fmt.Printf("thanks you %v %v for booking %v tickets. Your tickets will be sent to %v\n", firtsName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
			firtsnames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firtsnames = append(firtsnames, names[0])
			}
			fmt.Printf("The first names of bookings are:  %v\n", firtsnames)
			if remainingTickets == 0 {
				fmt.Println("Sorry, we are sold out!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("Please enter a valid name")
			}

			if !isValidEmail {
				fmt.Println("Please enter a valid email")
			}

			if !isValidTickets {
				fmt.Println("Please enter a valid number of tickets")
			}
		}

	}
}
