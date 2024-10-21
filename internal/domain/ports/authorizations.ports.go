package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type AuthorizationsRepo interface {
	Create(data entities.Authorizations) (entities.Authorizations, *entities.HttpException)
	Update(id entities.AuthStatus, data entities.Authorizations) (entities.Authorizations, *entities.HttpException)
	Delete(id entities.AuthStatus) *entities.HttpException
	List() ([]entities.Authorizations, int64, *entities.HttpException)
	FindOne(id entities.AuthStatus) (entities.Authorizations, *entities.HttpException)
}
