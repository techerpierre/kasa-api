package services

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type BookingService struct {
	output ports.BookingOutput
}

func CreateBookingService(output ports.BookingOutput) *BookingService {
	return &BookingService{
		output: output,
	}
}

func (s *BookingService) Create(data entities.Booking) (entities.Booking, *entities.Exception) {
	return s.output.Create(data)
}

func (s *BookingService) Update(id string, data entities.Booking) (entities.Booking, *entities.Exception) {
	return s.output.Update(id, data)
}

func (s *BookingService) Delete(id string) *entities.Exception {
	return s.output.Delete(id)
}

func (s *BookingService) List(listing entities.Listing) ([]entities.Booking, int64, *entities.Exception) {
	return s.output.List(listing)
}

func (s *BookingService) FindOne(id string) (entities.Booking, *entities.Exception) {
	return s.output.FindOne(id)
}
