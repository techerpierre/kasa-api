package repositories

import (
	"context"

	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/helpers"
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

func (r *AccommodationRepository) Create(data entities.Accommodation) (entities.Accommodation, *entities.Exception) {
	result, err := r.prisma.Accommodation.CreateOne(
		db.Accommodation.Title.SetIfPresent(helpers.NilIfEmptyString(data.Title)),
		db.Accommodation.Description.SetIfPresent(helpers.NilIfEmptyString(data.Description)),
		db.Accommodation.Cover.SetIfPresent(helpers.NilIfEmptyString(data.Cover)),
		db.Accommodation.Adress.SetIfPresent(helpers.NilIfEmptyString(data.Adress)),
		db.Accommodation.Zip.SetIfPresent(helpers.NilIfEmptyString(data.Zip)),
		db.Accommodation.City.SetIfPresent(helpers.NilIfEmptyString(data.City)),
		db.Accommodation.Country.SetIfPresent(helpers.NilIfEmptyString(data.Country)),
		db.Accommodation.Active.Set(data.Active),
		db.Accommodation.User.Link(
			db.User.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.UserID)),
		),
		db.Accommodation.Pictures.SetIfPresent(&data.Pictures),
		db.Accommodation.AdditionalAdress.SetIfPresent(data.AdditionalAdress),
		db.Accommodation.Equipments.SetIfPresent(&data.Equipments),
		db.Accommodation.Tags.SetIfPresent(&data.Tags),
	).Exec(context.Background())

	if err != nil {
		return entities.Accommodation{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.Accommodation{
		ID:               result.ID,
		Title:            result.Title,
		Description:      result.Description,
		Cover:            result.Cover,
		Pictures:         result.Pictures,
		Adress:           result.Adress,
		AdditionalAdress: helpers.PointerFromPrismaField(result.AdditionalAdress()),
		Zip:              result.Zip,
		City:             result.City,
		Country:          result.Country,
		Active:           result.Active,
		Equipments:       result.Equipments,
		Tags:             result.Tags,
		UserID:           result.UserID,
	}, nil
}

func (r *AccommodationRepository) Update(id string, data entities.Accommodation) (entities.Accommodation, *entities.Exception) {
	result, err := r.prisma.Accommodation.FindUnique(
		db.Accommodation.ID.Equals(id),
	).Update(
		db.Accommodation.Title.SetIfPresent(helpers.NilIfEmptyString(data.Title)),
		db.Accommodation.Description.SetIfPresent(helpers.NilIfEmptyString(data.Description)),
		db.Accommodation.Cover.SetIfPresent(helpers.NilIfEmptyString(data.Cover)),
		db.Accommodation.Adress.SetIfPresent(helpers.NilIfEmptyString(data.Adress)),
		db.Accommodation.Zip.SetIfPresent(helpers.NilIfEmptyString(data.Zip)),
		db.Accommodation.City.SetIfPresent(helpers.NilIfEmptyString(data.City)),
		db.Accommodation.Country.SetIfPresent(helpers.NilIfEmptyString(data.Country)),
		db.Accommodation.Active.Set(data.Active),
		db.Accommodation.User.Link(
			db.User.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.UserID)),
		),
		db.Accommodation.Pictures.SetIfPresent(&data.Pictures),
		db.Accommodation.AdditionalAdress.SetIfPresent(data.AdditionalAdress),
		db.Accommodation.Equipments.SetIfPresent(&data.Equipments),
		db.Accommodation.Tags.SetIfPresent(&data.Tags),
	).Exec(context.Background())

	if err != nil {
		if err == db.ErrNotFound {
			return entities.Accommodation{}, entities.CreateException(
				entities.ExceptionCode_RessourceNotFound,
				entities.ExceptionMessage_RessourceNotFound,
			)
		}

		return entities.Accommodation{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.Accommodation{
		ID:               result.ID,
		Title:            result.Title,
		Description:      result.Description,
		Cover:            result.Cover,
		Pictures:         result.Pictures,
		Adress:           result.Adress,
		AdditionalAdress: helpers.PointerFromPrismaField(result.AdditionalAdress()),
		Zip:              result.Zip,
		City:             result.City,
		Country:          result.Country,
		Active:           result.Active,
		Equipments:       result.Equipments,
		Tags:             result.Tags,
		UserID:           result.UserID,
	}, nil
}

func (r *AccommodationRepository) Delete(id string) *entities.Exception {
	_, err := r.prisma.Accommodation.FindUnique(
		db.Accommodation.ID.Equals(id),
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

func (r *AccommodationRepository) List(listing entities.Listing) ([]entities.Accommodation, int, *entities.Exception) {
	results, err := r.prisma.Accommodation.FindMany().Skip(
		listing.Page * listing.Pagesize,
	).Take(listing.Pagesize).Exec(context.Background())

	if err != nil {
		return nil, 0, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	countResult, err := r.prisma.Prisma.ExecuteRaw(
		`SELECT COUNT(*) FROM Accommodation`,
	).Exec(context.Background())

	if err != nil {
		return nil, 0, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	var accommodations []entities.Accommodation

	for _, result := range results {
		accommodations = append(accommodations, entities.Accommodation{
			ID:               result.ID,
			Title:            result.Title,
			Description:      result.Description,
			Cover:            result.Cover,
			Pictures:         result.Pictures,
			Adress:           result.Adress,
			AdditionalAdress: helpers.PointerFromPrismaField(result.AdditionalAdress()),
			Zip:              result.Zip,
			City:             result.City,
			Country:          result.Country,
			Active:           result.Active,
			Equipments:       result.Equipments,
			Tags:             result.Tags,
			UserID:           result.UserID,
		})
	}

	return accommodations, countResult.Count, nil
}

func (r *AccommodationRepository) FindOne(id string) (entities.Accommodation, *entities.Exception) {
	result, err := r.prisma.Accommodation.FindUnique(
		db.Accommodation.ID.Equals(id),
	).Exec(context.Background())

	if err != nil {
		if err == db.ErrNotFound {
			return entities.Accommodation{}, entities.CreateException(
				entities.ExceptionCode_RessourceNotFound,
				entities.ExceptionMessage_RessourceNotFound,
			)
		}

		return entities.Accommodation{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.Accommodation{
		ID:               result.ID,
		Title:            result.Title,
		Description:      result.Description,
		Cover:            result.Cover,
		Pictures:         result.Pictures,
		Adress:           result.Adress,
		AdditionalAdress: helpers.PointerFromPrismaField(result.AdditionalAdress()),
		Zip:              result.Zip,
		City:             result.City,
		Country:          result.Country,
		Active:           result.Active,
		Equipments:       result.Equipments,
		Tags:             result.Tags,
		UserID:           result.UserID,
	}, nil
}
