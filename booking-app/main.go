package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const totalTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets out of which we still have %v tickets available!\n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var userName string

	//ask user their name
	userName = "Tom"
	println(userName)

}
