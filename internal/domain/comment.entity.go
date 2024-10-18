package domain

type Comment struct {
	ID              string
	Content         string
	AccommodationID string
	Accommodation   Accommodation
	UserID          string
	User            User
}
