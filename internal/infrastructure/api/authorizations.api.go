package api

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/services"
)

type AuthorizationAPI struct {
	service *services.AuthorizationsService
}

func CreateAuthorizationAPI(service *services.AuthorizationsService) *AuthorizationAPI {
	return &AuthorizationAPI{
		service: service,
	}
}

func (a *AuthorizationAPI) Create(data entities.Authorizations) (entities.Authorizations, *entities.Exception) {
	return a.service.Create(data)
}

func (a *AuthorizationAPI) Update(id string, data entities.Authorizations) (entities.Authorizations, *entities.Exception) {
	return a.service.Update(id, data)
}

func (a *AuthorizationAPI) Delete(id string) *entities.Exception {
	return a.service.Delete(id)
}

func (a *AuthorizationAPI) List() ([]entities.Authorizations, int, *entities.Exception) {
	return a.service.List()
}

func (a *AuthorizationAPI) FindOne(id string) (entities.Authorizations, *entities.Exception) {
	return a.service.FindOne(id)
}
