package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("----- %v booking application -----\n", conferenceName)
	fmt.Printf("We habe total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// ask user for their name
	var userName string
	var userTickets int

	userName = "Marco"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
