package domain

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
	"github.com/techerpierre/kasa-api/internal/domain/ports"
)

type CommentService struct {
	repo ports.CommentRepo
}

func CreateCommentService(repo ports.CommentRepo) *CommentService {
	return &CommentService{
		repo: repo,
	}
}

func (s *CommentService) Create(data entities.Comment) (entities.Comment, *entities.HttpException) {
	return s.repo.Create(data)
}

func (s *CommentService) Update(id string, data entities.Comment) (entities.Comment, *entities.HttpException) {
	return s.repo.Update(id, data)
}

func (s *CommentService) Delete(id string) *entities.HttpException {
	return s.repo.Delete(id)
}

func (s *CommentService) List(listing entities.Listing) ([]entities.Comment, int64, *entities.HttpException) {
	return s.repo.List(listing)
}

func (s *CommentService) FindOne(id string) (entities.Comment, *entities.HttpException) {
	return s.repo.FindOne(id)
}
