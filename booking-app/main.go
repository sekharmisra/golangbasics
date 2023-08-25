package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const totalTickets uint = 50
	var remainingTickets uint = 50
	//var bookings []string
	bookings := []string{}

	fmt.Printf("Conference name is of type %T, totaltickets is of type %T, remainingtickets is of type %T \n", conferenceName, totalTickets, remainingTickets)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets out of which we still have %v tickets available!\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
	fmt.Printf("Booking is of type  %T\n", bookings)

	for {
		//Infinite loop
		var firstName string
		var lastName string
		var userTickets uint
		var email string

		fmt.Println("Enter your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter number of tickets to purchased:")
		fmt.Scan(&userTickets)

		fmt.Println("Enter your email:")
		fmt.Scan(&email)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole slice is: %v\n", bookings)
		fmt.Printf("The first element of the slice is: %v\n", bookings[0])
		fmt.Printf("The size of the slice is: %v\n", len(bookings))
		fmt.Printf("The type of the slice %T\n", bookings)

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email on %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("Remaining tickets %v for the %v\n", remainingTickets, conferenceName)

		fmt.Printf("All the bookings in the application: %v\n", bookings)

	}

}
