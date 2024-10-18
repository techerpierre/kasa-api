package domain

type AccommodationPorts interface {
	Create(data Accommodation) (Accommodation, error)
	Update(id string, data Accommodation) (Accommodation, error)
	Delete(id string) error
	List(listing Listing) ([]Accommodation, int64, error)
	FindOne(id string)
}
