package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type CommentRepo interface {
	Create(data entities.Comment) (entities.Comment, *entities.HttpException)
	Update(id string, data entities.Comment) (entities.Comment, *entities.HttpException)
	Delete(id string) *entities.HttpException
	List(listing entities.Listing) ([]entities.Comment, int64, *entities.HttpException)
	FindOne(id string) (entities.Comment, *entities.HttpException)
}
