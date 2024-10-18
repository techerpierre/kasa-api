package infrastructure

import (
	"context"
	"net/http"

	"github.com/techerpierre/kasa-api/internal/domain"
	"github.com/techerpierre/kasa-api/internal/helpers"
	db "github.com/techerpierre/kasa-api/models"
)

type UserRepo struct {
	prisma *db.PrismaClient
}

func CreateUserRepo(prisma *db.PrismaClient) *UserRepo {
	return &UserRepo{
		prisma: prisma,
	}
}

func (r *UserRepo) Create(data domain.User) (domain.User, *domain.HttpException) {
	if !r.validateUser(data) {
		return domain.User{}, &domain.HttpException{
			Status:  http.StatusBadRequest,
			Message: "Some user data required fields missing.",
		}
	}
	result, err := r.prisma.User.CreateOne(
		db.User.Email.Set(data.Email),
		db.User.Password.Set(data.Password),
		db.User.Firstname.Set(data.Firstname),
		db.User.Lastname.Set(data.Lastname),
		db.User.Authorizations.Link(
			db.Authorizations.ID.Equals(db.AuthStatus(data.AuthorizationsID)),
		),
		db.User.Status.Set(db.AuthStatus(data.Status)),
		db.User.Picture.Set(data.Picture),
		db.User.Cover.Set(data.Cover),
	).Exec(context.Background())

	if err != nil {
		return domain.User{}, &domain.HttpException{
			Status:  http.StatusInternalServerError,
			Message: "Unable to create new user.",
		}
	}
	var user domain.User
	if err := helpers.Pipe(&result, &user); err != nil {
		return domain.User{}, &domain.HttpException{
			Status:  http.StatusInternalServerError,
			Message: "Unable to pipe prisma response in user entity.",
		}
	}
	return user, nil
}

func (r *UserRepo) Update(id string, data domain.User) (domain.User, *domain.HttpException) {
	if !r.validateUser(data) {
		return domain.User{}, &domain.HttpException{
			Status:  http.StatusBadRequest,
			Message: "Some user data required fields missing.",
		}
	}
	result, err := r.prisma.User.FindUnique(
		db.User.ID.Equals(id),
	).Update(
		db.User.Email.Set(data.Email),
		db.User.Password.Set(data.Password),
		db.User.Firstname.Set(data.Firstname),
		db.User.Lastname.Set(data.Lastname),
		db.User.Authorizations.Link(
			db.Authorizations.ID.Equals(db.AuthStatus(data.AuthorizationsID)),
		),
		db.User.Status.Set(db.AuthStatus(data.Status)),
		db.User.Picture.Set(data.Picture),
		db.User.Cover.Set(data.Cover),
	).Exec(context.Background())

	if err != nil {
		return domain.User{}, &domain.HttpException{
			Status:  http.StatusInternalServerError,
			Message: "Unable to update this user.",
		}
	}

	if result == nil {
		return domain.User{}, &domain.HttpException{
			Status:  http.StatusNotFound,
			Message: "This user is not found.",
		}
	}

	var user domain.User
	if err := helpers.Pipe(&result, &user); err != nil {
		return domain.User{}, &domain.HttpException{
			Status:  http.StatusInternalServerError,
			Message: "Unable to pipe prisma response in user entity.",
		}
	}
	return user, nil
}

func (r *UserRepo) Delete(id string) *domain.HttpException {
	result, err := r.prisma.User.FindUnique(
		db.User.ID.Equals(id),
	).Delete().Exec(context.Background())

	if err != nil {
		return &domain.HttpException{
			Status:  http.StatusInternalServerError,
			Message: "Unable to Delete this user",
		}
	}

	if result == nil {
		return &domain.HttpException{
			Status:  http.StatusNotFound,
			Message: "This user is not found.",
		}
	}

	return nil
}

func (r *UserRepo) List(listing domain.Listing) ([]domain.User, int64, *domain.HttpException) {
	countResult, err := r.prisma.Prisma.ExecuteRaw(`SELECT COUNT(*) FROM "User"`).Exec(context.Background())

	if err != nil {
		return nil, 0, &domain.HttpException{
			Status:  http.StatusInternalServerError,
			Message: "Unable to count users.",
		}
	}

	skip, take := helpers.Paginate(int(listing.Page), int(listing.Pagesize), countResult.Count)

	results, err := r.prisma.User.FindMany().Skip(skip).Take(take).Skip(skip).Exec(context.Background())

	if err != nil {
		return nil, 0, &domain.HttpException{
			Status:  http.StatusInternalServerError,
			Message: "Unable to list users",
		}
	}

	var users []domain.User

	for _, result := range results {
		var user domain.User
		helpers.Pipe(&result, &user)
		users = append(users, user)
	}

	return users, int64(countResult.Count), nil
}

func (r *UserRepo) FindOne(id string) (domain.User, *domain.HttpException) {
	result, err := r.prisma.User.FindUnique(
		db.User.ID.Equals(id),
	).Exec(context.Background())

	if err != nil {
		return domain.User{}, &domain.HttpException{
			Status:  http.StatusInternalServerError,
			Message: "Unable to find this user.",
		}
	}

	if result == nil {
		return domain.User{}, &domain.HttpException{
			Status:  http.StatusNotFound,
			Message: "This user is not found.",
		}
	}

	var user domain.User
	helpers.Pipe(&result, &user)

	return user, nil
}

func (r *UserRepo) FindOneByEmail(email string) (domain.User, *domain.HttpException) {
	result, err := r.prisma.User.FindUnique(
		db.User.Email.Equals(email),
	).Exec(context.Background())

	if err != nil {
		return domain.User{}, &domain.HttpException{
			Status:  http.StatusInternalServerError,
			Message: "Unable to find this user.",
		}
	}

	if result == nil {
		return domain.User{}, &domain.HttpException{
			Status:  http.StatusNotFound,
			Message: "This user is not found.",
		}
	}

	var user domain.User
	helpers.Pipe(&result, &user)

	return user, nil
}

func (*UserRepo) validateUser(data domain.User) bool {
	requirements := []string{
		"Email",
		"Password",
		"Firstname",
		"Lastname",
		"AuthorizationsID",
	}
	return helpers.ValidateEntity(any(data), requirements)
}
