package main

import "strings"

func ValidateUserInput(remainingTickets int, userNameF string, userNameL string, email string, userTickets int) (bool, bool, bool) {
	isValidName := len(userNameF) >= 2 && len(userNameL) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
