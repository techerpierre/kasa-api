package repositories

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	db "github.com/techerpierre/kasa-api/models"
)

type BookingRepository struct {
	prisma *db.PrismaClient
}

func CreateBookingRepository(prisma *db.PrismaClient) *BookingRepository {
	return &BookingRepository{
		prisma: prisma,
	}
}

func (*BookingRepository) Create(data entities.Booking) (entities.Booking, *entities.Exception) {
	return entities.Booking{}, nil
}

func (*BookingRepository) Update(id string, data entities.Booking) (entities.Booking, *entities.Exception) {
	return entities.Booking{}, nil
}

func (*BookingRepository) Delete(id string) *entities.Exception {
	return nil
}

func (*BookingRepository) List(listing entities.Listing) ([]entities.Booking, int64, *entities.Exception) {
	return nil, 0, nil
}

func (*BookingRepository) FindOne(id string) (entities.Booking, *entities.Exception) {
	return entities.Booking{}, nil
}
