package main

import (
	"fmt"
	"strings"

	"github.com/sekidome/bookingtut/helper" // path needs to be the same as specified in the go.mod file
)

const conferenceTickets int = 50

var (
	conferenceName        = "Go Conference"
	remainingTickets uint = 50
	names            []string
)

func main() {
	userGreeting()

	for remainingTickets != 0 {
		firstName, lastName, email, order := getUserInput() // die funktion muss ausgefuehrt werden, um den entsprechenden Wert an die variable returnen zu können -> decleration aber gleichzeitig execution
		validName, validEmail, validTicket := helper.ValidInput(firstName, lastName, email, order, remainingTickets)

		if validName {
			fmt.Println("Name too short")
			continue
		}
		if validEmail {
			fmt.Println("Wrong Email")
			continue
		}
		if validTicket {
			fmt.Printf("Only %d Tickets remaining.\n", remainingTickets)
			continue
		}

		bookTicket(firstName, lastName, order)

		fmt.Printf("There are %d tickets left \n", remainingTickets)
		firstnames := genFirstnames()
		fmt.Printf("Hello, %s! Thanks for buying %d tickets. \n", firstName, order)
		fmt.Println(firstnames)
	}
}

func userGreeting() {
	fmt.Printf("Event: %s. \nTickets: %d \nTickets remaining: %d\n", conferenceName, conferenceTickets, remainingTickets)
}

func genFirstnames() []string {
	var firstnames []string
	for _, booking := range names {
		var name = strings.Fields(booking)
		firstnames = append(firstnames, name[0])
	}
	return firstnames
}

func getUserInput() (string, string, string, uint) {
	var (
		firstName string
		lastName  string
		email     string
		order     uint
	)
	fmt.Println("Enter your Firstname:")
	fmt.Scanln(&firstName)
	fmt.Println("Enter your Lastname:")
	fmt.Scanln(&lastName)
	fmt.Println("Enter your Email:")
	fmt.Scanln(&email)
	fmt.Println("Enter number of tickets:")
	fmt.Scan(&order) // Scan does not recognize spaces?

	return firstName, lastName, email, order
}

func bookTicket(firstName string, lastName string, order uint) {
	names = append(names, firstName+" "+lastName) // assign the existing slice to the new longer slice. If typing := instead of = there will only be the last element remaining.
	remainingTickets -= order
}
