package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type AuthorizationsOutput interface {
	Create(data entities.Authorizations) (entities.Authorizations, *entities.Exception)
	Update(id entities.AuthStatus, data entities.Authorizations) (entities.Authorizations, *entities.Exception)
	Delete(id entities.AuthStatus) *entities.Exception
	List() ([]entities.Authorizations, int64, *entities.Exception)
	FindOne(id entities.AuthStatus) (entities.Authorizations, *entities.Exception)
}
