package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/sekidome/bookingtut/helper" // path needs to be the same as specified in the go.mod file
)

const conferenceTickets int = 50

var (
	conferenceName        = "Go Conference"
	remainingTickets uint = 50
	names                 = make([]UserData, 0)
	wg                    = sync.WaitGroup{} // Tell mainThread to wait on other threads to complete before finishing the program.
	// When user would buy last ticket, the for loop ends and the programm exits, even if the sendTicket side thread isn t finished yet without the wg.
)

type UserData struct {
	firstName string
	lastName  string
	email     string
	order     uint
}

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

		userData := bookTicket(firstName, lastName, email, order)
		wg.Add(1)
		go sendTicket(userData)

		fmt.Printf("There are %d tickets left \n", remainingTickets)
		firstnames := genFirstnames()
		fmt.Printf("Hello, %s! Thanks for buying %d tickets. \n", firstName, order)
		fmt.Println(firstnames)
		fmt.Println(userData)
	}
	wg.Wait()
}

func userGreeting() {
	fmt.Printf("Event: %s. \nTickets: %d \nTickets remaining: %d\n", conferenceName, conferenceTickets, remainingTickets)
}

func genFirstnames() []string {
	var firstnames []string
	for _, booking := range names {
		firstnames = append(firstnames, booking.firstName)
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

func bookTicket(firstName string, lastName string, email string, order uint) UserData {

	remainingTickets -= order

	userData := UserData{
		firstName: firstName,
		lastName:  lastName,
		email:     email,
		order:     order,
	}

	names = append(names, userData) // assign the existing slice to the new longer slice. If typing := instead of = there will only be the last element remaining.
	fmt.Printf("Saved userdata: %v\n", names)
	return userData
}

func sendTicket(userData UserData) {
	time.Sleep(10 * time.Second)
	ticket := fmt.Sprintf("Thanks for buying %d tickets, %v. The tickets will be send to your email-address, %s.", userData.order, userData.firstName, userData.email)
	fmt.Println(ticket)
	wg.Done()
}
