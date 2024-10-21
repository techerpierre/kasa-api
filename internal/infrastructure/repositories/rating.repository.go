package repositories

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	db "github.com/techerpierre/kasa-api/models"
)

type RatingRepository struct {
	prisma *db.PrismaClient
}

func CreateRatingRepository(prisma *db.PrismaClient) *RatingRepository {
	return &RatingRepository{
		prisma: prisma,
	}
}

func (*RatingRepository) Create(data entities.Rating) (entities.Rating, *entities.Exception) {
	return entities.Rating{}, nil
}
func (*RatingRepository) Update(id string, data entities.Rating) (entities.Rating, *entities.Exception) {
	return entities.Rating{}, nil
}

func (*RatingRepository) Delete(id string) *entities.Exception {
	return nil
}

func (*RatingRepository) List(listing entities.Listing) ([]entities.Rating, int64, *entities.Exception) {
	return nil, 0, nil
}

func (*RatingRepository) FindOne(id string) (entities.Rating, *entities.Exception) {
	return entities.Rating{}, nil
}
