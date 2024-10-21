package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type AccommodationRepo interface {
	Create(data entities.Accommodation) (entities.Accommodation, *entities.Exception)
	Update(id string, data entities.Accommodation) (entities.Accommodation, *entities.Exception)
	Delete(id string) *entities.Exception
	List(listing entities.Listing) ([]entities.Accommodation, int64, *entities.Exception)
	FindOne(id string) (entities.Accommodation, *entities.Exception)
}