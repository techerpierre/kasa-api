package ports

import "github.com/techerpierre/kasa-api/internal/domain/entities"

type CommentOutput interface {
	Create(data entities.Comment) (entities.Comment, *entities.Exception)
	Update(id string, data entities.Comment) (entities.Comment, *entities.Exception)
	Delete(id string) *entities.Exception
	List(listing entities.Listing) ([]entities.Comment, int64, *entities.Exception)
	FindOne(id string) (entities.Comment, *entities.Exception)
}

type CommentInput interface {
	Create(data entities.Comment) (entities.Comment, *entities.Exception)
	Update(id string, data entities.Comment) (entities.Comment, *entities.Exception)
	Delete(id string) *entities.Exception
	List(listing entities.Listing) ([]entities.Comment, int64, *entities.Exception)
	FindOne(id string) (entities.Comment, *entities.Exception)
}
