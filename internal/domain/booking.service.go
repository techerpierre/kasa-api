package domain

type BookingService struct {
	repo BookingRepo
}

func CreateBookingService(repo BookingRepo) *BookingService {
	return &BookingService{
		repo: repo,
	}
}

func (s *BookingService) Create(data Booking) (Booking, error) {
	return s.repo.Create(data)
}

func (s *BookingService) Update(id string, data Booking) (Booking, error) {
	return s.repo.Update(id, data)
}

func (s *BookingService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *BookingService) List(listing Listing) ([]Booking, int64, error) {
	return s.repo.List(listing)
}

func (s *BookingService) FindOne(id string) (Booking, error) {
	return s.repo.FindOne(id)
}
