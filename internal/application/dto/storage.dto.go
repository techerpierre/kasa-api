package dto

import "strings"

type StorageEndpoint string
type UploadMode string

const (
	STORAGE_ENDPOINT_USERS_COVER             StorageEndpoint = "users/cover"
	STORAGE_ENDPOINT_ACCOMMODATIONS_COVER    StorageEndpoint = "accommodations/cover"
	STORAGE_ENDPOINT_ACCOMMODATIONS_PICTURES StorageEndpoint = "accommodations/pictures"
	UPLOAD_MODE_MULTIPLE                     UploadMode      = "multiple"
	UPLOAD_MODE_SINGLE                       UploadMode      = "single"
)

type CreateStorageInput struct {
	Endpoint    StorageEndpoint `json:"endpoint"`
	UploadMode  UploadMode      `json:"uploadMode"`
	Indentifier string          `json:"indentifier"`
}

func IsValidStorageEndpoint(value string) bool {
	validEnpoints := []StorageEndpoint{STORAGE_ENDPOINT_USERS_COVER, STORAGE_ENDPOINT_ACCOMMODATIONS_COVER, STORAGE_ENDPOINT_ACCOMMODATIONS_PICTURES}
	for _, endpoint := range validEnpoints {
		if strings.Compare(string(endpoint), value) == 0 {
			return true
		}
	}
	return false
}
