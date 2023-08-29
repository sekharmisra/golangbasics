package main

import (
	"fmt"
	"time"
)

const totalTickets uint = 50

var conferenceName = "Go conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	userTickets uint
}

func main() {

	//Step 0 - Print selected conference city!
	printSelectedConferenceCity()

	//Step 1 - Greet users
	greetUsers()

	for {
		//Infinite loop

		//Step 2: Get user inputs
		firstName, lastName, email, userTickets := getUserInput()

		//Step 3: Validate user inputs
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			//Step 4: Book the tickets - Call the function
			bookTickets(firstName, lastName, email, userTickets)

			//Send tickets for email
			go sendTickets(firstName, lastName, email, userTickets)

			//Step 5: Call print first name function
			firstNames := getFirstNames()
			fmt.Printf("The firstnames of the bookings in the application are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("All tickets are sold out for the conference. Please come next year!")
				//End the program by breaking the loop!
				break
			}
		} else {

			//User input validations!

			if !isValidName {
				fmt.Println("Firstname or lastname is too short!")
			}

			if !isValidEmail {
				fmt.Println("Email does not contain @ symbol!")
			}

			if !isValidTicketNumber {
				fmt.Println("Ticket number is invaid!")
			}
		}

	}

}

// Greet user function - Greets the users
func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets out of which we still have %v tickets available!\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

// Prints firstnames of the users
func getFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

// Get user input
func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets to purchased:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}

// Print selected conference city!
func printSelectedConferenceCity() {
	//Switch Statements:

	city := "London"

	switch city {

	case "New York":
		fmt.Println("New York city selected!")

	case "London":
		fmt.Println("London city selected!")

	case "New Dehli", "Bangalore":
		fmt.Println("New Delhi & Banglaore city selected!")

	default:
		fmt.Println("Invalid city selected!")

	}
}

// Finally: book tickets!
func bookTickets(firstName string, lastName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	//Create a map
	var userData = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		userTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("List of bookings is %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive confirmation email on %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("Remaining tickets %v for the %v\n", remainingTickets, conferenceName)
}

func sendTickets(firstName string, lastName string, email string, userTickets uint) {
	//Sleep for 10 seconds on the current thread..
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##############################")
	fmt.Printf("Sending ticket:\n %v \n to the email address %v\n", ticket, email)
	fmt.Println("##############################")
}
