package main

import "fmt"

func main() {
	const conferenceName = "Go Conference"
	const totalTickets uint = 50

	var remainingTickets uint = 50

	fmt.Printf("Welcome to %s booking application! \n", conferenceName)
	fmt.Printf("We have total of %d tickets and %d are still available. \n", remainingTickets, totalTickets)
	fmt.Println("Get your tickets here to attend")
}
