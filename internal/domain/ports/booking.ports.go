package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type BookingRepo interface {
	Create(data entities.Booking) (entities.Booking, *entities.Exception)
	Update(id string, data entities.Booking) (entities.Booking, *entities.Exception)
	Delete(id string) *entities.Exception
	List(listing entities.Listing) ([]entities.Booking, int64, *entities.Exception)
	FindOne(id string) (entities.Booking, *entities.Exception)
}
