package api

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/services"
)

type BookingAPI struct {
	service *services.BookingService
}

func CreateBookingAPI(service *services.BookingService) *BookingAPI {
	return &BookingAPI{
		service: service,
	}
}

func (a *BookingAPI) Create(data entities.Booking) (entities.Booking, *entities.Exception) {
	return a.service.Create(data)
}

func (a *BookingAPI) Update(id string, data entities.Booking) (entities.Booking, *entities.Exception) {
	return a.service.Update(id, data)
}

func (a *BookingAPI) Delete(id string) *entities.Exception {
	return a.service.Delete(id)
}

func (a *BookingAPI) List(listing entities.Listing) ([]entities.Booking, int64, *entities.Exception) {
	return a.service.List(listing)
}

func (a *BookingAPI) FindOne(id string) (entities.Booking, *entities.Exception) {
	return a.service.FindOne(id)
}
