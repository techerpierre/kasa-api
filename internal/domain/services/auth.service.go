package services

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type AuthService struct {
	userService     *UserService
	passwordService *PasswordService
	jwtService      *JwtService
}

func CreateAuthService(userService *UserService, passwordService *PasswordService, jwtService *JwtService) *AuthService {
	return &AuthService{
		userService:     userService,
		passwordService: passwordService,
		jwtService:      jwtService,
	}
}

func (s *AuthService) Login(email string, password string) (string, *entities.Exception) {
	user, exception := s.userService.FindOneByEmail(email)
	if exception != nil {
		return "", exception
	}

	ok, exception := s.passwordService.Compare(password, user.Password)
	if exception != nil {
		return "", exception
	}
	if !ok {
		return "", entities.CreateException(
			entities.ExceptionCode_Unauthorized,
			entities.ExceptionMessage_Unauthorized,
		)
	}

	payloads := entities.Payloads{
		ID:               user.ID,
		Email:            user.Email,
		Firstname:        user.Firstname,
		Lastname:         user.Lastname,
		AuthorizationsID: user.AuthorizationsID,
	}

	return s.jwtService.Sign(payloads)
}
