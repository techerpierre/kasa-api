package entities

type Rating struct {
	ID              string
	Value           int
	AccommodationID string
	Accommodation   Accommodation
	UserID          string
	User            User
}
