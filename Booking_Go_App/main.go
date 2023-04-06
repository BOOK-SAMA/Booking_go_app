package main

import (
	"fmt"
	"sync"
)

var conferenceName = "Morgan Carmero"

const conferencTickets uint = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName      string
	lastName       string
	email          string
	numberOFticket uint
}

var wg = sync.WaitGroup{}

func main() {
	//var conferenceName = "Morgan Carmero"
	greetUser()
	//var booking = [50]string{}

	/*var firstName string
	var lastName string
	var email string
	var userTicket uint
	*/
	for {
		firstName, lastName, email, userTicket := GetUserInput()
		IsvalidEmail, lsValidName, lsvaildticketnum := ValidateUserInput(firstName, lastName, email, userTicket, remainingTickets)

		if IsvalidEmail && lsvaildticketnum && lsValidName {
			BookingTicket(userTicket, firstName, lastName, email)

			wg.Add(1)

			go sendTicket(userTicket, firstName, lastName, email)

			//1.06.57

			FN := PrintFirstname()
			fmt.Printf("Custome name : %v", FN)

			if remainingTickets == 0 {
				fmt.Println("ticket is out")
				break
			}
		} else {
			fmt.Println("something worng!!!!!!")
			continue
		}

		if userTicket > remainingTickets {
			fmt.Println("fuck off you can't booking more ticket")
			continue
		}

	}
	wg.Wait()

}

func greetUser() {
	fmt.Printf("Welcome to %v booking app\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v remaining\n", conferencTickets, remainingTickets)
	fmt.Println("Get Your ticket Here")
}

func PrintFirstname() []string {
	firstName_only := []string{}
	//_ use to ignore varible
	for _, element := range bookings {
		// var firstName_get = names[index] or use
		firstName_only = append(firstName_only, element.firstName)
	}
	return firstName_only
}

/*
	func ValidateUserInput(firstName string, lastName string, email string, userTicket uint) (bool, bool, bool) {
		var lsValidName bool = len(firstName) >= 2 && len(lastName) >= 2
		IsvalidEmail := strings.Contains(email, "@")
		lsvaildticketnum := userTicket >= 0 && userTicket <= remainingTickets
		return IsvalidEmail, lsValidName, lsvaildticketnum
	}
*/
func GetUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTicket uint
	//ask user for their name
	fmt.Println("Please Enter Your FirstName :")
	fmt.Scan(&firstName)

	fmt.Println("Please Enter Your lastName :")
	fmt.Scan(&lastName)

	fmt.Println("Please Enter Your email :")
	fmt.Scan(&email)

	fmt.Println("Please Enter Your ticket :")
	fmt.Scan(&userTicket)

	return firstName, lastName, email, userTicket

}

func BookingTicket(userTicket uint, firstName string, lastName string, email string) {

	remainingTickets = remainingTickets - userTicket
	//bookings[0] = firstName + " " + lastName

	//create map for user
	//var myslice []string
	// var mymap map[int]int

	var userData = UserData{
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		numberOFticket: userTicket,
	}
	/*
		userData["firstName"] = firstName
		userData["lastname"] = lastName
		userData["email"] = email
		userData["numberOFticket"] = strconv.FormatUint(uint64(userTicket), 10)
	*/

	bookings = append(bookings, userData)
	//add element to slice use apped
	// lenth of array => len(arrys)
	fmt.Printf("List of bookings : %v \n", bookings)
	fmt.Printf("Your name is %v %v and Booked %v Tickets\n", firstName, lastName, userTicket)
	fmt.Printf("Remaining tickets : %v\n", remainingTickets)

}

// city := "london"
/* switch city {
	case "New York"
		// excute code for booking new york
	case "london"
		// excute code for booking londonno city
	case "thai","bankok"
		// excute code for booking
	default:
	fmt.Print(fuck off  select)

}
note 2.01.08
*/
func sendTicket(userTicket uint, firstname string, lastname string, email string) {
	var ticket = fmt.Sprintf("%v ticket for %v %v ", userTicket, firstname, lastname)
	fmt.Printf("sending ticket %v to email address %v", ticket, email)
	wg.Done()
}
