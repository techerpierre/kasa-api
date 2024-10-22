package repositories

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	db "github.com/techerpierre/kasa-api/models"
)

type AccommodationRepository struct {
	prisma *db.PrismaClient
}

func CreateAccommodationRepository(prisma *db.PrismaClient) *AccommodationRepository {
	return &AccommodationRepository{
		prisma: prisma,
	}
}

func (*AccommodationRepository) Create(data entities.Accommodation) (entities.Accommodation, *entities.Exception) {
	return entities.Accommodation{}, nil
}

func (*AccommodationRepository) Update(id string, data entities.Accommodation) (entities.Accommodation, *entities.Exception) {
	return entities.Accommodation{}, nil
}

func (*AccommodationRepository) Delete(id string) *entities.Exception {
	return nil
}

func (*AccommodationRepository) List(listing entities.Listing) ([]entities.Accommodation, int, *entities.Exception) {
	return nil, 0, nil
}

func (*AccommodationRepository) FindOne(id string) (entities.Accommodation, *entities.Exception) {
	return entities.Accommodation{}, nil
}
