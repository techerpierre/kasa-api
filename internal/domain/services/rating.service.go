package domain

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type RatingService struct {
	repo ports.RatingRepo
}

func CreateRatingService(repo ports.RatingRepo) *RatingService {
	return &RatingService{
		repo: repo,
	}
}

func (s *RatingService) Create(data entities.Rating) (entities.Rating, *entities.HttpException) {
	return s.repo.Create(data)
}

func (s *RatingService) Update(id string, data entities.Rating) (entities.Rating, *entities.HttpException) {
	return s.repo.Update(id, data)
}

func (s *RatingService) Delete(id string) *entities.HttpException {
	return s.repo.Delete(id)
}

func (s *RatingService) List(listing entities.Listing) ([]entities.Rating, int64, *entities.HttpException) {
	return s.repo.List(listing)
}

func (s *RatingService) FindOne(id string) (entities.Rating, *entities.HttpException) {
	return s.repo.FindOne(id)
}
