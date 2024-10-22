package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type AccommodationOutput interface {
	Create(data entities.Accommodation) (entities.Accommodation, *entities.Exception)
	Update(id string, data entities.Accommodation) (entities.Accommodation, *entities.Exception)
	Delete(id string) *entities.Exception
	List(listing entities.Listing) ([]entities.Accommodation, int, *entities.Exception)
	FindOne(id string) (entities.Accommodation, *entities.Exception)
}

type AccommodationInput interface {
	Create(data entities.Accommodation) (entities.Accommodation, *entities.Exception)
	Update(id string, data entities.Accommodation) (entities.Accommodation, *entities.Exception)
	Delete(id string) *entities.Exception
	List(listing entities.Listing) ([]entities.Accommodation, int, *entities.Exception)
	FindOne(id string) (entities.Accommodation, *entities.Exception)
}
