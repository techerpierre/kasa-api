package ports

import (
	"mime/multipart"

	"github.com/techerpierre/kasa-api/internal/domain/entities"
)

type StorageOutput interface {
	Write(path string, file multipart.File, handler *multipart.FileHeader) (string, *entities.Exception)
	Read(path string) ([]byte, *entities.Exception)
}

type StorageInput interface {
	Write(path string, file multipart.File, handler *multipart.FileHeader) (string, *entities.Exception)
	Read(path string) ([]byte, *entities.Exception)
}
