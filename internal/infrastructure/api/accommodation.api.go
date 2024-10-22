package api

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/services"
)

type AccommodationAPI struct {
	service *services.AccomodationService
}

func CreateAccommodationAPI(service *services.AccomodationService) *AccommodationAPI {
	return &AccommodationAPI{
		service: service,
	}
}

func (a *AccommodationAPI) Create(data entities.Accommodation) (entities.Accommodation, *entities.Exception) {
	return a.service.Create(data)
}

func (a *AccommodationAPI) Update(id string, data entities.Accommodation) (entities.Accommodation, *entities.Exception) {
	return a.service.Update(id, data)
}

func (a *AccommodationAPI) Delete(id string) *entities.Exception {
	return a.service.Delete(id)
}

func (a *AccommodationAPI) List(listing entities.Listing) ([]entities.Accommodation, int, *entities.Exception) {
	return a.service.List(listing)
}

func (a *AccommodationAPI) FindOne(id string) (entities.Accommodation, *entities.Exception) {
	return a.service.FindOne(id)
}
