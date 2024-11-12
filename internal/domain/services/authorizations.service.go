package services

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type AuthorizationsService struct {
	output     ports.AuthorizationsOutput
	jwtService *JwtService
}

func CreateAuthorizationsService(output ports.AuthorizationsOutput, jwtService *JwtService) *AuthorizationsService {
	return &AuthorizationsService{
		output:     output,
		jwtService: jwtService,
	}
}

func (s *AuthorizationsService) Create(data entities.Authorizations) (entities.Authorizations, *entities.Exception) {
	return s.output.Create(data)
}

func (s *AuthorizationsService) Update(id string, data entities.Authorizations) (entities.Authorizations, *entities.Exception) {
	return s.output.Update(id, data)
}

func (s *AuthorizationsService) Delete(id string) *entities.Exception {
	return s.output.Delete(id)
}

func (s *AuthorizationsService) List() ([]entities.Authorizations, int, *entities.Exception) {
	return s.output.List()
}

func (s *AuthorizationsService) FindOne(id string) (entities.Authorizations, *entities.Exception) {
	return s.output.FindOne(id)
}

func (s *AuthorizationsService) IsAuthorized(token string, authorization entities.AuthorizationType) (bool, entities.Payloads, *entities.Exception) {
	payloads, exception := s.jwtService.Verify(token)
	if exception != nil {
		return false, entities.Payloads{}, exception
	}

	authorizations, exception := s.output.FindOne(payloads.AuthorizationsID)
	if exception != nil {
		return false, entities.Payloads{}, exception
	}

	return s.getAuthorizationByType(authorizations, authorization), payloads, nil
}

func (s *AuthorizationsService) getAuthorizationByType(authorizations entities.Authorizations, authorization entities.AuthorizationType) bool {
	if authorization == entities.Authorization_NoAuthorization {
		return true
	}

	authorizationMap := map[entities.AuthorizationType]bool{
		entities.Authorization_CreateAuthorization: authorizations.CreateAuthorization,
		entities.Authorization_UpdateAuthorization: authorizations.UpdateAuthorization,
		entities.Authorization_DeleteAuthorization: authorizations.DeleteAuthorization,
		entities.Authorization_CreateUser:          authorizations.CreateUser,
		entities.Authorization_UpdateUser:          authorizations.UpdateUser,
		entities.Authorization_DeleteUser:          authorizations.DeleteUser,
		entities.Authorization_CreateAccommodation: authorizations.CreateAccommodation,
		entities.Authorization_UpdateAccommodation: authorizations.UpdateAccommodation,
		entities.Authorization_DeleteAccommodation: authorizations.DeleteAccommodation,
		entities.Authorization_CreateBooking:       authorizations.CreateBooking,
		entities.Authorization_UpdateBooking:       authorizations.UpdateBooking,
		entities.Authorization_DeleteBooking:       authorizations.DeleteBooking,
		entities.Authorization_CreateRating:        authorizations.CreateRating,
		entities.Authorization_UpdateRating:        authorizations.UpdateRating,
		entities.Authorization_DeleteRating:        authorizations.DeleteRating,
		entities.Authorization_CreateComment:       authorizations.CreateComment,
		entities.Authorization_UpdateComment:       authorizations.UpdateComment,
		entities.Authorization_DeleteComment:       authorizations.DeleteComment,
	}

	if result, found := authorizationMap[authorization]; found {
		return result
	}
	return false
}
