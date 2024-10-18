package domain

type RatingRepo interface {
	Create(data Rating) (Rating, *HttpException)
	Update(id string, data Rating) (Rating, *HttpException)
	Delete(id string) *HttpException
	List(listing Listing) ([]Rating, int64, *HttpException)
	FindOne(id string) (Rating, *HttpException)
}
