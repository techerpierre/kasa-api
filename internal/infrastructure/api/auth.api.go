package api

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/services"
)

type AuthApi struct {
	service *services.AuthService
}

func CreateAuthApi(service *services.AuthService) *AuthApi {
	return &AuthApi{
		service: service,
	}
}

func (a *AuthApi) Login(email string, password string) (string, *entities.Exception) {
	return a.service.Login(email, password)
}
