package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type JwtOutput interface {
	Sign(payloads entities.Payloads) (string, *entities.Exception)
	Verify(token string) (entities.Payloads, *entities.Exception)
}
