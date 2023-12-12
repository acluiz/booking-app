package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceName = "Go Conference"
	const totalTickets uint = 50

	var remainingTickets uint = totalTickets
	bookings := []string{}

	fmt.Printf("Welcome to %s booking application! \n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available. \n", remainingTickets, totalTickets)
	fmt.Println("Get your tickets here to attend")

	for remainingTickets > 0 {
		var userTickets uint

		var firstName string
		var lastName string
		var email string
	
		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)
	
		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter the number of tickets you want:")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2 
		isValidEmail := strings.Contains(email, "@")
		isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketsNumber {
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName + " " + lastName)

			fmt.Printf("Thank you %s %s for booking %d tickets. You will receive a confirmation email at %s\n", firstName, lastName, userTickets, email)
			fmt.Printf("%d tickets remaining for %s\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {
				names := strings.Fields(booking)
				firstName := names[0]
				firstNames = append(firstNames, firstName)
			}

			fmt.Printf("The first names of bookings are: %v\n", firstNames)

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
