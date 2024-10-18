package domain

type UserService struct {
	repo UserRepo
}

func CreateUserService(repo UserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(data User) (User, *HttpException) {
	return s.repo.Create(data)
}

func (s *UserService) Update(id string, data User) (User, *HttpException) {
	return s.repo.Update(id, data)
}

func (s *UserService) Delete(id string) *HttpException {
	return s.repo.Delete(id)
}

func (s *UserService) List(listing Listing) ([]User, int64, *HttpException) {
	return s.repo.List(listing)
}

func (s *UserService) FindOne(id string) (User, *HttpException) {
	return s.repo.FindOne(id)
}
