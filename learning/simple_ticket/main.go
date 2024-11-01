package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceTickets int
var conferenceName string
var remainingTickets uint
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// for managing Concurrency
var waitGroup = sync.WaitGroup{}

func main() {
	greetUsers()
	for {
		getData := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(getData.firstName, getData.lastName,
			getData.email, getData.numberOfTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(getData)

			waitGroup.Add(1)
			go sendTicket(getData)

			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
	waitGroup.Wait()
}

func greetUsers() {
	fmt.Println("Hello and welcome to the ticket booking application!")
	fmt.Println("Please enter your conference name:")
	fmt.Scan(&conferenceName)
	fmt.Println("Now, please enter the total capacity of the conference:")
	fmt.Scan(&conferenceTickets)
	remainingTickets = uint(conferenceTickets)

	fmt.Printf("Welcome to the %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
}

func getUserInput() UserData {
	var getData UserData
	fmt.Println("Enter your first name: ")
	fmt.Scan(&getData.firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&getData.lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&getData.email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&getData.numberOfTickets)

	return getData
}

func bookTicket(getUser UserData) {
	remainingTickets = remainingTickets - getUser.numberOfTickets
	bookings = append(bookings, getUser)
	fmt.Printf("List of bookings: %v\n", bookings)
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", getUser.firstName, getUser.lastName, getUser.numberOfTickets, getUser.email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(getData UserData) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", getData.numberOfTickets, getData.firstName, getData.lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, getData.email)
	fmt.Println("#################")
	waitGroup.Done()
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}
