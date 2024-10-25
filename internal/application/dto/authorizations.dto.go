package dto

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type AuthorizationsDTO struct {
	ID string `json:"id"`
}

type AuthorizationsInputDTO struct {
	ID string `json:"id"`
}

func PipeAuthorizationsInDTO(source *entities.Authorizations, target *AuthorizationsDTO) {
	target.ID = source.ID
}

func PipeInputDTOInAuthorizations(source *AuthorizationsInputDTO, target *entities.Authorizations) {
	target.ID = source.ID
}
