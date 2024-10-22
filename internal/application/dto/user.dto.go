package dto

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type UserDTO struct {
	ID               string  `json:"id"`
	Email            string  `json:"email"`
	Firstname        string  `json:"firstname"`
	Lastname         string  `json:"lastname"`
	Picture          *string `json:"picture"`
	Cover            *string `json:"cover"`
	AuthorizationsID string  `json:"authorizationsId"`
}

type UserInputDTO struct {
	Email            string  `json:"email"`
	Password         string  `json:"password"`
	Firstname        string  `json:"firstname"`
	Lastname         string  `json:"lastname"`
	Picture          *string `json:"picture"`
	Cover            *string `json:"cover"`
	AuthorizationsID string  `json:"authorizationsId"`
}

func PipeUserInDTO(source *entities.User, target *UserDTO) {
	target.ID = source.ID
	target.Email = source.Email
	target.Firstname = source.Firstname
	target.Lastname = source.Lastname
	target.Picture = source.Picture
	target.Cover = source.Cover
	target.AuthorizationsID = source.AuthorizationsID
}

func PipeInputDTOInUser(source *UserInputDTO, target *entities.User) {
	target.Email = source.Email
	target.Password = source.Password
	target.Firstname = source.Firstname
	target.Lastname = source.Lastname
	target.Picture = source.Picture
	target.Cover = source.Cover
	target.AuthorizationsID = source.AuthorizationsID
}
