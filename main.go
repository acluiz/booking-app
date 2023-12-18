package main

import (
	"fmt"
	"strconv"
)

const conferenceName = "Go Conference"
const totalTickets uint = 50

var	remainingTickets uint = totalTickets
var	bookings = make([]map[string]string, 0)

func greetUsers() {
	fmt.Printf("Welcome to %s booking application! \n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available. \n", remainingTickets, totalTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() ([]string) {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking["firstName"])
	}

	return firstNames
}

func getUserInput()(string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address:")
	fmt.Scan(&email)

	fmt.Println("Enter the number of tickets you want:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	userData := make(map[string]string)

	userData["firstName"] = firstName
	userData["lastName"] = lastName
	userData["email"] = email
	userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userData)

	fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastName, userTickets, email)
	fmt.Printf("%d tickets remaining for %s\n", remainingTickets, conferenceName)
}

func main() {
	greetUsers()

	for remainingTickets > 0 {
		firstName, lastName, email, userTickets:= getUserInput()
		isValidName, isValidEmail, isValidTicketsNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketsNumber {
			bookTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)
			fmt.Printf("totalTickets: %v, remainingTickets: %v\n", totalTickets, remainingTickets)

			noTicketsRemaining := remainingTickets == 0

			if noTicketsRemaining {
				fmt.Println("Our conference is booked out. Come back next year!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("The first name or the last name you entered is too short")
			} 
			
			if !isValidEmail {
				fmt.Println("The email address you entered doesn't contain a '@' sign")
			} 
			
			if !isValidTicketsNumber {
				fmt.Println("The numberof tickets you entered is invalid")
			}
		}
	}
}
