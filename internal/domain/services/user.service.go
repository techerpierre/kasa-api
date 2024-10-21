package domain

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type UserService struct {
	repo ports.UserRepo
}

func CreateUserService(repo ports.UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(data entities.User) (entities.User, *entities.Exception) {
	return s.repo.Create(data)
}

func (s *UserService) Update(id string, data entities.User) (entities.User, *entities.Exception) {
	return s.repo.Update(id, data)
}

func (s *UserService) Delete(id string) *entities.Exception {
	return s.repo.Delete(id)
}

func (s *UserService) List(listing entities.Listing) ([]entities.User, int64, *entities.Exception) {
	return s.repo.List(listing)
}

func (s *UserService) FindOne(id string) (entities.User, *entities.Exception) {
	return s.repo.FindOne(id)
}
