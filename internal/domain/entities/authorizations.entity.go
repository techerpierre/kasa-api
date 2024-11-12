package entities

type AuthorizationType int

const (
	Authorization_CreateUser AuthorizationType = iota
	Authorization_CreateAuthorization
	Authorization_UpdateAuthorization
	Authorization_DeleteAuthorization
	Authorization_UpdateUser
	Authorization_DeleteUser
	Authorization_CreateAccommodation
	Authorization_UpdateAccommodation
	Authorization_DeleteAccommodation
	Authorization_CreateBooking
	Authorization_UpdateBooking
	Authorization_DeleteBooking
	Authorization_CreateRating
	Authorization_UpdateRating
	Authorization_DeleteRating
	Authorization_CreateComment
	Authorization_UpdateComment
	Authorization_DeleteComment
	Authorization_NoAuthorization
)

type Authorizations struct {
	ID                  string
	CreateAuthorization bool
	UpdateAuthorization bool
	DeleteAuthorization bool
	CreateUser          bool
	UpdateUser          bool
	DeleteUser          bool
	CreateAccommodation bool
	UpdateAccommodation bool
	DeleteAccommodation bool
	CreateBooking       bool
	UpdateBooking       bool
	DeleteBooking       bool
	CreateRating        bool
	UpdateRating        bool
	DeleteRating        bool
	CreateComment       bool
	UpdateComment       bool
	DeleteComment       bool
	Users               []User
}
