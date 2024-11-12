package repositories

import (
	"context"
	"strconv"

	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/helpers"
	db "github.com/techerpierre/kasa-api/models"
)

type AuthorizationsRepository struct {
	prisma *db.PrismaClient
}

func CreateAuthorizationsRepository(prisma *db.PrismaClient) *AuthorizationsRepository {
	return &AuthorizationsRepository{
		prisma: prisma,
	}
}

func (r *AuthorizationsRepository) Create(data entities.Authorizations) (entities.Authorizations, *entities.Exception) {

	result, err := r.prisma.Authorizations.CreateOne(
		db.Authorizations.ID.SetIfPresent(helpers.NilIfEmptyString(data.ID)),
		db.Authorizations.CreateAuthorization.Set(data.CreateAuthorization),
		db.Authorizations.UpdateAuthorization.Set(data.UpdateAuthorization),
		db.Authorizations.DeleteAuthorization.Set(data.UpdateAuthorization),
		db.Authorizations.CreateUser.Set(data.CreateUser),
		db.Authorizations.UpdateUser.Set(data.UpdateUser),
		db.Authorizations.DeleteUser.Set(data.DeleteUser),
		db.Authorizations.CreateAccommodation.Set(data.CreateAccommodation),
		db.Authorizations.UpdateAccommodation.Set(data.UpdateAccommodation),
		db.Authorizations.DeleteAccommodation.Set(data.DeleteAccommodation),
		db.Authorizations.CreateBooking.Set(data.CreateBooking),
		db.Authorizations.UpdateBooking.Set(data.UpdateBooking),
		db.Authorizations.DeleteBooking.Set(data.DeleteBooking),
		db.Authorizations.CreateRating.Set(data.CreateRating),
		db.Authorizations.UpdateRating.Set(data.UpdateRating),
		db.Authorizations.DeleteRating.Set(data.DeleteRating),
		db.Authorizations.CreateComment.Set(data.CreateComment),
		db.Authorizations.UpdateComment.Set(data.UpdateComment),
		db.Authorizations.DeleteComment.Set(data.DeleteComment),
	).Exec(context.Background())

	if err != nil {
		return entities.Authorizations{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.Authorizations{
		ID:                  result.ID,
		CreateAuthorization: result.CreateAuthorization,
		UpdateAuthorization: result.UpdateAuthorization,
		DeleteAuthorization: result.DeleteAuthorization,
		CreateUser:          result.CreateUser,
		UpdateUser:          result.UpdateUser,
		DeleteUser:          result.DeleteUser,
		CreateAccommodation: result.CreateAccommodation,
		UpdateAccommodation: result.UpdateAccommodation,
		DeleteAccommodation: result.DeleteAccommodation,
		CreateBooking:       result.CreateBooking,
		UpdateBooking:       result.UpdateBooking,
		DeleteBooking:       result.DeleteBooking,
		CreateRating:        result.CreateRating,
		UpdateRating:        result.UpdateRating,
		DeleteRating:        result.DeleteRating,
		CreateComment:       result.CreateComment,
		UpdateComment:       result.UpdateComment,
		DeleteComment:       result.DeleteComment,
	}, nil
}

func (r *AuthorizationsRepository) Update(id string, data entities.Authorizations) (entities.Authorizations, *entities.Exception) {
	result, err := r.prisma.Authorizations.FindUnique(
		db.Authorizations.ID.Equals(id),
	).Update(
		db.Authorizations.CreateAuthorization.Set(data.CreateAuthorization),
		db.Authorizations.UpdateAuthorization.Set(data.UpdateAuthorization),
		db.Authorizations.DeleteAuthorization.Set(data.UpdateAuthorization),
		db.Authorizations.CreateUser.Set(data.CreateUser),
		db.Authorizations.UpdateUser.Set(data.UpdateUser),
		db.Authorizations.DeleteUser.Set(data.DeleteUser),
		db.Authorizations.CreateAccommodation.Set(data.CreateAccommodation),
		db.Authorizations.UpdateAccommodation.Set(data.UpdateAccommodation),
		db.Authorizations.DeleteAccommodation.Set(data.DeleteAccommodation),
		db.Authorizations.CreateBooking.Set(data.CreateBooking),
		db.Authorizations.UpdateBooking.Set(data.UpdateBooking),
		db.Authorizations.DeleteBooking.Set(data.DeleteBooking),
		db.Authorizations.CreateRating.Set(data.CreateRating),
		db.Authorizations.UpdateRating.Set(data.UpdateRating),
		db.Authorizations.DeleteRating.Set(data.DeleteRating),
		db.Authorizations.CreateComment.Set(data.CreateComment),
		db.Authorizations.UpdateComment.Set(data.UpdateComment),
		db.Authorizations.DeleteComment.Set(data.DeleteComment),
	).Exec(context.Background())

	if err != nil {
		if err == db.ErrNotFound {
			return entities.Authorizations{}, entities.CreateException(
				entities.ExceptionCode_RessourceNotFound,
				entities.ExceptionMessage_RessourceNotFound,
			)
		}

		return entities.Authorizations{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.Authorizations{
		ID:                  result.ID,
		CreateAuthorization: result.CreateAuthorization,
		UpdateAuthorization: result.UpdateAuthorization,
		DeleteAuthorization: result.DeleteAuthorization,
		CreateUser:          result.CreateUser,
		UpdateUser:          result.UpdateUser,
		DeleteUser:          result.DeleteUser,
		CreateAccommodation: result.CreateAccommodation,
		UpdateAccommodation: result.UpdateAccommodation,
		DeleteAccommodation: result.DeleteAccommodation,
		CreateBooking:       result.CreateBooking,
		UpdateBooking:       result.UpdateBooking,
		DeleteBooking:       result.DeleteBooking,
		CreateRating:        result.CreateRating,
		UpdateRating:        result.UpdateRating,
		DeleteRating:        result.DeleteRating,
		CreateComment:       result.CreateComment,
		UpdateComment:       result.UpdateComment,
		DeleteComment:       result.DeleteComment,
	}, nil
}

func (r *AuthorizationsRepository) Delete(id string) *entities.Exception {
	_, err := r.prisma.Authorizations.FindUnique(
		db.Authorizations.ID.Equals(id),
	).Delete().Exec(context.Background())

	if err != nil {
		if err == db.ErrNotFound {
			return entities.CreateException(
				entities.ExceptionCode_RessourceNotFound,
				entities.ExceptionMessage_RessourceNotFound,
			)
		}

		return entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	return nil
}

func (r *AuthorizationsRepository) List() ([]entities.Authorizations, int, *entities.Exception) {
	results, err := r.prisma.Authorizations.FindMany().Exec(context.Background())

	if err != nil {
		return nil, 0, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	var countResult []CountResult

	err = r.prisma.Prisma.QueryRaw(
		`SELECT COUNT(*) FROM "Authorizations"`,
	).Exec(context.Background(), &countResult)

	if err != nil {
		return nil, 0, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	count, _ := strconv.Atoi(countResult[0].Count)

	var authorizations []entities.Authorizations

	for _, result := range results {
		authorizations = append(authorizations, entities.Authorizations{
			ID:                  result.ID,
			CreateAuthorization: result.CreateAuthorization,
			UpdateAuthorization: result.UpdateAuthorization,
			DeleteAuthorization: result.DeleteAuthorization,
			CreateUser:          result.CreateUser,
			UpdateUser:          result.UpdateUser,
			DeleteUser:          result.DeleteUser,
			CreateAccommodation: result.CreateAccommodation,
			UpdateAccommodation: result.UpdateAccommodation,
			DeleteAccommodation: result.DeleteAccommodation,
			CreateBooking:       result.CreateBooking,
			UpdateBooking:       result.UpdateBooking,
			DeleteBooking:       result.DeleteBooking,
			CreateRating:        result.CreateRating,
			UpdateRating:        result.UpdateRating,
			DeleteRating:        result.DeleteRating,
			CreateComment:       result.CreateComment,
			UpdateComment:       result.UpdateComment,
			DeleteComment:       result.DeleteComment,
		})
	}

	return authorizations, count, nil
}

func (r *AuthorizationsRepository) FindOne(id string) (entities.Authorizations, *entities.Exception) {
	result, err := r.prisma.Authorizations.FindUnique(
		db.Authorizations.ID.Equals(id),
	).Exec(context.Background())

	if err != nil {
		if err == db.ErrNotFound {
			return entities.Authorizations{}, entities.CreateException(
				entities.ExceptionCode_RessourceNotFound,
				entities.ExceptionMessage_RessourceNotFound,
			)
		}

		return entities.Authorizations{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.Authorizations{
		ID:                  result.ID,
		CreateAuthorization: result.CreateAuthorization,
		UpdateAuthorization: result.UpdateAuthorization,
		DeleteAuthorization: result.DeleteAuthorization,
		CreateUser:          result.CreateUser,
		UpdateUser:          result.UpdateUser,
		DeleteUser:          result.DeleteUser,
		CreateAccommodation: result.CreateAccommodation,
		UpdateAccommodation: result.UpdateAccommodation,
		DeleteAccommodation: result.DeleteAccommodation,
		CreateBooking:       result.CreateBooking,
		UpdateBooking:       result.UpdateBooking,
		DeleteBooking:       result.DeleteBooking,
		CreateRating:        result.CreateRating,
		UpdateRating:        result.UpdateRating,
		DeleteRating:        result.DeleteRating,
		CreateComment:       result.CreateComment,
		UpdateComment:       result.UpdateComment,
		DeleteComment:       result.DeleteComment,
	}, nil
}
