package domain

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type AuthorizationsService struct {
	repo ports.AuthorizationsRepo
}

func CreateAuthorizationsService(repo ports.AuthorizationsRepo) *AuthorizationsService {
	return &AuthorizationsService{
		repo: repo,
	}
}

func (s *AuthorizationsService) Create(data entities.Authorizations) (entities.Authorizations, *entities.HttpException) {
	return s.repo.Create(data)
}

func (s *AuthorizationsService) Update(id entities.AuthStatus, data entities.Authorizations) (entities.Authorizations, *entities.HttpException) {
	return s.repo.Update(id, data)
}

func (s *AuthorizationsService) Delete(id entities.AuthStatus) *entities.HttpException {
	return s.repo.Delete(id)
}

func (s *AuthorizationsService) List() ([]entities.Authorizations, int64, *entities.HttpException) {
	return s.repo.List()
}

func (s *AuthorizationsService) FindOne(id entities.AuthStatus) (entities.Authorizations, *entities.HttpException) {
	return s.repo.FindOne(id)
}
