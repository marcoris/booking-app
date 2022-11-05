package helper

import "strings"

func ValidateUserInput(firstName string, lastName string, userMail string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidMail := strings.Contains(userMail, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidMail, isValidTicketNumber
}
