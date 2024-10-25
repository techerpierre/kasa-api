package repositories

import (
	"context"
	"fmt"
	"strconv"

	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/helpers"
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

func (r *RatingRepository) Create(data entities.Rating) (entities.Rating, *entities.Exception) {
	result, err := r.prisma.Rating.CreateOne(
		db.Rating.Value.SetIfPresent(helpers.NilIfEmptyInt(data.Value)),
		db.Rating.Accommodation.Link(
			db.Accommodation.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.AccommodationID)),
		),
		db.Rating.User.Link(
			db.User.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.UserID)),
		),
	).Exec(context.Background())

	if err != nil {
		return entities.Rating{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.Rating{
		ID:              result.ID,
		Value:           result.Value,
		AccommodationID: result.AccommodationID,
		UserID:          result.UserID,
	}, nil
}
func (r *RatingRepository) Update(id string, data entities.Rating) (entities.Rating, *entities.Exception) {
	result, err := r.prisma.Rating.FindUnique(
		db.Rating.ID.Equals(id),
	).Update(
		db.Rating.Value.SetIfPresent(helpers.NilIfEmptyInt(data.Value)),
		db.Rating.Accommodation.Link(
			db.Accommodation.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.AccommodationID)),
		),
		db.Rating.User.Link(
			db.User.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.UserID)),
		),
	).Exec(context.Background())

	if err != nil {
		if err == db.ErrNotFound {
			return entities.Rating{}, entities.CreateException(
				entities.ExceptionCode_RessourceNotFound,
				entities.ExceptionMessage_RessourceNotFound,
			)
		}

		return entities.Rating{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.Rating{
		ID:              result.ID,
		Value:           result.Value,
		AccommodationID: result.AccommodationID,
		UserID:          result.UserID,
	}, nil
}

func (r *RatingRepository) Delete(id string) *entities.Exception {
	_, err := r.prisma.Rating.FindUnique(
		db.Rating.ID.Equals(id),
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

func (r *RatingRepository) List(listing entities.Listing) ([]entities.Rating, int, *entities.Exception) {
	results, err := r.prisma.Rating.FindMany().Skip(
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
		`SELECT COUNT(*) as count FROM "Rating"`,
	).Exec(context.Background(), &countResult)

	if err != nil {
		fmt.Printf("dzeffezzfzef: %s", err.Error())
		return nil, 0, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	count, _ := strconv.Atoi(countResult[0].Count)

	var ratings []entities.Rating

	for _, result := range results {
		ratings = append(ratings, entities.Rating{
			ID:              result.ID,
			Value:           result.Value,
			AccommodationID: result.AccommodationID,
			UserID:          result.UserID,
		})
	}

	return ratings, count, nil
}

func (r *RatingRepository) FindOne(id string) (entities.Rating, *entities.Exception) {
	result, err := r.prisma.Rating.FindUnique(
		db.Rating.ID.Equals(id),
	).Exec(context.Background())

	if err != nil {
		if err == db.ErrNotFound {
			return entities.Rating{}, entities.CreateException(
				entities.ExceptionCode_RessourceNotFound,
				entities.ExceptionMessage_RessourceNotFound,
			)
		}

		return entities.Rating{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.Rating{
		ID:              result.ID,
		Value:           result.Value,
		AccommodationID: result.AccommodationID,
		UserID:          result.UserID,
	}, nil
}
