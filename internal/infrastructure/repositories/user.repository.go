package repositories

import (
	"context"
	"strconv"

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
		if err == db.ErrNotFound {
			return entities.User{}, entities.CreateException(
				entities.ExceptionCode_RessourceNotFound,
				entities.ExceptionMessage_RessourceNotFound,
			)
		}

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

func (r *UserRepository) List(listing entities.Listing) ([]entities.User, int, *entities.Exception) {
	var filters []db.UserWhereParam

	for _, filter := range listing.Filters {
		filterQuery, exception := r.getFilterQuery(filter)
		if exception != nil {
			return nil, 0, exception
		}
		filters = append(filters, filterQuery)
	}

	results, err := r.prisma.User.FindMany(filters...).Skip(
		listing.Page * listing.Pagesize,
	).Take(listing.Pagesize).Exec(context.Background())

	if err != nil {
		return nil, 0, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	var countResult []CountResult

	err = r.prisma.Prisma.QueryRaw(
		`SELECT COUNT(*) FROM "User"`,
	).Exec(context.Background(), &countResult)

	if err != nil {
		return nil, 0, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	count, _ := strconv.Atoi(countResult[0].Count)

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

	return users, count, nil
}

func (r *UserRepository) FindOne(id string) (entities.User, *entities.Exception) {
	result, err := r.prisma.User.FindUnique(
		db.User.ID.Equals(id),
	).Exec(context.Background())

	if err != nil {
		if err == db.ErrNotFound {
			return entities.User{}, entities.CreateException(
				entities.ExceptionCode_RessourceNotFound,
				entities.ExceptionMessage_RessourceNotFound,
			)
		}

		return entities.User{}, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
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
		if err == db.ErrNotFound {
			return entities.User{}, entities.CreateException(
				entities.ExceptionCode_RessourceNotFound,
				entities.ExceptionMessage_RessourceNotFound,
			)
		}

		return entities.User{}, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
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

func (*UserRepository) getFilterQuery(filter entities.Filter) (db.UserWhereParam, *entities.Exception) {
	queries := map[string]db.UserWhereParam{
		"email":            db.User.Email.ContainsIfPresent(filter.Value),
		"firstname":        db.User.Firstname.ContainsIfPresent(filter.Value),
		"lastname":         db.User.Lastname.ContainsIfPresent(filter.Value),
		"authorizationsId": db.User.AuthorizationsID.ContainsIfPresent(filter.Value),
	}

	if query, found := queries[filter.Field]; found {
		return query, nil
	}

	return db.User.ID.Contains(""), entities.CreateException(
		entities.ExceptionCode_BadInputFormat,
		entities.ExceptionMessage_BadInputFormat,
	)
}
