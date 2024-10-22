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

func (*AuthorizationsRepository) Update(id string, data entities.Authorizations) (entities.Authorizations, *entities.Exception) {
	return entities.Authorizations{}, nil
}

func (*AuthorizationsRepository) Delete(id string) *entities.Exception {
	return nil
}

func (*AuthorizationsRepository) List() ([]entities.Authorizations, int, *entities.Exception) {
	return nil, 0, nil
}

func (*AuthorizationsRepository) FindOne(id string) (entities.Authorizations, *entities.Exception) {
	return entities.Authorizations{}, nil
}
