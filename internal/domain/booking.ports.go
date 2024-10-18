package domain

type BookingRepo interface {
	Create(data Booking) (Booking, *HttpException)
	Update(id string, data Booking) (Booking, *HttpException)
	Delete(id string) *HttpException
	List(listing Listing) ([]Booking, int64, *HttpException)
	FindOne(id string) (Booking, *HttpException)
}
