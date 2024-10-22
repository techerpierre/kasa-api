package repositories

import (
	"context"

	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/helpers"
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

func (r *BookingRepository) Create(data entities.Booking) (entities.Booking, *entities.Exception) {
	result, err := r.prisma.Booking.CreateOne(
		db.Booking.Start.Set(data.Start),
		db.Booking.End.Set(data.End),
		db.Booking.Accommodation.Link(
			db.Accommodation.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.AccommodationID)),
		),
		db.Booking.Client.Link(
			db.User.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.ClientID)),
		),
	).Exec(context.Background())

	if err != nil {
		return entities.Booking{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.Booking{
		ID:              result.ID,
		Start:           result.Start,
		End:             result.End,
		AccommodationID: result.AccommodationID,
		ClientID:        result.ClientID,
	}, nil
}

func (r *BookingRepository) Update(id string, data entities.Booking) (entities.Booking, *entities.Exception) {
	result, err := r.prisma.Booking.FindUnique(
		db.Booking.ID.Equals(id),
	).Update(
		db.Booking.Start.Set(data.Start),
		db.Booking.End.Set(data.End),
		db.Booking.Accommodation.Link(
			db.Accommodation.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.AccommodationID)),
		),
		db.Booking.Client.Link(
			db.User.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.ClientID)),
		),
	).Exec(context.Background())

	if err != nil {
		return entities.Booking{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.Booking{
		ID:              result.ID,
		Start:           result.Start,
		End:             result.End,
		AccommodationID: result.AccommodationID,
		ClientID:        result.ClientID,
	}, nil
}

func (r *BookingRepository) Delete(id string) *entities.Exception {
	_, err := r.prisma.Booking.FindUnique(
		db.Booking.ID.Equals(id),
	).Delete().Exec(context.Background())

	if err != nil {
		return entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	return nil
}

func (r *BookingRepository) List(listing entities.Listing) ([]entities.Booking, int, *entities.Exception) {
	results, err := r.prisma.Booking.FindMany().Skip(
		listing.Page * listing.Pagesize,
	).Take(listing.Pagesize).Exec(context.Background())

	if err != nil {
		return nil, 0, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	countResult, err := r.prisma.Prisma.ExecuteRaw(
		`SELECT COUNT(*) FROM Booking`,
	).Exec(context.Background())

	if err != nil {
		return nil, 0, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	var bookings []entities.Booking

	for _, result := range results {
		bookings = append(bookings, entities.Booking{
			ID:              result.ID,
			Start:           result.Start,
			End:             result.End,
			AccommodationID: result.AccommodationID,
			ClientID:        result.ClientID,
		})
	}

	return bookings, countResult.Count, nil
}

func (r *BookingRepository) FindOne(id string) (entities.Booking, *entities.Exception) {
	result, err := r.prisma.Booking.FindUnique(
		db.Booking.ID.Equals(id),
	).Exec(context.Background())

	if err != nil {
		if err == db.ErrNotFound {
			return entities.Booking{}, entities.CreateException(
				entities.ExceptionCode_RessourceNotFound,
				entities.ExceptionMessage_RessourceNotFound,
			)
		}

		return entities.Booking{}, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	return entities.Booking{
		ID:              result.ID,
		Start:           result.Start,
		End:             result.End,
		AccommodationID: result.AccommodationID,
		ClientID:        result.ClientID,
	}, nil
}
