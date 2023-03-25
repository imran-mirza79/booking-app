package helper

import (
	"fmt"
	"time"
)

func SendTicket(ticket Ticket)  {
	time.Sleep(5* time.Second)
	fmt.Println("\n===========================")
	fmt.Printf("%v tockets for %v %v\n",ticket.UserTickets, ticket.FirstName, ticket.LastName )
	fmt.Println("\n===========================")

}