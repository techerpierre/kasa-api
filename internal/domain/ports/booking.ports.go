package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type BookingRepo interface {
	Create(data entities.Booking) (entities.Booking, *entities.HttpException)
	Update(id string, data entities.Booking) (entities.Booking, *entities.HttpException)
	Delete(id string) *entities.HttpException
	List(listing entities.Listing) ([]entities.Booking, int64, *entities.HttpException)
	FindOne(id string) (entities.Booking, *entities.HttpException)
}
