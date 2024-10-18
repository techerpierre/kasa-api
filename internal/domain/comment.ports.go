package domain

type CommentRepo interface {
	Create(data Comment) (Comment, *HttpException)
	Update(id string, data Comment) (Comment, *HttpException)
	Delete(id string) *HttpException
	List(listing Listing) ([]Comment, int64, *HttpException)
	FindOne(id string) (Comment, *HttpException)
}
