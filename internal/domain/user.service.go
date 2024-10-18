package domain

type UserService struct {
	repo UserRepo
}

func CreateUserService(repo UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(data User) (User, error) {
	return s.repo.Create(data)
}

func (s *UserService) Update(id string, data User) (User, error) {
	return s.repo.Update(id, data)
}

func (s *UserService) Delete(id string) error {
	return s.repo.Delete(id)
}

func (s *UserService) List(listing Listing) ([]User, int64, error) {
	return s.repo.List(listing)
}

func (s *UserService) FindOne(id string) (User, error) {
	return s.repo.FindOne(id)
}
