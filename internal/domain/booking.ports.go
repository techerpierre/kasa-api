package domain

type BookingRepo interface {
	Create(data Booking) (Booking, error)
	Update(id string, data Booking) (Booking, error)
	Delete(id string) error
	List(listing Listing) ([]Booking, int64, error)
	FindOne(id string) (Booking, error)
}
