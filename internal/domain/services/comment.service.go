package services

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type CommentService struct {
	output ports.CommentOutput
}

func CreateCommentService(output ports.CommentOutput) *CommentService {
	return &CommentService{
		output: output,
	}
}

func (s *CommentService) Create(data entities.Comment) (entities.Comment, *entities.Exception) {
	return s.output.Create(data)
}

func (s *CommentService) Update(id string, data entities.Comment) (entities.Comment, *entities.Exception) {
	return s.output.Update(id, data)
}

func (s *CommentService) Delete(id string) *entities.Exception {
	return s.output.Delete(id)
}

func (s *CommentService) List(listing entities.Listing) ([]entities.Comment, int, *entities.Exception) {
	return s.output.List(listing)
}

func (s *CommentService) FindOne(id string) (entities.Comment, *entities.Exception) {
	return s.output.FindOne(id)
}
