package entities

type User struct {
	ID               string
	Email            string
	Password         string
	Firstname        string
	Lastname         string
	Status           AuthStatus
	Picture          *string
	Cover            *string
	AuthorizationsID AuthStatus
	Authorizations   Authorizations
	Accomodations    []Accommodation
	Bookings         []Booking
	Ratings          []Rating
	Comments         []Comment
}
