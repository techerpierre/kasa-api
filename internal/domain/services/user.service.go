package services

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type UserService struct {
	output          ports.UserOutput
	passwordService *PasswordService
}

func CreateUserService(output ports.UserOutput, passwordService *PasswordService) *UserService {
	return &UserService{
		output:          output,
		passwordService: passwordService,
	}
}

func (s *UserService) Create(data entities.User) (entities.User, *entities.Exception) {
	hash, exception := s.passwordService.Hash(data.Password)
	if exception != nil {
		return entities.User{}, exception
	}
	data.Password = hash

	return s.output.Create(data)
}

func (s *UserService) Update(id string, data entities.User) (entities.User, *entities.Exception) {
	return s.output.Update(id, data)
}

func (s *UserService) Delete(id string) *entities.Exception {
	return s.output.Delete(id)
}

func (s *UserService) List(listing entities.Listing) ([]entities.User, int, *entities.Exception) {
	return s.output.List(listing)
}

func (s *UserService) FindOne(id string) (entities.User, *entities.Exception) {
	return s.output.FindOne(id)
}
