package services

import (
	"mime/multipart"

	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type StorageService struct {
	output ports.StorageOutput
}

func CreateStorageService(output ports.StorageOutput) *StorageService {
	return &StorageService{
		output: output,
	}
}

func (s *StorageService) Write(path string, file multipart.File, handler *multipart.FileHeader) (string, *entities.Exception) {
	return s.output.Write(path, file, handler)
}

func (s *StorageService) Read(path string) ([]byte, *entities.Exception) {
	return s.output.Read(path)
}
