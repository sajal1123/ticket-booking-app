package main

import (
	"fmt"
	"sync"
	"time"
)

var event string = "FIFA World Cup"

const finalTickets = 80000

var remainingTickets = 80000
var bookings = make([]UserData, 0)

type UserData struct {
	firstName    string
	lastName     string
	email        string
	numOfTickets int
}

var wg = sync.WaitGroup{}

func main() {

	greeting(event, remainingTickets, finalTickets)

	userNameF, userNameL, email, userTickets := getUserInput()

	isValidName, isValidEmail, isValidTicketNumber := ValidateUserInput(remainingTickets, userNameF, userNameL, email, userTickets)

	if isValidEmail && isValidName && isValidTicketNumber {

		bookTicket(userTickets, userNameF, userNameL, email)
		wg.Add(1)
		go sendTicket(userTickets, userNameF, userNameL, email)

		fNs := getFirstNames()

		fmt.Println("List of attendees: ", fNs)

		if remainingTickets <= 0 {
			fmt.Println("Tickets are sold out!")
		}

	} else {
		if isValidName != true {
			fmt.Println("First name or last name you entered is too short.")
		}
		if isValidEmail != true {
			fmt.Println("Your email address is invalid.")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of ticket is invalid.")
		}

		fmt.Printf("Your user data is invalid. Please try again.\n")
	}
	wg.Wait()
}

func greeting(e string, remainingTickets int, finalTickets int) {
	fmt.Println("Welcome to the", e)
	fmt.Printf("We have %v out of %v remaining\n", remainingTickets, finalTickets)
	fmt.Println("Get your tickers here")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, int) {
	var userNameF string
	var userNameL string
	var email string
	var userTickets int

	fmt.Println("Enter your first name:")
	fmt.Scan(&userNameF)

	fmt.Println("Enter your last name:")
	fmt.Scan(&userNameL)

	fmt.Println("Enter your email:")
	fmt.Scan(&email)

	fmt.Println("How many tickets would you like?")
	fmt.Scan(&userTickets)

	return userNameF, userNameL, email, userTickets
}

func bookTicket(userTickets int, userNameF string, userNameL string, email string) {
	remainingTickets = remainingTickets - userTickets

	var userData = UserData{
		firstName:    userNameF,
		lastName:     userNameL,
		email:        email,
		numOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v", bookings)

	fmt.Printf("Thank you for booking your tickets %v(%v).\n", userNameF, email)
	fmt.Printf("You have booked %v tickets.\nThere are %v tickers remaining.\n", userTickets, remainingTickets)

}

func sendTicket(userTickets int, firstName string, lastName string, email string) {
	var tickets = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	time.Sleep(100 * time.Second)
	fmt.Println("###################")
	fmt.Printf("Sending ticket:\n%v \nto email %v\n", tickets, email)
	fmt.Println("###################")
	wg.Done()

}
