package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type AuthorizationsOutput interface {
	Create(data entities.Authorizations) (entities.Authorizations, *entities.Exception)
	Update(id string, data entities.Authorizations) (entities.Authorizations, *entities.Exception)
	Delete(id string) *entities.Exception
	List() ([]entities.Authorizations, int, *entities.Exception)
	FindOne(id string) (entities.Authorizations, *entities.Exception)
}

type AuthorizationsInput interface {
	Create(data entities.Authorizations) (entities.Authorizations, *entities.Exception)
	Update(id string, data entities.Authorizations) (entities.Authorizations, *entities.Exception)
	Delete(id string) *entities.Exception
	List() ([]entities.Authorizations, int, *entities.Exception)
	FindOne(id string) (entities.Authorizations, *entities.Exception)
}
