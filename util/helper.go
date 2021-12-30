package util

import (
	"strings"
)

func ValidateInput(firstName string, lastName string, userEmail string,
	userTickets uint, remainTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(userEmail, "@")
	isValidTicketsNum := userTickets > 0 && userTickets <= remainTickets
	return isValidName, isValidEmail, isValidTicketsNum
}
