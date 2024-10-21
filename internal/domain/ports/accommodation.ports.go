package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type AccommodationRepo interface {
	Create(data entities.Accommodation) (entities.Accommodation, *entities.HttpException)
	Update(id string, data entities.Accommodation) (entities.Accommodation, *entities.HttpException)
	Delete(id string) *entities.HttpException
	List(listing entities.Listing) ([]entities.Accommodation, int64, *entities.HttpException)
	FindOne(id string) (entities.Accommodation, *entities.HttpException)
}
