package services

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type PasswordService struct {
	output ports.PasswordOutput
}

func CreatePasswordService(output ports.PasswordOutput) *PasswordService {
	return &PasswordService{
		output: output,
	}
}

func (s *PasswordService) Hash(password string) (string, *entities.Exception) {
	return s.output.Hash(password)
}

func (s *PasswordService) Compare(plainPassword, hash string) (bool, *entities.Exception) {
	return s.output.Compare(plainPassword, hash)
}
