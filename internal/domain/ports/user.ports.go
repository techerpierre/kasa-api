package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type UserRepo interface {
	Create(data entities.User) (entities.User, *entities.HttpException)
	Update(id string, data entities.User) (entities.User, *entities.HttpException)
	Delete(id string) *entities.HttpException
	List(listing entities.Listing) ([]entities.User, int64, *entities.HttpException)
	FindOne(id string) (entities.User, *entities.HttpException)
	FindOneByEmail(email string) (entities.User, *entities.HttpException)
}
