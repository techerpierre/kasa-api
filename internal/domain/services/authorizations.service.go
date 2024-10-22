package services

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type AuthorizationsService struct {
	output ports.AuthorizationsOutput
}

func CreateAuthorizationsService(output ports.AuthorizationsOutput) *AuthorizationsService {
	return &AuthorizationsService{
		output: output,
	}
}

func (s *AuthorizationsService) Create(data entities.Authorizations) (entities.Authorizations, *entities.Exception) {
	return s.output.Create(data)
}

func (s *AuthorizationsService) Update(id string, data entities.Authorizations) (entities.Authorizations, *entities.Exception) {
	return s.output.Update(id, data)
}

func (s *AuthorizationsService) Delete(id string) *entities.Exception {
	return s.output.Delete(id)
}

func (s *AuthorizationsService) List() ([]entities.Authorizations, int, *entities.Exception) {
	return s.output.List()
}

func (s *AuthorizationsService) FindOne(id string) (entities.Authorizations, *entities.Exception) {
	return s.output.FindOne(id)
}
