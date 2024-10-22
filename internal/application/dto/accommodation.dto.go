package dto

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type AccommodationDTO struct {
	ID               string   `json:"id"`
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	Cover            string   `json:"cover"`
	Pictures         []string `json:"pictures"`
	Adress           string   `json:"adress"`
	AdditionalAdress *string  `json:"additionalAdress"`
	Zip              string   `json:"zip"`
	City             string   `json:"city"`
	Country          string   `json:"country"`
	Active           bool     `json:"active"`
	Equipments       []string `json:"equipments"`
	Tags             []string `json:"tags"`
	UserID           string   `json:"userId"`
}

type AccommodationInputDTO struct {
	Title            string   `json:"title"`
	Description      string   `json:"description"`
	Cover            string   `json:"cover"`
	Pictures         []string `json:"pictures"`
	Adress           string   `json:"adress"`
	AdditionalAdress *string  `json:"additionalAdress"`
	Zip              string   `json:"zip"`
	City             string   `json:"city"`
	Country          string   `json:"country"`
	Active           bool     `json:"active"`
	Equipments       []string `json:"equipments"`
	Tags             []string `json:"tags"`
	UserID           string   `json:"userId"`
}

func PipeAccommodationInDTO(source *entities.Accommodation, target *AccommodationDTO) {
	target.ID = source.ID
	target.Title = source.Title
	target.Description = source.Description
	target.Cover = source.Cover
	target.Pictures = source.Pictures
	target.Adress = source.Adress
	target.AdditionalAdress = source.AdditionalAdress
	target.Zip = source.Zip
	target.City = source.City
	target.Country = source.Country
	target.Active = source.Active
	target.Equipments = source.Equipments
	target.Tags = source.Tags
	target.UserID = source.UserID
}

func PipeInputDTOInAccommodation(source *AccommodationInputDTO, target *entities.Accommodation) {
	target.Title = source.Title
	target.Description = source.Description
	target.Cover = source.Cover
	target.Pictures = source.Pictures
	target.Adress = source.Adress
	target.AdditionalAdress = source.AdditionalAdress
	target.Zip = source.Zip
	target.City = source.City
	target.Country = source.Country
	target.Active = source.Active
	target.Equipments = source.Equipments
	target.Tags = source.Tags
	target.UserID = source.UserID
}
