package api

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/services"
)

type CommentAPI struct {
	service *services.CommentService
}

func CreateCommentAPI(service *services.CommentService) *CommentAPI {
	return &CommentAPI{
		service: service,
	}
}

func (a *CommentAPI) Create(data entities.Comment) (entities.Comment, *entities.Exception) {
	return a.service.Create(data)
}

func (a *CommentAPI) Update(id string, data entities.Comment) (entities.Comment, *entities.Exception) {
	return a.service.Update(id, data)
}

func (a *CommentAPI) Delete(id string) *entities.Exception {
	return a.service.Delete(id)
}

func (a *CommentAPI) List(listing entities.Listing) ([]entities.Comment, int64, *entities.Exception) {
	return a.service.List(listing)
}

func (a *CommentAPI) FindOne(id string) (entities.Comment, *entities.Exception) {
	return a.service.FindOne(id)
}
