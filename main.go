package main
import (
	"fmt"
	"strings"
)

func main(){
	conferenceName :="Go Conference"
	const conferenceTickets int=50
	var remainingTickets uint=50
    var bookings []string
	

	//fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Println("We have total of %v tickets and %v are still available.\n", conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")


	
	

	for{
		var firstName string
	    var lastName string
	    var email string
	    var userTickets uint
    
	    //ask user for their name
    
	    fmt.Println("Enter your first name: ")
	    fmt.Scan(&firstName)
	    fmt.Println("Enter your last name: ")
	    fmt.Scan(&lastName)
	    fmt.Println("Enter your email: ")
	    fmt.Scan(&email)
	    fmt.Println("Enter number of tickets: ")
	    fmt.Scan(&userTickets)

		isValidName := len(firstName) >=2 && len(lastName)>=2
		isValidEmail := strings.Contains(email,"@")
		isValidTicketNumber:= userTickets > 0 && userTickets<= remainingTickets
		//isValidCity:= city == "Singapore" || city=="London"

		if isValidName && isValidEmail && isValidTicketNumber{

		    remainingTickets=remainingTickets-userTickets
	        //bookings[0]=firstName + " " +lastName
	        bookings=append(bookings,firstName+" "+lastName)
        
	        // fmt.Printf("The whole Slice: %v\n", bookings)
	        // fmt.Printf("The first Value: %v\n", bookings[0])
	        // fmt.Printf("Slice type: %T\n", bookings)
	        // fmt.Printf("Slice length: %v\n", len(bookings))
	        
	        fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n",firstName, lastName, userTickets,email)
            fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
            
            printFirstNames(bookings)
		    
    
            
		    if remainingTickets == 0 {
		    	//end program 
		    	fmt.Println("Our conferenceis booked out. Come back next year.")
		    	break
		    }
		    	
		}else{
			if !isValidName{
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail{
				fmt.Println("email address you entered doesn,t contains @ sign")
			}
			if !isValidTicketNumber{
				fmt.Println("number of tickets you entered is invalid")
			}
		}     
	} 
}   
func greetUsers(confName string,confTickets int,  remainingTickets uint){
	fmt.Println("Welcome to %v booking application\n", confName)
	fmt.Printf(We have total of %v tickets and %v are still available.\n", confTickets, remainingTickets)
	fmt.Println("get your tickets here to attend)
}
func printFirstNames(bookings []string){

	firstNames := []string{}
	for _, booking := range bookings{
		var names=strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("The first name of bookings are: %v\n", firstNames)
}




//Upto  2 hour 03 minutes 