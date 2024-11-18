package api

import (
	"mime/multipart"

	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/services"
)

type StorageAPI struct {
	sevice *services.StorageService
}

func CreateStorageAPI(sevice *services.StorageService) *StorageAPI {
	return &StorageAPI{
		sevice: sevice,
	}
}

func (a *StorageAPI) Write(path string, file multipart.File, handler *multipart.FileHeader) (string, *entities.Exception) {
	return a.sevice.Write(path, file, handler)
}

func (a *StorageAPI) Read(path string) ([]byte, *entities.Exception) {
	return a.sevice.Read(path)
}
