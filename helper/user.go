package helper

type UserData struct {
	Name FullName
	Email string
	NmberOfTickets uint
}

type FullName struct {
	FirstName string
	LastName string
}
