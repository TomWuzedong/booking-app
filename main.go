package main

import (
	"booking-app/util"
	"fmt"
	"sync"
	"time"
)

const confTickets = 50

var confName = "Go Conference"
var remainTickets uint = confTickets
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	welcome()
	firstName, lastName, userEmail, userTickets := getInput()
	isValidName, isValidEmail, isValidTicketsNum :=
		util.ValidateInput(firstName, lastName, userEmail, userTickets, remainTickets)

	if isValidName && isValidEmail && isValidTicketsNum {
		fmt.Println("############")

		bookTickets(firstName, lastName, userEmail, userTickets)
		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, userEmail)

		firstNames := getFirstNames()
		fmt.Printf("users with bookings are: %v\n", firstNames)
		fmt.Println("")
		if remainTickets == 0 {
			fmt.Println("All tickets are booked out")
		}
	} else {
		handleInvalidInput(isValidName, isValidEmail, isValidTicketsNum)
	}
	wg.Wait()
}

func welcome() {
	fmt.Printf("Welcome to the %v booking application\n", confName)
	fmt.Printf("Total of %v with %v available\n", confTickets, remainTickets)
	fmt.Println("Get your tickets here to attend")
}

func getInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var userEmail string
	var userTickets uint

	fmt.Print("Please enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Print("Please enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Print("Please enter your email address: ")
	fmt.Scan(&userEmail)

	fmt.Print("Please enter the number of tickets to book: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, userEmail, userTickets
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func bookTickets(firstName string, lastName string, userEmail string, userTickets uint) {
	remainTickets = remainTickets - userTickets

	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           userEmail,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)

	fmt.Printf("User %v booked %v tickets using %v\n",
		lastName, userTickets, userEmail)
	fmt.Printf("There are %v tickets currently remaining, %v bookings made: %v \n",
		remainTickets, len(bookings), bookings)
}

func handleInvalidInput(isValidName bool, isValidEmail bool, isValidTicketsNum bool) {
	if !isValidName {
		fmt.Println("first or last name entered is invalid")
	}
	if !isValidEmail {
		fmt.Println("email entered is invalid")
	}
	if !isValidTicketsNum {
		fmt.Println("number of tickets to book is invalid")
	}
	fmt.Println("please try again")
}

func sendTickets(userTickets uint, firstName string, lastName string, userEmail string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v\n", userTickets, firstName, lastName)
	fmt.Println("############")
	fmt.Printf("Sending ticket:\n %v to email address: %v\n", ticket, userEmail)
	fmt.Println("############")
	wg.Done()
}
