package repositories

import (
	"context"
	"strconv"

	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/helpers"
	db "github.com/techerpierre/kasa-api/models"
)

type CommentRepository struct {
	prisma *db.PrismaClient
}

func CreateCommentRepository(prisma *db.PrismaClient) *CommentRepository {
	return &CommentRepository{
		prisma: prisma,
	}
}

func (r *CommentRepository) Create(data entities.Comment) (entities.Comment, *entities.Exception) {
	result, err := r.prisma.Comment.CreateOne(
		db.Comment.Content.SetIfPresent(helpers.NilIfEmptyString(data.Content)),
		db.Comment.Accommodation.Link(
			db.Accommodation.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.AccommodationID)),
		),
		db.Comment.User.Link(
			db.User.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.UserID)),
		),
	).Exec(context.Background())

	if err != nil {
		return entities.Comment{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.Comment{
		ID:              result.ID,
		Content:         result.Content,
		AccommodationID: result.AccommodationID,
		UserID:          result.UserID,
	}, nil
}

func (r *CommentRepository) Update(id string, data entities.Comment) (entities.Comment, *entities.Exception) {
	result, err := r.prisma.Comment.FindUnique(
		db.Comment.ID.Equals(id),
	).Update(
		db.Comment.Content.SetIfPresent(helpers.NilIfEmptyString(data.Content)),
		db.Comment.Accommodation.Link(
			db.Accommodation.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.AccommodationID)),
		),
		db.Comment.User.Link(
			db.User.ID.EqualsIfPresent(helpers.NilIfEmptyString(data.UserID)),
		),
	).Exec(context.Background())

	if err != nil {
		if err == db.ErrNotFound {
			return entities.Comment{}, entities.CreateException(
				entities.ExceptionCode_RessourceNotFound,
				entities.ExceptionMessage_RessourceNotFound,
			)
		}

		return entities.Comment{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.Comment{
		ID:              result.ID,
		Content:         result.Content,
		AccommodationID: result.AccommodationID,
		UserID:          result.UserID,
	}, nil
}

func (r *CommentRepository) Delete(id string) *entities.Exception {
	_, err := r.prisma.Comment.FindUnique(
		db.Comment.ID.Equals(id),
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

func (r *CommentRepository) List(listing entities.Listing) ([]entities.Comment, int, *entities.Exception) {
	var filters []db.CommentWhereParam

	for _, filter := range listing.Filters {
		filterQuery, exception := r.getFilterQuery(filter)
		if exception != nil {
			return nil, 0, exception
		}
		filters = append(filters, filterQuery)
	}

	results, err := r.prisma.Comment.FindMany(filters...).Skip(
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
		`SELECT COUNT(*) FROM "Comment"`,
	).Exec(context.Background(), &countResult)

	if err != nil {
		return nil, 0, entities.CreateException(
			entities.ExceptionCode_NotHandledError,
			entities.ExceptionMessage_NotHandledError,
		)
	}

	count, _ := strconv.Atoi(countResult[0].Count)

	var comments []entities.Comment

	for _, result := range results {
		comments = append(comments, entities.Comment{
			ID:              result.ID,
			Content:         result.Content,
			AccommodationID: result.AccommodationID,
			UserID:          result.UserID,
		})
	}

	return comments, count, nil
}

func (r *CommentRepository) FindOne(id string) (entities.Comment, *entities.Exception) {
	result, err := r.prisma.Comment.FindUnique(
		db.Comment.ID.Equals(id),
	).Exec(context.Background())

	if err != nil {
		if err == db.ErrNotFound {
			return entities.Comment{}, entities.CreateException(
				entities.ExceptionCode_RessourceNotFound,
				entities.ExceptionMessage_RessourceNotFound,
			)
		}

		return entities.Comment{}, entities.CreateException(
			entities.ExceptionCode_BadInputFormat,
			entities.ExceptionMessage_BadInputFormat,
		)
	}

	return entities.Comment{
		ID:              result.ID,
		Content:         result.Content,
		AccommodationID: result.AccommodationID,
		UserID:          result.UserID,
	}, nil
}

func (*CommentRepository) getFilterQuery(fiter entities.Filter) (db.CommentWhereParam, *entities.Exception) {
	queries := map[string]db.CommentWhereParam{
		"content":         db.Comment.Content.EqualsIfPresent(fiter.Value),
		"accommodationId": db.Comment.AccommodationID.EqualsIfPresent(fiter.Value),
		"userId":          db.Comment.UserID.EqualsIfPresent(fiter.Value),
	}

	if query, found := queries[fiter.Field]; found {
		return query, nil
	}

	return db.Comment.ID.Contains(""), entities.CreateException(
		entities.ExceptionCode_BadInputFormat,
		entities.ExceptionMessage_BadInputFormat,
	)
}
