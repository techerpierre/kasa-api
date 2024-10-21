package services

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type UserService struct {
	output ports.UserOutput
}

func CreateUserService(output ports.UserOutput) *UserService {
	return &UserService{
		output: output,
	}
}

func (s *UserService) Create(data entities.User) (entities.User, *entities.Exception) {
	return s.output.Create(data)
}

func (s *UserService) Update(id string, data entities.User) (entities.User, *entities.Exception) {
	return s.output.Update(id, data)
}

func (s *UserService) Delete(id string) *entities.Exception {
	return s.output.Delete(id)
}

func (s *UserService) List(listing entities.Listing) ([]entities.User, int64, *entities.Exception) {
	return s.output.List(listing)
}

func (s *UserService) FindOne(id string) (entities.User, *entities.Exception) {
	return s.output.FindOne(id)
}
