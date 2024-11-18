package repositories

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/techerpierre/kasa-api/internal/domain/entities"
)

type StorageRepository struct{}

func CreateStorageRepository() *StorageRepository {
	return &StorageRepository{}
}

func (s *StorageRepository) Write(path string, file multipart.File, handler *multipart.FileHeader) (string, *entities.Exception) {
	out, err := os.Create(fmt.Sprintf("%s%s", os.Getenv("FILE_STORE_PATH"), path))
	if err != nil {
		fmt.Println("Error in file creation:", err)
		return "", entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		fmt.Println("Error in file copy:", err)
		return "", entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	return path, nil
}

func (s *StorageRepository) Read(path string) ([]byte, *entities.Exception) {
	file, err := os.ReadFile(fmt.Sprintf("%s%s", os.Getenv("FILE_STORE_PATH"), path))
	if err != nil {
		fmt.Println("Error in file copy:", err)
		if err == os.ErrNotExist {
			return nil, entities.CreateException(
				entities.ExceptionCode_RessourceNotFound,
				entities.ExceptionMessage_RessourceNotFound,
			)
		}
		return nil, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	return file, nil
}
