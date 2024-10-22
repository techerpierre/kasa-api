package api

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/services"
)

type RatingAPI struct {
	service *services.RatingService
}

func CreateRatingAPI(service *services.RatingService) *RatingAPI {
	return &RatingAPI{
		service: service,
	}
}

func (a *RatingAPI) Create(data entities.Rating) (entities.Rating, *entities.Exception) {
	return a.service.Create(data)
}

func (a *RatingAPI) Update(id string, data entities.Rating) (entities.Rating, *entities.Exception) {
	return a.service.Update(id, data)
}

func (a *RatingAPI) Delete(id string) *entities.Exception {
	return a.service.Delete(id)
}

func (a *RatingAPI) List(listing entities.Listing) ([]entities.Rating, int, *entities.Exception) {
	return a.service.List(listing)
}

func (a *RatingAPI) FindOne(id string) (entities.Rating, *entities.Exception) {
	return a.service.FindOne(id)
}
