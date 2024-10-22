package services

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type RatingService struct {
	output ports.RatingOutput
}

func CreateRatingService(output ports.RatingOutput) *RatingService {
	return &RatingService{
		output: output,
	}
}

func (s *RatingService) Create(data entities.Rating) (entities.Rating, *entities.Exception) {
	return s.output.Create(data)
}

func (s *RatingService) Update(id string, data entities.Rating) (entities.Rating, *entities.Exception) {
	return s.output.Update(id, data)
}

func (s *RatingService) Delete(id string) *entities.Exception {
	return s.output.Delete(id)
}

func (s *RatingService) List(listing entities.Listing) ([]entities.Rating, int, *entities.Exception) {
	return s.output.List(listing)
}

func (s *RatingService) FindOne(id string) (entities.Rating, *entities.Exception) {
	return s.output.FindOne(id)
}
