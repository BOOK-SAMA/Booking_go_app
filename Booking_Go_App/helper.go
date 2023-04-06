package main

import "strings"

func ValidateUserInput(firstName string, lastName string, email string, userTicket uint, remainingTickets uint) (bool, bool, bool) {
	var lsValidName bool = len(firstName) >= 2 && len(lastName) >= 2
	IsvalidEmail := strings.Contains(email, "@")
	lsvaildticketnum := userTicket >= 0 && userTicket <= remainingTickets
	return IsvalidEmail, lsValidName, lsvaildticketnum
}
