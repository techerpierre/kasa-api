package services

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type JwtService struct {
	output ports.JwtOutput
}

func CreateJwtService(output ports.JwtOutput) *JwtService {
	return &JwtService{
		output: output,
	}
}

func (s *JwtService) Sign(payloads entities.Payloads) (string, *entities.Exception) {
	return s.output.Sign(payloads)
}

func (s *JwtService) Verify(token string) (entities.Payloads, *entities.Exception) {
	return s.output.Verify(token)
}
