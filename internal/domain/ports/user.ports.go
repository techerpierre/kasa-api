package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type UserOutput interface {
	Create(data entities.User) (entities.User, *entities.Exception)
	Update(id string, data entities.User) (entities.User, *entities.Exception)
	Delete(id string) *entities.Exception
	List(listing entities.Listing) ([]entities.User, int64, *entities.Exception)
	FindOne(id string) (entities.User, *entities.Exception)
	FindOneByEmail(email string) (entities.User, *entities.Exception)
}
