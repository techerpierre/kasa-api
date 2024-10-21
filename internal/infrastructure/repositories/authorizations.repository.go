package repositories

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
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

func (*AuthorizationsRepository) Create(data entities.Authorizations) (entities.Authorizations, *entities.Exception) {
	return entities.Authorizations{}, nil
}

func (*AuthorizationsRepository) Update(id entities.AuthStatus, data entities.Authorizations) (entities.Authorizations, *entities.Exception) {
	return entities.Authorizations{}, nil
}

func (*AuthorizationsRepository) Delete(id entities.AuthStatus) *entities.Exception {
	return nil
}

func (*AuthorizationsRepository) List() ([]entities.Authorizations, int64, *entities.Exception) {
	return nil, 0, nil
}

func (*AuthorizationsRepository) FindOne(id entities.AuthStatus) (entities.Authorizations, *entities.Exception) {
	return entities.Authorizations{}, nil
}
