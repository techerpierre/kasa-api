package dto

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type RatingDTO struct {
	ID              string `json:"id"`
	Value           int    `json:"value"`
	AccommodationID string `json:"accomodationId"`
	UserID          string `json:"userId"`
}

type RatingInputDTO struct {
	Value           int    `json:"value"`
	AccommodationID string `json:"accomodationId"`
	UserID          string `json:"userId"`
}

type RatingFiltersDTO struct {
	AccomodationID *string `form:"accommodationId"`
	UserID         *string `form:"userId"`
}

func PipeRatingInDTO(source *entities.Rating, target *RatingDTO) {
	target.ID = source.ID
	target.Value = source.Value
	target.AccommodationID = source.AccommodationID
	target.UserID = source.UserID
}

func PipeInputDTOInRating(source *RatingInputDTO, target *entities.Rating) {
	target.Value = source.Value
	target.AccommodationID = source.AccommodationID
	target.UserID = source.UserID
}

func MakeRatingFilters(filterDTO RatingFiltersDTO) []entities.Filter {
	return []entities.Filter{
		{Field: "accommodationId", Value: filterDTO.AccomodationID},
		{Field: "userId", Value: filterDTO.UserID},
	}
}
