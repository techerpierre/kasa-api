package repositories

import (
	"context"

	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/helpers"
	db "github.com/techerpierre/kasa-api/models"
)

type UserRepository struct {
	prisma *db.PrismaClient
}

func CreateUserRepository(prisma *db.PrismaClient) *UserRepository {
	return &UserRepository{
		prisma: prisma,
	}
}

func (r *UserRepository) Create(data entities.User) (entities.User, *entities.Exception) {
	result, err := r.prisma.User.CreateOne(
		db.User.Email.SetIfPresent(helpers.NilIfEmptyString(data.Email)),
		db.User.Password.SetIfPresent(helpers.NilIfEmptyString(data.Password)),
		db.User.Firstname.SetIfPresent(helpers.NilIfEmptyString(data.Firstname)),
		db.User.Lastname.SetIfPresent(helpers.NilIfEmptyString(data.Lastname)),
		db.User.Authorizations.Link(
			db.Authorizations.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.AuthorizationsID)),
		),
		db.User.Picture.SetIfPresent(data.Picture),
		db.User.Cover.SetIfPresent(data.Cover),
	).Exec(context.Background())

	if err != nil {
		return entities.User{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.User{
		ID:               result.ID,
		Email:            result.Email,
		Password:         result.Password,
		Firstname:        result.Firstname,
		Lastname:         result.Lastname,
		Picture:          helpers.PointerFromPrismaField(result.Picture()),
		Cover:            helpers.PointerFromPrismaField(result.Cover()),
		AuthorizationsID: result.AuthorizationsID,
	}, nil
}

func (r *UserRepository) Update(id string, data entities.User) (entities.User, *entities.Exception) {
	result, err := r.prisma.User.FindUnique(
		db.User.ID.Equals(id),
	).Update(
		db.User.Email.SetIfPresent(helpers.NilIfEmptyString(data.Email)),
		db.User.Password.SetIfPresent(helpers.NilIfEmptyString(data.Password)),
		db.User.Firstname.SetIfPresent(helpers.NilIfEmptyString(data.Firstname)),
		db.User.Lastname.SetIfPresent(helpers.NilIfEmptyString(data.Lastname)),
		db.User.Authorizations.Link(
			db.Authorizations.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.AuthorizationsID)),
		),
		db.User.Picture.SetIfPresent(data.Picture),
		db.User.Cover.SetIfPresent(data.Cover),
	).Exec(context.Background())

	if err != nil {
		return entities.User{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.User{
		ID:               result.ID,
		Email:            result.Email,
		Password:         result.Password,
		Firstname:        result.Firstname,
		Lastname:         result.Lastname,
		Picture:          helpers.PointerFromPrismaField(result.Picture()),
		Cover:            helpers.PointerFromPrismaField(result.Cover()),
		AuthorizationsID: result.AuthorizationsID,
	}, nil
}

func (r *UserRepository) Delete(id string) *entities.Exception {
	_, err := r.prisma.User.FindUnique(
		db.User.ID.Equals(id),
	).Delete().Exec(context.Background())

	if err != nil {
		return entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return nil
}

func (r *UserRepository) List(listing entities.Listing) ([]entities.User, int, *entities.Exception) {
	results, err := r.prisma.User.FindMany().Skip(
		listing.Page * listing.Pagesize,
	).Take(listing.Pagesize).Exec(context.Background())

	if err != nil {
		return nil, 0, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	countResult, err := r.prisma.Prisma.ExecuteRaw(
		"SELECT COUNT(*) FROM USER",
	).Exec(context.Background())

	if err != nil {
		return nil, 0, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	var users []entities.User

	for _, result := range results {
		users = append(users, entities.User{
			ID:               result.ID,
			Email:            result.Email,
			Password:         result.Password,
			Firstname:        result.Firstname,
			Lastname:         result.Lastname,
			Picture:          helpers.PointerFromPrismaField(result.Picture()),
			Cover:            helpers.PointerFromPrismaField(result.Cover()),
			AuthorizationsID: result.AuthorizationsID,
		})
	}

	return users, countResult.Count, nil
}

func (r *UserRepository) FindOne(id string) (entities.User, *entities.Exception) {
	result, err := r.prisma.User.FindUnique(
		db.User.ID.Equals(id),
	).Exec(context.Background())

	if err != nil {
		return entities.User{}, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	if result == nil {
		return entities.User{}, entities.CreateException(
			entities.ExceptionCode_RessourceNotFound,
			entities.ExceptionMessage_RessourceNotFound,
		)
	}

	return entities.User{
		ID:               result.ID,
		Email:            result.Email,
		Password:         result.Password,
		Firstname:        result.Firstname,
		Lastname:         result.Lastname,
		Picture:          helpers.PointerFromPrismaField(result.Picture()),
		Cover:            helpers.PointerFromPrismaField(result.Cover()),
		AuthorizationsID: result.AuthorizationsID,
	}, nil
}

func (r *UserRepository) FindOneByEmail(email string) (entities.User, *entities.Exception) {
	result, err := r.prisma.User.FindUnique(
		db.User.Email.Equals(email),
	).Exec(context.Background())

	if err != nil {
		return entities.User{}, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	if result == nil {
		return entities.User{}, entities.CreateException(
			entities.ExceptionCode_RessourceNotFound,
			entities.ExceptionMessage_RessourceNotFound,
		)
	}

	return entities.User{
		ID:               result.ID,
		Email:            result.Email,
		Password:         result.Password,
		Firstname:        result.Firstname,
		Lastname:         result.Lastname,
		Picture:          helpers.PointerFromPrismaField(result.Picture()),
		Cover:            helpers.PointerFromPrismaField(result.Cover()),
		AuthorizationsID: result.AuthorizationsID,
	}, nil
}
