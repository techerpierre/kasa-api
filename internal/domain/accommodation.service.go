package domain

type AccomodationService struct {
	repo AccommodationRepo
}

func CreateAccomodationService(repo AccommodationRepo) *AccomodationService {
	return &AccomodationService{
		repo: repo,
	}
}

func (s *AccomodationService) Create(data Accommodation) (Accommodation, error) {
	return s.repo.Create(data)
}

func (s *AccomodationService) Update(id string, data Accommodation) (Accommodation, error) {
	return s.repo.Update(id, data)
}

func (s *AccomodationService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *AccomodationService) List(listing Listing) ([]Accommodation, int64, error) {
	return s.repo.List(listing)
}

func (s *AccomodationService) FindOne(id string) (Accommodation, error) {
	return s.repo.FindOne(id)
}
