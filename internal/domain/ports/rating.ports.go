package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type RatingOutput interface {
	Create(data entities.Rating) (entities.Rating, *entities.Exception)
	Update(id string, data entities.Rating) (entities.Rating, *entities.Exception)
	Delete(id string) *entities.Exception
	List(listing entities.Listing) ([]entities.Rating, int64, *entities.Exception)
	FindOne(id string) (entities.Rating, *entities.Exception)
}

type RatingInput interface {
	Create(data entities.Rating) (entities.Rating, *entities.Exception)
	Update(id string, data entities.Rating) (entities.Rating, *entities.Exception)
	Delete(id string) *entities.Exception
	List(listing entities.Listing) ([]entities.Rating, int64, *entities.Exception)
	FindOne(id string) (entities.Rating, *entities.Exception)
}
