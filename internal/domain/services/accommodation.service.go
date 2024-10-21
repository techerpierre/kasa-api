package domain

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type AccomodationService struct {
	output ports.AccommodationOutput
}

func CreateAccomodationService(output ports.AccommodationOutput) *AccomodationService {
	return &AccomodationService{
		output: output,
	}
}

func (s *AccomodationService) Create(data entities.Accommodation) (entities.Accommodation, *entities.Exception) {
	return s.output.Create(data)
}

func (s *AccomodationService) Update(id string, data entities.Accommodation) (entities.Accommodation, *entities.Exception) {
	return s.output.Update(id, data)
}

func (s *AccomodationService) Delete(id string) *entities.Exception {
	return s.output.Delete(id)
}

func (s *AccomodationService) List(listing entities.Listing) ([]entities.Accommodation, int64, *entities.Exception) {
	return s.output.List(listing)
}

func (s *AccomodationService) FindOne(id string) (entities.Accommodation, *entities.Exception) {
	return s.output.FindOne(id)
}
