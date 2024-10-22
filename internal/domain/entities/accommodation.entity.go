package entities

type Accommodation struct {
	ID               string
	Title            string
	Description      string
	Cover            string
	Pictures         []string
	Adress           string
	AdditionalAdress *string
	Zip              string
	City             string
	Country          string
	Active           bool
	Equipments       []string
	Tags             []string
	UserID           string
	User             User
	Ratings          []Rating
	Comments         []Comment
	Bookings         []Booking
}
