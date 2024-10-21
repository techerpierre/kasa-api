package domain

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type BookingService struct {
	repo ports.BookingRepo
}

func CreateBookingService(repo ports.BookingRepo) *BookingService {
	return &BookingService{
		repo: repo,
	}
}

func (s *BookingService) Create(data entities.Booking) (entities.Booking, *entities.HttpException) {
	return s.repo.Create(data)
}

func (s *BookingService) Update(id string, data entities.Booking) (entities.Booking, *entities.HttpException) {
	return s.repo.Update(id, data)
}

func (s *BookingService) Delete(id string) *entities.HttpException {
	return s.repo.Delete(id)
}

func (s *BookingService) List(listing entities.Listing) ([]entities.Booking, int64, *entities.HttpException) {
	return s.repo.List(listing)
}

func (s *BookingService) FindOne(id string) (entities.Booking, *entities.HttpException) {
	return s.repo.FindOne(id)
}
