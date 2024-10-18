package domain

type RatingService struct {
	repo RatingRepo
}

func CreateRatingService(repo RatingRepo) *RatingService {
	return &RatingService{
		repo: repo,
	}
}

func (s *RatingService) Create(data Rating) (Rating, *HttpException) {
	return s.repo.Create(data)
}

func (s *RatingService) Update(id string, data Rating) (Rating, *HttpException) {
	return s.repo.Update(id, data)
}

func (s *RatingService) Delete(id string) *HttpException {
	return s.repo.Delete(id)
}

func (s *RatingService) List(listing Listing) ([]Rating, int64, *HttpException) {
	return s.repo.List(listing)
}

func (s *RatingService) FindOne(id string) (Rating, *HttpException) {
	return s.repo.FindOne(id)
}
