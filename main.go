package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
)

var conferenceName = "Go Conference"
const totalConferenceTickets int = 50
var remainingTickets uint = uint(totalConferenceTickets)
var bookings []helper.UserData

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	for remainingTickets > 0 {

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(firstName, lastName, email, remainingTickets, userTickets)

		if isValidEmail && isValidName && isValidTicket {
			bookTicket(userTickets, firstName, lastName, email)
			fmt.Printf("The first names of bookings are: %v\n", getFirstNames())
			wg.Add(1)
			go generateTicket(userTickets, firstName, lastName)
			if remainingTickets == 0 {
				fmt.Println("Our conference is booked our")
				break
			}

		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short. Please enter Valid Name")
			}

			if !isValidEmail {
				fmt.Println("Please enter valid email")
			}

			if !isValidTicket {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}

	}
	wg.Wait()
}


func greetUsers() {
	fmt.Printf("Welcome to %v conference\n", conferenceName)
}

func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.Name.FirstName)
	}
	return firstNames
}


func getUserInput() (string, string, string, uint) {

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Enter your Last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets -= userTickets

	// create a map for user

	var fullName = helper.FullName {
		FirstName: firstName,
		LastName: lastName,
	}

	var userData = helper.UserData {
		Name : fullName,
		Email: email,
		NmberOfTickets: userTickets,
	}

	bookings = append(bookings,userData)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets are remaining for %v\n", remainingTickets, conferenceName)
	fmt.Printf("A confirmation mail will be send to %v\n", email)

	fmt.Printf("%v bookings so far from\n", len(bookings))
	fmt.Println(bookings)
}

func generateTicket(userTickets uint, firstName string, lastName string) {
	var ticket = helper.Ticket {
		UserTickets: userTickets,
		FirstName: firstName,
		LastName: lastName,
	}
	helper.SendTicket(ticket)
	wg.Done()
}
