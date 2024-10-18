package domain

type CommentRepo interface {
	Create(data Comment) (Comment, error)
	Update(id string, data Comment) (Comment, error)
	Delete(id string) error
	List(listing Listing) ([]Comment, int64, error)
	FindOne(id string) (Comment, error)
}
