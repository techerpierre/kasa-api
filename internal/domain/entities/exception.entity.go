package entities

type ExceptionCode int

const (
	ExceptionCode_RessourceNotFound ExceptionCode = 1
	ExceptionCode_BadInputFormat    ExceptionCode = 2
	ExceptionCode_NotAllowedScope   ExceptionCode = 3
)

const (
	ExceptionMessage__RessourceNotFound string = "Ressource not found."
	ExceptionMessage__BadInputFormat    string = "Bad input format."
	ExceptionMessage__NotAllowedScope   string = "Not allowed scope."
)

type Exception struct {
	Code    ExceptionCode
	Message string
}

func CreateException(code ExceptionCode, message string) Exception {
	return Exception{
		Code:    code,
		Message: message,
	}
}
