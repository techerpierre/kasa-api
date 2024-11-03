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

type UserFiltersDTO struct {
	Email            *string `form:"email"`
	Firstname        *string `form:"firstname"`
	Lastname         *string `form:"lastname"`
	AuthorizationsID *string `form:"authorizationsId"`
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

func MakeUserFilters(filterDTO UserFiltersDTO) []entities.Filter {
	return []entities.Filter{
		{Field: "email", Value: filterDTO.Email},
		{Field: "firstname", Value: filterDTO.Firstname},
		{Field: "lastname", Value: filterDTO.Lastname},
		{Field: "authorizationsId", Value: filterDTO.AuthorizationsID},
	}
}
