package domain

type RatingRepo interface {
	Create(data Rating) (Rating, error)
	Update(id string, data Rating) (Rating, error)
	Delete(id string) error
	List(listing Listing) ([]Rating, int64, error)
	FindOne(id string) (Rating, error)
}
