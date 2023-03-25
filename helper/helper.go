package helper

import "strings"

func validateName(name string) bool {
	return len(name) >= 2
}

func validateEmail(email string) bool {
	return strings.Contains(email, "@")
}

func validateNoTickets(availableTickets uint, requestedTickets uint) bool {
	return requestedTickets > 0 && requestedTickets <= availableTickets
}

func ValidateUserInput(firstName string, lastName string, email string, availableTickets uint, requestedTickets uint) (bool, bool, bool) {
	return validateName(firstName) || validateName(lastName), validateEmail(email), validateNoTickets(uint(availableTickets), uint(requestedTickets))
}