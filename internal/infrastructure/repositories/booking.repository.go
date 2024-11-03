package repositories

import (
	"context"
	"strconv"

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
		if err == db.ErrNotFound {
			return entities.Booking{}, entities.CreateException(
				entities.ExceptionCode_RessourceNotFound,
				entities.ExceptionMessage_RessourceNotFound,
			)
		}

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

func (r *BookingRepository) List(listing entities.Listing) ([]entities.Booking, int, *entities.Exception) {
	var filters []db.BookingWhereParam

	for _, filter := range listing.Filters {
		filterQuery, exception := r.getFilterQuery(filter)
		if exception != nil {
			return nil, 0, exception
		}
		filters = append(filters, filterQuery)
	}

	results, err := r.prisma.Booking.FindMany(filters...).Skip(
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
		`SELECT COUNT(*) FROM "Booking"`,
	).Exec(context.Background(), &countResult)

	if err != nil {
		return nil, 0, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	count, _ := strconv.Atoi(countResult[0].Count)

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

	return bookings, count, nil
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

func (*BookingRepository) getFilterQuery(filter entities.Filter) (db.BookingWhereParam, *entities.Exception) {
	queries := map[string]db.BookingWhereParam{
		"start":           db.Booking.Start.EqualsIfPresent(helpers.StringToTime(filter.Value, "0000-00-00 00:00:00")),
		"end":             db.Booking.End.EqualsIfPresent(helpers.StringToTime(filter.Value, "0000-00-00 00:00:00")),
		"accommodationId": db.Booking.AccommodationID.ContainsIfPresent(filter.Value),
		"clientId":        db.Booking.ClientID.ContainsIfPresent(filter.Value),
	}

	if query, found := queries[filter.Field]; found {
		return query, nil
	}

	return db.Booking.ID.Contains(""), entities.CreateException(
		entities.ExceptionCode_BadInputFormat,
		entities.ExceptionMessage_BadInputFormat,
	)
}
