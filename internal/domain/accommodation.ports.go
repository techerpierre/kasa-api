package domain

type AccommodationRepo interface {
	Create(data Accommodation) (Accommodation, *HttpException)
	Update(id string, data Accommodation) (Accommodation, *HttpException)
	Delete(id string) *HttpException
	List(listing Listing) ([]Accommodation, int64, *HttpException)
	FindOne(id string) (Accommodation, *HttpException)
}
