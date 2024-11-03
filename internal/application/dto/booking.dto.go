package dto

import (
	"time"

	"github.com/techerpierre/kasa-api/internal/domain/entities"
)

type BookingDTO struct {
	ID              string    `json:"id"`
	Start           time.Time `json:"start"`
	End             time.Time `json:"end"`
	AccommodationID string    `json:"accommodationId"`
	ClientID        string    `json:"clientId"`
}

type BookingInputDTO struct {
	Start           time.Time `json:"start"`
	End             time.Time `json:"end"`
	AccommodationID string    `json:"accommodationId"`
	ClientID        string    `json:"clientId"`
}

type BookingFiltersDTO struct {
	Start           *string `form:"start"`
	End             *string `form:"end"`
	AccommodationID *string `form:"accommodationId"`
	ClientID        *string `form:"clientId"`
}

func PipeBookingInDTO(source *entities.Booking, target *BookingDTO) {
	target.ID = source.ID
	target.Start = source.Start
	target.End = source.End
	target.AccommodationID = source.AccommodationID
	target.ClientID = source.ClientID
}

func PipeInputDTOInBooking(source *BookingInputDTO, target *entities.Booking) {
	target.Start = source.Start
	target.End = source.End
	target.AccommodationID = source.AccommodationID
	target.ClientID = source.ClientID
}

func MakeBookingFilters(filterDTO BookingFiltersDTO) []entities.Filter {
	return []entities.Filter{
		{Field: "start", Value: filterDTO.Start},
		{Field: "end", Value: filterDTO.End},
		{Field: "accommodationId", Value: filterDTO.AccommodationID},
		{Field: "clientId", Value: filterDTO.ClientID},
	}
}
