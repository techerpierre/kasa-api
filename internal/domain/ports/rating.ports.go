package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type RatingRepo interface {
	Create(data entities.Rating) (entities.Rating, *entities.HttpException)
	Update(id string, data entities.Rating) (entities.Rating, *entities.HttpException)
	Delete(id string) *entities.HttpException
	List(listing entities.Listing) ([]entities.Rating, int64, *entities.HttpException)
	FindOne(id string) (entities.Rating, *entities.HttpException)
}
