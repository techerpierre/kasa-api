package api

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/services"
)

type UserAPI struct {
	service *services.UserService
}

func CreateUserAPI(service *services.UserService) *UserAPI {
	return &UserAPI{
		service: service,
	}
}

func (a *UserAPI) Create(data entities.User) (entities.User, *entities.Exception) {
	return a.service.Create(data)
}

func (a *UserAPI) Update(id string, data entities.User) (entities.User, *entities.Exception) {
	return a.service.Update(id, data)
}
func (a *UserAPI) Delete(id string) *entities.Exception {
	return a.service.Delete(id)
}

func (a *UserAPI) List(listing entities.Listing) ([]entities.User, int, *entities.Exception) {
	return a.service.List(listing)
}

func (a *UserAPI) FindOne(id string) (entities.User, *entities.Exception) {
	return a.service.FindOne(id)
}
