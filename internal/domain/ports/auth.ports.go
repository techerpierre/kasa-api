package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type AuthInput interface {
	Login(email string, password string) (string, *entities.Exception)
}
