package main

import "fmt"

func main() {
	// Varaibles
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remaingTickets = 10

	// fmt.Println("Welcome to", conferenceName, " booking application")
	// fmt.Println("We have total of", conferenceTickets, "tickets and only", remaingTickets, "tickets are left")
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and only %v tickets are left\n", conferenceTickets, remaingTickets)
	fmt.Println("Get your tickets here to attend", conferenceTickets)

	var firstName string
	var lastName string
	var email string
	var userTickets int = 5
	fmt.Println("Please enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("Please enter your email")

	fmt.Scan(&email)

	fmt.Println("Please enter the number of tickets you want to book")

	fmt.Scan(&userTickets)

	fmt.Println("Hello", firstName, lastName, "you have booked", userTickets, "tickets for", conferenceName)

	// Array and slices

}
