package domain

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type AccomodationService struct {
	repo ports.AccommodationRepo
}

func CreateAccomodationService(repo ports.AccommodationRepo) *AccomodationService {
	return &AccomodationService{
		repo: repo,
	}
}

func (s *AccomodationService) Create(data entities.Accommodation) (entities.Accommodation, *entities.HttpException) {
	return s.repo.Create(data)
}

func (s *AccomodationService) Update(id string, data entities.Accommodation) (entities.Accommodation, *entities.HttpException) {
	return s.repo.Update(id, data)
}

func (s *AccomodationService) Delete(id string) *entities.HttpException {
	return s.repo.Delete(id)
}

func (s *AccomodationService) List(listing entities.Listing) ([]entities.Accommodation, int64, *entities.HttpException) {
	return s.repo.List(listing)
}

func (s *AccomodationService) FindOne(id string) (entities.Accommodation, *entities.HttpException) {
	return s.repo.FindOne(id)
}
