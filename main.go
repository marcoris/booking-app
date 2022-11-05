package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	userMail    string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for {
		// ask user for their name
		firstName, lastName, userMail, userTickets := getUserInput()
		isValidName, isValidMail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, userMail, userTickets, remainingTickets)

		if isValidName && isValidMail && isValidTicketNumber {
			bookTicket(firstName, lastName, userMail, userTickets)

			wg.Add(1)
			go sendTicket(firstName, lastName, userMail, userTickets)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our Go conference is booked out. Come back next year :-)")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}
			if !isValidMail {
				fmt.Println("email address you entered is not correct")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}
		}
	}

	wg.Wait()
}

func greetUsers() {
	fmt.Printf("----- %v booking application -----\n", conferenceName)
	fmt.Printf("We habe total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Printf("Get your tickets here to attend\n\n")
}

func getFirstNames() []string {
	var firstNames []string
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userMail string
	var userTickets uint

	fmt.Print("Enter first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter Email: ")
	fmt.Scan(&userMail)

	fmt.Print("Enter Tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, userMail, userTickets
}

func bookTicket(firstName string, lastName string, userMail string, userTickets uint) {
	remainingTickets -= userTickets
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		userMail:    userMail,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("%v %v thx for booking %v tickets. You will receive a confirmation email at %v.\n\n", firstName, lastName, userTickets, userMail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

// Just simulate sending the tickets to the usermail
func sendTicket(firstName string, lastName string, userMail string, userTickets uint) {
	time.Sleep(30 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#####################")
	fmt.Printf("Sending ticket:\n%v to email address:\n%v\n", ticket, userMail)
	fmt.Println("#####################")

	wg.Done()
}
