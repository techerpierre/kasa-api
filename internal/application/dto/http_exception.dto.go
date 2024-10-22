package dto

import (
	"net/http"

	"github.com/techerpierre/kasa-api/internal/domain/entities"
)

type HTTPException struct {
	Message   string                 `json:"message"`
	ErrorCode entities.ExceptionCode `json:"errorCode"`
}

func getHttpCode(exceptionCode entities.ExceptionCode) int {
	switch exceptionCode {
	case entities.ExceptionCode_RessourceNotFound:
		return http.StatusNotFound
	case entities.ExceptionCode_BadInputFormat:
		return http.StatusBadRequest
	case entities.ExceptionCode_NotAllowedScope:
		return http.StatusMethodNotAllowed
	default:
		return http.StatusInternalServerError
	}
}

func HTTPExceptionFromException(exception *entities.Exception) (HTTPException, int) {
	return HTTPException{
		Message:   exception.Message,
		ErrorCode: exception.Code,
	}, getHttpCode(exception.Code)
}
