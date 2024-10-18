package entities

import "time"

type Booking struct {
	ID              string
	Start           time.Time
	End             time.Time
	AccommodationID string
	Accommodation   Accommodation
	ClientID        string
	Client          User
}
