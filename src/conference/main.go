package main

import (
	"conference/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"
var conferenceTicket = 50
var remainingTicket = 50

// var bookings = make([]map[string]string, 0) //creating list of map
var bookings = make([]userData, 0)

type userData struct {
	firstName    string
	lastName     string
	emailAddress string
	noOfTickets  int
}

var wg = sync.WaitGroup{}

func main() {

	greet()

	firstName, lastName, emailAddress, userTickets := getUserInput()
	//validate user input
	isValidName, isValidEmail, isValidTicketNo := helper.ValidateUserInput(firstName, lastName, emailAddress, userTickets, remainingTicket)

	fmt.Println(isValidName, isValidEmail, isValidTicketNo)

	if isValidName && isValidEmail && isValidTicketNo {

		//booking ticket logic
		bookTickets(firstName, lastName, userTickets, emailAddress)

		wg.Add(1)
		go sendTicket(userTickets, firstName, lastName, emailAddress)

		firstN := getFirstName()
		fmt.Printf("The total booking details are %v ", bookings)
		fmt.Printf("The first name  details are %v ", firstN)

		if remainingTicket == 0 {
			fmt.Println("Ticket is booked out")

		}
	} else if userTickets == remainingTicket {
		// do something
	} else {
		if !isValidName {
			fmt.Printf("The name you entered is wrong. You entered %v ", isValidName)

		}
		if !isValidEmail {
			fmt.Printf("The email you entered is wrong. You entered %v ", isValidEmail)
		}
		if !isValidTicketNo {
			fmt.Printf("The Ticket number you entered is wrong. You entered %v ", isValidTicketNo)
		}

	}

	wg.Wait()

}

func greet() {
	fmt.Printf("Welcome to %v  \n \n", conferenceName)
	fmt.Printf("Welcome to the %v booking system \n", conferenceName)
	fmt.Printf("The total conference ticket is %v and remaining ticket is %v \n ", conferenceTicket, remainingTicket)

}
func getFirstName() []string {
	firstN := []string{}
	//slicing first name from full name
	for _, booking := range bookings { //index replaced by _

		firstN = append(firstN, booking.firstName)
		// var name = strings.Fields(booking)
		// //var first = name[0]
		// firstN = append(firstN, name[0])

	}
	return firstN

}

func getUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var emailAddress string
	var userTickets int

	fmt.Println("Enter the  first name of the user :")
	fmt.Scan(&firstName)
	fmt.Println("Enter the last name of the user :")
	fmt.Scan(&lastName)
	fmt.Println("Enter the email address of the user :")
	fmt.Scan(&emailAddress)
	fmt.Println("Enter the amount of the tickets you want to purchage :")
	fmt.Scan(&userTickets)
	return firstName, lastName, emailAddress, userTickets
}
func bookTickets(firstName string, lastName string, userTickets int, emailAddress string) {
	remainingTicket = remainingTicket - userTickets
	// creating a map for the user data
	//var userData = make(map[string]string) //Create map
	var userData = userData{
		firstName:    firstName,
		lastName:     lastName,
		emailAddress: emailAddress,
		noOfTickets:  userTickets,
	}
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = emailAddress
	// userData["noOfTickets"] = strconv.FormatInt(int64(userTickets), 10)

	bookings = append(bookings, userData) //slice
	//fmt.Printf(bookings[0])

	//fmt.Println("Welcome to the", conferenceName, "Booking System")

}
func sendTicket(userTickets int, firstName string, lastName string, emailAddress string) {
	time.Sleep(8 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v ", userTickets, firstName, lastName)
	fmt.Println("#########")
	fmt.Printf("Sending ticket \n %v  \n to email address %v", ticket, emailAddress)
	fmt.Println("#########")
	wg.Done() // wait group says i am done
}
