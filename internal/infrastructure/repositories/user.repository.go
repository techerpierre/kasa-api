package repositories

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
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

func (*UserRepository) Create(data entities.User) (entities.User, *entities.Exception) {
	return entities.User{}, nil
}

func (*UserRepository) Update(id string, data entities.User) (entities.User, *entities.Exception) {
	return entities.User{}, nil
}
func (*UserRepository) Delete(id string) *entities.Exception {
	return nil
}

func (*UserRepository) List(listing entities.Listing) ([]entities.User, int64, *entities.Exception) {
	return nil, 0, nil
}

func (*UserRepository) FindOne(id string) (entities.User, *entities.Exception) {
	return entities.User{}, nil
}

func (*UserRepository) FindOneByEmail(email string) (entities.User, *entities.Exception) {
	return entities.User{}, nil
}
