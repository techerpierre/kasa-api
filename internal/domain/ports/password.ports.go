package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type PasswordOutput interface {
	Hash(password string) (string, *entities.Exception)
	Compare(plainPassword, hash string) (bool, *entities.Exception)
}
