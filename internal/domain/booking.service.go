package domain

type BookingService struct {
	repo BookingRepo
}

func CreateBookingService(repo BookingRepo) *BookingService {
	return &BookingService{
		repo: repo,
	}
}

func (s *BookingService) Create(data Booking) (Booking, *HttpException) {
	return s.repo.Create(data)
}

func (s *BookingService) Update(id string, data Booking) (Booking, *HttpException) {
	return s.repo.Update(id, data)
}

func (s *BookingService) Delete(id string) *HttpException {
	return s.repo.Delete(id)
}

func (s *BookingService) List(listing Listing) ([]Booking, int64, *HttpException) {
	return s.repo.List(listing)
}

func (s *BookingService) FindOne(id string) (Booking, *HttpException) {
	return s.repo.FindOne(id)
}
