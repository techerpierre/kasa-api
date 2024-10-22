package repositories

import (
	"context"

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
		db.Authorizations.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.ID)),
	).Exec(context.Background())

	if err != nil {
		return entities.Authorizations{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.Authorizations{
		ID: result.ID,
	}, nil
}

func (r *AuthorizationsRepository) Update(id string, data entities.Authorizations) (entities.Authorizations, *entities.Exception) {
	result, err := r.prisma.Authorizations.FindUnique(
		db.Authorizations.ID.Equals(id),
	).Update().Exec(context.Background())

	if err != nil {
		return entities.Authorizations{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.Authorizations{
		ID: result.ID,
	}, nil
}

func (r *AuthorizationsRepository) Delete(id string) *entities.Exception {
	_, err := r.prisma.Authorizations.FindUnique(
		db.Authorizations.ID.Equals(id),
	).Delete().Exec(context.Background())

	if err != nil {
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

	countResult, err := r.prisma.Prisma.ExecuteRaw(
		`SELECT COUNT(*) FROM Authorizations`,
	).Exec(context.Background())

	if err != nil {
		return nil, 0, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	var authorizations []entities.Authorizations

	for _, result := range results {
		authorizations = append(authorizations, entities.Authorizations{
			ID: result.ID,
		})
	}

	return authorizations, countResult.Count, nil
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
		ID: result.ID,
	}, nil
}
