package entities

type ExceptionCode int

const (
	ExceptionCode_NotHandledError   ExceptionCode = 0
	ExceptionCode_RessourceNotFound ExceptionCode = 1
	ExceptionCode_BadInputFormat    ExceptionCode = 2
	ExceptionCode_NotAllowedScope   ExceptionCode = 3
	ExceptionCode_Unauthorized      ExceptionCode = 4
)

const (
	ExceptionMessage_NotHandledError   string = "Not handled error."
	ExceptionMessage_RessourceNotFound string = "Ressource not found."
	ExceptionMessage_BadInputFormat    string = "Bad input format."
	ExceptionMessage_NotAllowedScope   string = "Not allowed scope."
	ExceptionMessage_Unauthorized      string = "unauthorized."
)

type Exception struct {
	Code    ExceptionCode
	Message string
}

func CreateException(code ExceptionCode, message string) *Exception {
	return &Exception{
		Code:    code,
		Message: message,
	}
}
