package repositories

import (
	"github.com/techerpierre/kasa-api/internal/domain/entities"
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

func (*CommentRepository) Create(data entities.Comment) (entities.Comment, *entities.Exception) {
	return entities.Comment{}, nil
}

func (*CommentRepository) Update(id string, data entities.Comment) (entities.Comment, *entities.Exception) {
	return entities.Comment{}, nil
}

func (*CommentRepository) Delete(id string) *entities.Exception {
	return nil
}

func (*CommentRepository) List(listing entities.Listing) ([]entities.Comment, int64, *entities.Exception) {
	return nil, 0, nil
}

func (*CommentRepository) FindOne(id string) (entities.Comment, *entities.Exception) {
	return entities.Comment{}, nil
}
