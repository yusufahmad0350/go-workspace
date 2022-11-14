package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	var conferenceTicket = 50
	var remainingTicket = 50
	var userTickets int
	var i int
	var firstName string
	var lastName string
	var emailAddress string
	fmt.Println("Enter the amount of the tickets you want to purchage :")
	fmt.Scan(&userTickets)
	var noBookings = []string{}

	for i = 0; i < userTickets; i++ {

		fmt.Println("Enter the  first name of the user :")
		fmt.Scan(&firstName)
		fmt.Println("Enter the last name of the user :")
		fmt.Scan(&lastName)
		fmt.Println("Enter the email address of the user :")
		fmt.Scan(&emailAddress)

		noBookings = append(noBookings, firstName+" "+lastName, "and email is "+emailAddress)

	}

	var bookings = []string{}
	bookings = append(bookings, firstName+" "+lastName) //slice
	fmt.Printf(bookings[0])

	//fmt.Println("Welcome to the", conferenceName, "Booking System")
	fmt.Printf("Welcome to the %v booking system \n", conferenceName)
	fmt.Printf("The total conference ticket is %v and remaining ticket is %v \n ", conferenceTicket, remainingTicket)
	//fmt.Printf("We have %v available and %v are remaining ticket \n", conferenceTicket, remainingTicket)
	//fmt.Printf("the type of conference name is %T and conference ticket is %T and remaining ticket is %T and user name is %T and user ticket is %T\n", conferenceName, conferenceTicket, remainingTicket, userName, userTickets)
	fmt.Printf("The name of the user is %v %v for booking %v tickets. you will recieve confirmation on %v shortly \n", firstName, lastName, userTickets, emailAddress)
	remainingTicket = remainingTicket - userTickets
	fmt.Printf("The remaining tickets is %v \n", remainingTicket)
	fmt.Printf("The total booking details are %v ", bookings)

	fmt.Printf("The final booking is ", noBookings)

}
