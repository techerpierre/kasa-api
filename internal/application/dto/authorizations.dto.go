package dto

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type AuthorizationsDTO struct {
	ID                  string `json:"id"`
	CreateAuthorization bool   `json:"createAuthorization"`
	UpdateAuthorization bool   `json:"updateAuthorization"`
	DeleteAuthorization bool   `json:"deleteAuthorization"`
	CreateUser          bool   `json:"createUser"`
	UpdateUser          bool   `json:"updateUser"`
	DeleteUser          bool   `json:"deleteUser"`
	CreateAccommodation bool   `json:"createAccommodation"`
	UpdateAccommodation bool   `json:"updateAccommodation"`
	DeleteAccommodation bool   `json:"deleteAccommodation"`
	CreateBooking       bool   `json:"createBooking"`
	UpdateBooking       bool   `json:"updateBooking"`
	DeleteBooking       bool   `json:"deleteBooking"`
	CreateRating        bool   `json:"createRating"`
	UpdateRating        bool   `json:"updateRating"`
	DeleteRating        bool   `json:"deleteRating"`
	CreateComment       bool   `json:"createComment"`
	UpdateComment       bool   `json:"updateComment"`
	DeleteComment       bool   `json:"deleteComment"`
}

type AuthorizationsInputDTO struct {
	ID                  string `json:"id"`
	CreateAuthorization bool   `json:"createAuthorization"`
	UpdateAuthorization bool   `json:"updateAuthorization"`
	DeleteAuthorization bool   `json:"deleteAuthorization"`
	CreateUser          bool   `json:"createUser"`
	UpdateUser          bool   `json:"updateUser"`
	DeleteUser          bool   `json:"deleteUser"`
	CreateAccommodation bool   `json:"createAccommodation"`
	UpdateAccommodation bool   `json:"updateAccommodation"`
	DeleteAccommodation bool   `json:"deleteAccommodation"`
	CreateBooking       bool   `json:"createBooking"`
	UpdateBooking       bool   `json:"updateBooking"`
	DeleteBooking       bool   `json:"deleteBooking"`
	CreateRating        bool   `json:"createRating"`
	UpdateRating        bool   `json:"updateRating"`
	DeleteRating        bool   `json:"deleteRating"`
	CreateComment       bool   `json:"createComment"`
	UpdateComment       bool   `json:"updateComment"`
	DeleteComment       bool   `json:"deleteComment"`
}

func PipeAuthorizationsInDTO(source *entities.Authorizations, target *AuthorizationsDTO) {
	target.ID = source.ID
	target.CreateAuthorization = source.CreateAuthorization
	target.UpdateAuthorization = source.UpdateAuthorization
	target.DeleteAuthorization = source.DeleteAuthorization
	target.CreateUser = source.CreateUser
	target.UpdateUser = source.UpdateUser
	target.DeleteUser = source.DeleteUser
	target.CreateAccommodation = source.CreateAccommodation
	target.UpdateAccommodation = source.UpdateAccommodation
	target.DeleteAccommodation = source.DeleteAccommodation
	target.CreateBooking = source.CreateBooking
	target.UpdateBooking = source.UpdateBooking
	target.DeleteBooking = source.DeleteBooking
	target.CreateRating = source.CreateRating
	target.UpdateRating = source.UpdateRating
	target.DeleteRating = source.DeleteRating
	target.CreateComment = source.CreateComment
	target.UpdateComment = source.UpdateComment
	target.DeleteComment = source.DeleteComment
}

func PipeInputDTOInAuthorizations(source *AuthorizationsInputDTO, target *entities.Authorizations) {
	target.ID = source.ID
	target.CreateAuthorization = source.CreateAuthorization
	target.UpdateAuthorization = source.UpdateAuthorization
	target.DeleteAuthorization = source.DeleteAuthorization
	target.CreateUser = source.CreateUser
	target.UpdateUser = source.UpdateUser
	target.DeleteUser = source.DeleteUser
	target.CreateAccommodation = source.CreateAccommodation
	target.UpdateAccommodation = source.UpdateAccommodation
	target.DeleteAccommodation = source.DeleteAccommodation
	target.CreateBooking = source.CreateBooking
	target.UpdateBooking = source.UpdateBooking
	target.DeleteBooking = source.DeleteBooking
	target.CreateRating = source.CreateRating
	target.UpdateRating = source.UpdateRating
	target.DeleteRating = source.DeleteRating
	target.CreateComment = source.CreateComment
	target.UpdateComment = source.UpdateComment
	target.DeleteComment = source.DeleteComment
}
