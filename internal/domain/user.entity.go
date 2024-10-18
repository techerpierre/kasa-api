package domain

type User struct {
	ID               string
	Email            string
	Password         string
	Firstname        string
	Lastname         string
	Status           AuthStatus
	Picture          string
	Cover            string
	AuthorizationsID string
	Authorizations   Authorizations
	Accomodations    []Accommodation
	Bookings         []Booking
	Ratings          []Rating
	Comments         []Comment
}
